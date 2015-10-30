package goview

var builtInTypes = []string{
	"int",
	"uint",
	"int64",
	"uint64",
	"int32",
	"uint32",
	"int16",
	"uint16",
	"int8",
	"uint8",
	"byte",
	"rune",
	"error",
	"string",
	"float32",
	"float64",
	"complex64",
	"complex128",
	"uintptr",
	"bool",
	"map",
	"true",
	"false",
	"nil",
	"iota",
}

var builtInFuncs = []string{
	"len",
	"cap",
	"close",
	"complex",
	"delete",
	"imag",
	"panic",
	"print",
	"println",
	"real",
	"recover",
	"make",
	"append",
	"new",
	"copy",
}

var (
	builtInFuncMap = makeMap(builtInFuncs)
	builtInTypeMap = makeMap(builtInTypes)
)

func makeMap(lst []string) map[string]struct{} {
	ret := make(map[string]struct{})
	for _, s := range lst {
		ret[s] = struct{}{}
	}

	return ret
}
