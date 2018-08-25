package gcimporter

import (
	"go/token"
	"go/types"
	"strings"
)

// parent is the package which declared the type; parent == nil means
// the package currently imported. The parent package is needed for
// exported struct fields and interface methods which don't contain
// explicit package information in the export data.
func readType(p *reader, parent *types.Package) types.Type {
	// if the type was seen before, i is its index (>= 0)
	i := readTagOrIndex(p)
	if i >= 0 {
		return p.typList[i]
	}

	// otherwise, i is the type tag (< 0)
	switch i {
	case namedTag:
		// read type object
		pos := readPos(p)
		parent, name := readQualifiedName(p)
		scope := parent.Scope()
		obj := scope.Lookup(name)

		// if the object doesn't exist yet, create and insert it
		if obj == nil {
			obj = types.NewTypeName(pos, parent, name, nil)
			scope.Insert(obj)
		}

		if _, ok := obj.(*types.TypeName); !ok {
			errorf("pkg = %s, name = %s => %s", parent, name, obj)
		}

		// associate new named type with obj if it doesn't exist yet
		t0 := types.NewNamed(obj.(*types.TypeName), nil, nil)

		// but record the existing type, if any
		t := obj.Type().(*types.Named)
		p.record(t)

		// read underlying type
		t0.SetUnderlying(readType(p, parent))

		// interfaces don't have associated methods
		if types.IsInterface(t0) {
			return t
		}

		// read associated methods
		for i := readInt(p); i > 0; i-- {
			// TODO(gri) replace this with something closer to fieldName
			pos := readPos(p)
			name := readString(p)
			if !exported(name) {
				readPkg(p)
			}

			// TODO(gri): do we need a full param list for the receiver?
			recv, _ := readParamList(p)
			params, isddd := readParamList(p)
			result, _ := readParamList(p)
			readInt(p) // go:nointerface pragma - discarded

			sig := types.NewSignature(recv.At(0), params, result, isddd)
			t0.AddMethod(types.NewFunc(pos, parent, name, sig))
		}

		return t

	case arrayTag:
		t := new(types.Array)
		if p.trackAllTypes {
			p.record(t)
		}

		n := readInt64(p)
		*t = *types.NewArray(readType(p, parent), n)
		return t

	case sliceTag:
		t := new(types.Slice)
		if p.trackAllTypes {
			p.record(t)
		}

		*t = *types.NewSlice(readType(p, parent))
		return t

	case dddTag:
		t := new(dddSlice)
		if p.trackAllTypes {
			p.record(t)
		}

		t.elem = readType(p, parent)
		return t

	case structTag:
		t := new(types.Struct)
		if p.trackAllTypes {
			p.record(t)
		}

		*t = *types.NewStruct(readFieldList(p, parent))
		return t

	case pointerTag:
		t := new(types.Pointer)
		if p.trackAllTypes {
			p.record(t)
		}

		*t = *types.NewPointer(readType(p, parent))
		return t

	case signatureTag:
		t := new(types.Signature)
		if p.trackAllTypes {
			p.record(t)
		}

		params, isddd := readParamList(p)
		result, _ := readParamList(p)
		*t = *types.NewSignature(nil, params, result, isddd)
		return t

	case interfaceTag:
		// Create a dummy entry in the type list. This is safe because we
		// cannot expect the interface type to appear in a cycle, as any
		// such cycle must contain a named type which would have been
		// first defined earlier.
		n := len(p.typList)
		if p.trackAllTypes {
			p.record(nil)
		}

		var embeddeds []*types.Named
		for n := readInt(p); n > 0; n-- {
			readPos(p)
			embeddeds = append(embeddeds, readType(p, parent).(*types.Named))
		}

		t := types.NewInterface(readMethodList(p, parent), embeddeds)
		p.interfaceList = append(p.interfaceList, t)
		if p.trackAllTypes {
			p.typList[n] = t
		}
		return t

	case mapTag:
		t := new(types.Map)
		if p.trackAllTypes {
			p.record(t)
		}

		key := readType(p, parent)
		val := readType(p, parent)
		*t = *types.NewMap(key, val)
		return t

	case chanTag:
		t := new(types.Chan)
		if p.trackAllTypes {
			p.record(t)
		}

		var dir types.ChanDir
		// tag values must match the constants in cmd/compile/internal/gc/go.go
		switch d := readInt(p); d {
		case 1 /* Crecv */ :
			dir = types.RecvOnly
		case 2 /* Csend */ :
			dir = types.SendOnly
		case 3 /* Cboth */ :
			dir = types.SendRecv
		default:
			errorf("unexpected channel dir %d", d)
		}
		val := readType(p, parent)
		*t = *types.NewChan(dir, val)
		return t

	default:
		errorf("unexpected type tag %d", i) // panics
		panic("unreachable")
	}
}

func readParam(p *reader, named bool) (*types.Var, bool) {
	t := readType(p, nil)
	td, isddd := t.(*dddSlice)
	if isddd {
		t = types.NewSlice(td.elem)
	}

	var pkg *types.Package
	var name string
	if named {
		name = readString(p)
		if name == "" {
			errorf("expected named parameter")
		}
		if name != "_" {
			pkg = readPkg(p)
		}
		if i := strings.Index(name, "·"); i > 0 {
			name = name[:i] // cut off gc-specific parameter numbering
		}
	}

	// read and discard compiler-specific info
	readString(p)

	return types.NewVar(token.NoPos, pkg, name, t), isddd
}

func readParamList(p *reader) (*types.Tuple, bool) {
	n := readInt(p)
	if n == 0 {
		return nil, false
	}
	// negative length indicates unnamed parameters
	named := true
	if n < 0 {
		n = -n
		named = false
	}
	// n > 0
	params := make([]*types.Var, n)
	isddd := false
	for i := range params {
		params[i], isddd = readParam(p, named)
	}
	return types.NewTuple(params...), isddd
}

func readMethodList(p *reader, parent *types.Package) []*types.Func {
	var methods []*types.Func
	if n := readInt(p); n > 0 {
		methods = make([]*types.Func, n)
		for i := range methods {
			methods[i] = readMethod(p, parent)
		}
	}
	return methods
}

func readMethod(p *reader, parent *types.Package) *types.Func {
	pos := readPos(p)
	pkg, name, _ := readFieldName(p, parent)
	params, isddd := readParamList(p)
	result, _ := readParamList(p)
	sig := types.NewSignature(nil, params, result, isddd)
	return types.NewFunc(pos, pkg, name, sig)
}

func readField(p *reader, parent *types.Package) (*types.Var, string) {
	pos := readPos(p)
	pkg, name, alias := readFieldName(p, parent)
	typ := readType(p, parent)
	tag := readString(p)

	anonymous := false
	if name == "" {
		// anonymous field - typ must be T or *T and T must be a type name
		switch typ := deref(typ).(type) {
		case *types.Basic: // basic types are named types
			pkg = nil // // objects defined in Universe scope have no package
			name = typ.Name()
		case *types.Named:
			name = typ.Obj().Name()
		default:
			errorf("named base type expected")
		}
		anonymous = true
	} else if alias {
		// anonymous field: we have an explicit name because it's an alias
		anonymous = true
	}

	return types.NewField(pos, pkg, name, typ, anonymous), tag
}

func readFieldList(p *reader, parent *types.Package) (
	fields []*types.Var, tags []string,
) {
	if n := readInt(p); n > 0 {
		fields = make([]*types.Var, n)
		tags = make([]string, n)
		for i := range fields {
			fields[i], tags[i] = readField(p, parent)
		}
	}
	return
}
