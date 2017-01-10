var dag = {
    "h": 56,
    "w": 16,
    "n": {
        "archive/tar": {
            "x": 12,
            "y": 14,
            "i": [
                "fmt",
                "io/ioutil",
                "path"
            ],
            "o": []
        },
        "archive/zip": {
            "x": 12,
            "y": 32,
            "i": [
                "compress/flate",
                "encoding/binary",
                "hash/crc32",
                "io/ioutil",
                "path"
            ],
            "o": []
        },
        "bufio": {
            "x": 8,
            "y": 42,
            "i": [
                "bytes"
            ],
            "o": [
                "compress/bzip2",
                "compress/flate",
                "compress/lzw",
                "crypto/rand",
                "encoding/csv",
                "encoding/gob",
                "encoding/xml",
                "go/build",
                "image",
                "internal/trace",
                "mime",
                "mime/quotedprintable",
                "net/http/internal",
                "net/textproto",
                "runtime/pprof"
            ]
        },
        "bytes": {
            "x": 7,
            "y": 31,
            "i": [
                "io",
                "unicode",
                "unicode/utf8"
            ],
            "o": [
                "bufio",
                "debug/gosym",
                "debug/macho",
                "encoding/base32",
                "encoding/hex",
                "encoding/json",
                "encoding/pem",
                "go/scanner",
                "io/ioutil",
                "math/big",
                "net/url",
                "os/exec",
                "regexp/syntax",
                "text/scanner",
                "text/tabwriter",
                "text/template/parse",
                "vendor/golang_org/x/net/http2/hpack"
            ]
        },
        "compress/bzip2": {
            "x": 9,
            "y": 17,
            "i": [
                "bufio",
                "sort"
            ],
            "o": []
        },
        "compress/flate": {
            "x": 11,
            "y": 22,
            "i": [
                "bufio",
                "fmt",
                "sort"
            ],
            "o": [
                "archive/zip",
                "compress/gzip",
                "compress/zlib"
            ]
        },
        "compress/gzip": {
            "x": 12,
            "y": 41,
            "i": [
                "compress/flate",
                "encoding/binary",
                "hash/crc32"
            ],
            "o": [
                "net/http"
            ]
        },
        "compress/lzw": {
            "x": 9,
            "y": 39,
            "i": [
                "bufio",
                "fmt"
            ],
            "o": [
                "image/gif"
            ]
        },
        "compress/zlib": {
            "x": 12,
            "y": 43,
            "i": [
                "compress/flate",
                "hash/adler32"
            ],
            "o": [
                "debug/elf",
                "image/png"
            ]
        },
        "container/heap": {
            "x": 12,
            "y": 2,
            "i": [
                "sort"
            ],
            "o": [
                "go/types"
            ]
        },
        "container/list": {
            "x": 11,
            "y": 5,
            "i": [],
            "o": [
                "crypto/tls"
            ]
        },
        "container/ring": {
            "x": 0,
            "y": 29,
            "i": [],
            "o": []
        },
        "context": {
            "x": 9,
            "y": 13,
            "i": [
                "fmt"
            ],
            "o": [
                "net",
                "os/exec"
            ]
        },
        "crypto": {
            "x": 8,
            "y": 36,
            "i": [
                "hash",
                "strconv"
            ],
            "o": [
                "crypto/md5",
                "crypto/rsa",
                "crypto/sha1",
                "crypto/sha256",
                "crypto/sha512"
            ]
        },
        "crypto/aes": {
            "x": 8,
            "y": 20,
            "i": [
                "crypto/cipher",
                "strconv"
            ],
            "o": [
                "crypto/ecdsa",
                "crypto/rand"
            ]
        },
        "crypto/cipher": {
            "x": 7,
            "y": 21,
            "i": [
                "crypto/subtle",
                "io"
            ],
            "o": [
                "crypto/aes",
                "crypto/des"
            ]
        },
        "crypto/des": {
            "x": 10,
            "y": 11,
            "i": [
                "crypto/cipher",
                "encoding/binary"
            ],
            "o": [
                "crypto/x509"
            ]
        },
        "crypto/dsa": {
            "x": 10,
            "y": 47,
            "i": [
                "math/big"
            ],
            "o": [
                "crypto/x509"
            ]
        },
        "crypto/ecdsa": {
            "x": 10,
            "y": 38,
            "i": [
                "crypto/aes",
                "crypto/elliptic",
                "crypto/sha512",
                "encoding/asn1"
            ],
            "o": [
                "crypto/x509"
            ]
        },
        "crypto/elliptic": {
            "x": 9,
            "y": 43,
            "i": [
                "math/big"
            ],
            "o": [
                "crypto/ecdsa"
            ]
        },
        "crypto/hmac": {
            "x": 11,
            "y": 17,
            "i": [
                "crypto/subtle",
                "hash"
            ],
            "o": [
                "crypto/tls"
            ]
        },
        "crypto/md5": {
            "x": 10,
            "y": 49,
            "i": [
                "crypto"
            ],
            "o": [
                "crypto/x509"
            ]
        },
        "crypto/rand": {
            "x": 9,
            "y": 34,
            "i": [
                "bufio",
                "crypto/aes",
                "math/big"
            ],
            "o": [
                "crypto/rsa",
                "mime/multipart"
            ]
        },
        "crypto/rc4": {
            "x": 11,
            "y": 9,
            "i": [
                "strconv"
            ],
            "o": [
                "crypto/tls"
            ]
        },
        "crypto/rsa": {
            "x": 10,
            "y": 43,
            "i": [
                "crypto",
                "crypto/rand"
            ],
            "o": [
                "crypto/x509"
            ]
        },
        "crypto/sha1": {
            "x": 10,
            "y": 51,
            "i": [
                "crypto"
            ],
            "o": [
                "crypto/x509"
            ]
        },
        "crypto/sha256": {
            "x": 10,
            "y": 53,
            "i": [
                "crypto"
            ],
            "o": [
                "crypto/x509"
            ]
        },
        "crypto/sha512": {
            "x": 9,
            "y": 45,
            "i": [
                "crypto"
            ],
            "o": [
                "crypto/ecdsa"
            ]
        },
        "crypto/subtle": {
            "x": 6,
            "y": 22,
            "i": [],
            "o": [
                "crypto/cipher",
                "crypto/hmac"
            ]
        },
        "crypto/tls": {
            "x": 12,
            "y": 17,
            "i": [
                "container/list",
                "crypto/hmac",
                "crypto/rc4",
                "crypto/x509"
            ],
            "o": [
                "net/http",
                "net/smtp"
            ]
        },
        "crypto/x509": {
            "x": 11,
            "y": 38,
            "i": [
                "crypto/des",
                "crypto/dsa",
                "crypto/ecdsa",
                "crypto/md5",
                "crypto/rsa",
                "crypto/sha1",
                "crypto/sha256",
                "crypto/x509/pkix",
                "encoding/hex",
                "encoding/pem",
                "io/ioutil",
                "net",
                "os/exec"
            ],
            "o": [
                "crypto/tls"
            ]
        },
        "crypto/x509/pkix": {
            "x": 10,
            "y": 55,
            "i": [
                "encoding/asn1"
            ],
            "o": [
                "crypto/x509"
            ]
        },
        "database/sql": {
            "x": 9,
            "y": 15,
            "i": [
                "database/sql/driver",
                "sort"
            ],
            "o": []
        },
        "database/sql/driver": {
            "x": 8,
            "y": 16,
            "i": [
                "fmt"
            ],
            "o": [
                "database/sql"
            ]
        },
        "debug/dwarf": {
            "x": 12,
            "y": 19,
            "i": [
                "encoding/binary",
                "fmt",
                "path",
                "sort"
            ],
            "o": [
                "debug/elf",
                "debug/macho",
                "debug/pe"
            ]
        },
        "debug/elf": {
            "x": 13,
            "y": 34,
            "i": [
                "compress/zlib",
                "debug/dwarf"
            ],
            "o": [
                "go/internal/gccgoimporter"
            ]
        },
        "debug/gosym": {
            "x": 8,
            "y": 24,
            "i": [
                "bytes",
                "encoding/binary",
                "fmt",
                "strings"
            ],
            "o": []
        },
        "debug/macho": {
            "x": 13,
            "y": 13,
            "i": [
                "bytes",
                "debug/dwarf"
            ],
            "o": []
        },
        "debug/pe": {
            "x": 13,
            "y": 10,
            "i": [
                "debug/dwarf"
            ],
            "o": []
        },
        "debug/plan9obj": {
            "x": 8,
            "y": 18,
            "i": [
                "encoding/binary",
                "fmt"
            ],
            "o": []
        },
        "encoding": {
            "x": 11,
            "y": 7,
            "i": [],
            "o": [
                "encoding/gob",
                "encoding/json",
                "encoding/xml"
            ]
        },
        "encoding/ascii85": {
            "x": 6,
            "y": 31,
            "i": [
                "io",
                "strconv"
            ],
            "o": []
        },
        "encoding/asn1": {
            "x": 9,
            "y": 41,
            "i": [
                "math/big"
            ],
            "o": [
                "crypto/ecdsa",
                "crypto/x509/pkix"
            ]
        },
        "encoding/base32": {
            "x": 8,
            "y": 34,
            "i": [
                "bytes",
                "strconv",
                "strings"
            ],
            "o": []
        },
        "encoding/base64": {
            "x": 9,
            "y": 37,
            "i": [
                "io",
                "strconv"
            ],
            "o": [
                "encoding/json",
                "encoding/pem",
                "mime"
            ]
        },
        "encoding/binary": {
            "x": 7,
            "y": 23,
            "i": [
                "io",
                "reflect"
            ],
            "o": [
                "archive/zip",
                "compress/gzip",
                "crypto/des",
                "debug/dwarf",
                "debug/gosym",
                "debug/plan9obj",
                "image/png",
                "index/suffixarray",
                "math/big"
            ]
        },
        "encoding/csv": {
            "x": 9,
            "y": 28,
            "i": [
                "bufio",
                "fmt",
                "strings"
            ],
            "o": []
        },
        "encoding/gob": {
            "x": 13,
            "y": 27,
            "i": [
                "bufio",
                "encoding",
                "fmt"
            ],
            "o": [
                "net/rpc"
            ]
        },
        "encoding/hex": {
            "x": 10,
            "y": 45,
            "i": [
                "bytes",
                "fmt"
            ],
            "o": [
                "crypto/x509"
            ]
        },
        "encoding/json": {
            "x": 12,
            "y": 26,
            "i": [
                "bytes",
                "encoding",
                "encoding/base64",
                "fmt",
                "sort",
                "strings",
                "unicode/utf16"
            ],
            "o": [
                "expvar",
                "html/template"
            ]
        },
        "encoding/pem": {
            "x": 10,
            "y": 28,
            "i": [
                "bytes",
                "encoding/base64",
                "sort",
                "strings"
            ],
            "o": [
                "crypto/x509"
            ]
        },
        "encoding/xml": {
            "x": 12,
            "y": 34,
            "i": [
                "bufio",
                "encoding",
                "fmt",
                "strings"
            ],
            "o": []
        },
        "errors": {
            "x": 3,
            "y": 26,
            "i": [],
            "o": [
                "io",
                "strconv",
                "syscall"
            ]
        },
        "expvar": {
            "x": 14,
            "y": 32,
            "i": [
                "encoding/json",
                "net/http"
            ],
            "o": []
        },
        "flag": {
            "x": 13,
            "y": 19,
            "i": [
                "fmt",
                "sort"
            ],
            "o": [
                "net/http/httptest",
                "testing",
                "testing/quick"
            ]
        },
        "fmt": {
            "x": 7,
            "y": 25,
            "i": [
                "os",
                "reflect"
            ],
            "o": [
                "archive/tar",
                "compress/flate",
                "compress/lzw",
                "context",
                "database/sql/driver",
                "debug/dwarf",
                "debug/gosym",
                "debug/plan9obj",
                "encoding/csv",
                "encoding/gob",
                "encoding/hex",
                "encoding/json",
                "encoding/xml",
                "flag",
                "go/token",
                "log",
                "math/big",
                "mime",
                "mime/quotedprintable",
                "net/http/internal",
                "net/internal/socktest",
                "net/url",
                "os/user",
                "runtime/pprof",
                "text/scanner",
                "text/template/parse",
                "vendor/golang_org/x/net/http2/hpack"
            ]
        },
        "go/ast": {
            "x": 11,
            "y": 13,
            "i": [
                "go/scanner"
            ],
            "o": [
                "go/doc",
                "go/parser",
                "go/printer"
            ]
        },
        "go/build": {
            "x": 13,
            "y": 21,
            "i": [
                "bufio",
                "go/doc",
                "go/parser",
                "log"
            ],
            "o": [
                "go/internal/gcimporter"
            ]
        },
        "go/constant": {
            "x": 12,
            "y": 8,
            "i": [
                "go/token",
                "math/big"
            ],
            "o": [
                "go/types"
            ]
        },
        "go/doc": {
            "x": 12,
            "y": 21,
            "i": [
                "go/ast",
                "path",
                "regexp",
                "text/template"
            ],
            "o": [
                "go/build"
            ]
        },
        "go/format": {
            "x": 14,
            "y": 17,
            "i": [
                "go/parser",
                "go/printer"
            ],
            "o": []
        },
        "go/importer": {
            "x": 15,
            "y": 24,
            "i": [
                "go/internal/gccgoimporter",
                "go/internal/gcimporter"
            ],
            "o": []
        },
        "go/internal/gccgoimporter": {
            "x": 14,
            "y": 24,
            "i": [
                "debug/elf",
                "go/types",
                "os/exec",
                "text/scanner"
            ],
            "o": [
                "go/importer"
            ]
        },
        "go/internal/gcimporter": {
            "x": 14,
            "y": 22,
            "i": [
                "go/build",
                "go/types",
                "text/scanner"
            ],
            "o": [
                "go/importer"
            ]
        },
        "go/parser": {
            "x": 12,
            "y": 12,
            "i": [
                "go/ast",
                "io/ioutil"
            ],
            "o": [
                "go/build",
                "go/format",
                "go/types"
            ]
        },
        "go/printer": {
            "x": 13,
            "y": 36,
            "i": [
                "go/ast",
                "text/tabwriter"
            ],
            "o": [
                "go/format"
            ]
        },
        "go/scanner": {
            "x": 10,
            "y": 14,
            "i": [
                "bytes",
                "go/token",
                "path/filepath"
            ],
            "o": [
                "go/ast"
            ]
        },
        "go/token": {
            "x": 9,
            "y": 19,
            "i": [
                "fmt",
                "sort"
            ],
            "o": [
                "go/constant",
                "go/scanner"
            ]
        },
        "go/types": {
            "x": 13,
            "y": 8,
            "i": [
                "container/heap",
                "go/constant",
                "go/parser"
            ],
            "o": [
                "go/internal/gccgoimporter",
                "go/internal/gcimporter"
            ]
        },
        "hash": {
            "x": 7,
            "y": 35,
            "i": [
                "io"
            ],
            "o": [
                "crypto",
                "crypto/hmac",
                "hash/adler32",
                "hash/crc32",
                "hash/crc64",
                "hash/fnv"
            ]
        },
        "hash/adler32": {
            "x": 11,
            "y": 47,
            "i": [
                "hash"
            ],
            "o": [
                "compress/zlib"
            ]
        },
        "hash/crc32": {
            "x": 11,
            "y": 45,
            "i": [
                "hash"
            ],
            "o": [
                "archive/zip",
                "compress/gzip",
                "image/png"
            ]
        },
        "hash/crc64": {
            "x": 8,
            "y": 44,
            "i": [
                "hash"
            ],
            "o": []
        },
        "hash/fnv": {
            "x": 8,
            "y": 46,
            "i": [
                "hash"
            ],
            "o": []
        },
        "html": {
            "x": 12,
            "y": 50,
            "i": [
                "strings"
            ],
            "o": [
                "html/template"
            ]
        },
        "html/template": {
            "x": 13,
            "y": 40,
            "i": [
                "encoding/json",
                "html",
                "text/template"
            ],
            "o": [
                "net/http/pprof",
                "net/rpc"
            ]
        },
        "image": {
            "x": 12,
            "y": 38,
            "i": [
                "bufio",
                "image/color",
                "strconv"
            ],
            "o": [
                "image/internal/imageutil",
                "image/png"
            ]
        },
        "image/color": {
            "x": 11,
            "y": 49,
            "i": [],
            "o": [
                "image",
                "image/color/palette"
            ]
        },
        "image/color/palette": {
            "x": 14,
            "y": 49,
            "i": [
                "image/color"
            ],
            "o": [
                "image/gif"
            ]
        },
        "image/draw": {
            "x": 14,
            "y": 44,
            "i": [
                "image/internal/imageutil"
            ],
            "o": [
                "image/gif"
            ]
        },
        "image/gif": {
            "x": 15,
            "y": 44,
            "i": [
                "compress/lzw",
                "image/color/palette",
                "image/draw"
            ],
            "o": []
        },
        "image/internal/imageutil": {
            "x": 13,
            "y": 44,
            "i": [
                "image"
            ],
            "o": [
                "image/draw",
                "image/jpeg"
            ]
        },
        "image/jpeg": {
            "x": 14,
            "y": 46,
            "i": [
                "image/internal/imageutil"
            ],
            "o": []
        },
        "image/png": {
            "x": 13,
            "y": 38,
            "i": [
                "compress/zlib",
                "encoding/binary",
                "hash/crc32",
                "image"
            ],
            "o": []
        },
        "index/suffixarray": {
            "x": 12,
            "y": 6,
            "i": [
                "encoding/binary",
                "regexp"
            ],
            "o": []
        },
        "internal/nettrace": {
            "x": 9,
            "y": 9,
            "i": [],
            "o": [
                "net"
            ]
        },
        "internal/race": {
            "x": 2,
            "y": 26,
            "i": [
                "unsafe"
            ],
            "o": [
                "sync"
            ]
        },
        "internal/singleflight": {
            "x": 9,
            "y": 11,
            "i": [
                "sync"
            ],
            "o": [
                "net"
            ]
        },
        "internal/syscall/windows/sysdll": {
            "x": 0,
            "y": 25,
            "i": [],
            "o": []
        },
        "internal/testenv": {
            "x": 15,
            "y": 26,
            "i": [
                "os/exec",
                "testing"
            ],
            "o": []
        },
        "internal/trace": {
            "x": 11,
            "y": 20,
            "i": [
                "bufio",
                "math/rand",
                "os/exec"
            ],
            "o": []
        },
        "io": {
            "x": 5,
            "y": 30,
            "i": [
                "errors",
                "sync"
            ],
            "o": [
                "bytes",
                "crypto/cipher",
                "encoding/ascii85",
                "encoding/base64",
                "encoding/binary",
                "hash",
                "os",
                "runtime/trace",
                "strings"
            ]
        },
        "io/ioutil": {
            "x": 10,
            "y": 36,
            "i": [
                "bytes",
                "path/filepath",
                "strconv"
            ],
            "o": [
                "archive/tar",
                "archive/zip",
                "crypto/x509",
                "go/parser",
                "net/textproto",
                "text/template"
            ]
        },
        "log": {
            "x": 12,
            "y": 4,
            "i": [
                "fmt"
            ],
            "o": [
                "go/build",
                "log/syslog",
                "net/http",
                "net/mail",
                "testing/iotest"
            ]
        },
        "log/syslog": {
            "x": 13,
            "y": 17,
            "i": [
                "log",
                "net",
                "strings"
            ],
            "o": []
        },
        "math": {
            "x": 4,
            "y": 29,
            "i": [
                "unsafe"
            ],
            "o": [
                "math/cmplx",
                "math/rand",
                "strconv"
            ]
        },
        "math/big": {
            "x": 8,
            "y": 32,
            "i": [
                "bytes",
                "encoding/binary",
                "fmt",
                "math/rand",
                "strings"
            ],
            "o": [
                "crypto/dsa",
                "crypto/elliptic",
                "crypto/rand",
                "encoding/asn1",
                "go/constant"
            ]
        },
        "math/cmplx": {
            "x": 5,
            "y": 32,
            "i": [
                "math"
            ],
            "o": []
        },
        "math/rand": {
            "x": 7,
            "y": 33,
            "i": [
                "math",
                "sync"
            ],
            "o": [
                "internal/trace",
                "math/big",
                "net",
                "testing/quick"
            ]
        },
        "mime": {
            "x": 11,
            "y": 28,
            "i": [
                "bufio",
                "encoding/base64",
                "fmt",
                "sort",
                "strings"
            ],
            "o": [
                "mime/multipart",
                "net/mail"
            ]
        },
        "mime/multipart": {
            "x": 12,
            "y": 36,
            "i": [
                "crypto/rand",
                "mime",
                "mime/quotedprintable",
                "net/textproto"
            ],
            "o": [
                "net/http"
            ]
        },
        "mime/quotedprintable": {
            "x": 11,
            "y": 43,
            "i": [
                "bufio",
                "fmt"
            ],
            "o": [
                "mime/multipart"
            ]
        },
        "net": {
            "x": 10,
            "y": 16,
            "i": [
                "context",
                "internal/nettrace",
                "internal/singleflight",
                "math/rand",
                "sort"
            ],
            "o": [
                "crypto/x509",
                "log/syslog",
                "net/http/httptrace",
                "net/textproto"
            ]
        },
        "net/http": {
            "x": 13,
            "y": 25,
            "i": [
                "compress/gzip",
                "crypto/tls",
                "log",
                "mime/multipart",
                "net/http/httptrace",
                "net/http/internal",
                "net/url",
                "path"
            ],
            "o": [
                "expvar",
                "net/http/cgi",
                "net/http/cookiejar",
                "net/http/httptest",
                "net/http/httputil",
                "net/http/pprof",
                "net/rpc"
            ]
        },
        "net/http/cgi": {
            "x": 14,
            "y": 19,
            "i": [
                "net/http",
                "regexp"
            ],
            "o": [
                "net/http/fcgi"
            ]
        },
        "net/http/cookiejar": {
            "x": 14,
            "y": 37,
            "i": [
                "net/http"
            ],
            "o": []
        },
        "net/http/fcgi": {
            "x": 15,
            "y": 19,
            "i": [
                "net/http/cgi"
            ],
            "o": []
        },
        "net/http/httptest": {
            "x": 14,
            "y": 15,
            "i": [
                "flag",
                "net/http"
            ],
            "o": []
        },
        "net/http/httptrace": {
            "x": 12,
            "y": 0,
            "i": [
                "net"
            ],
            "o": [
                "net/http"
            ]
        },
        "net/http/httputil": {
            "x": 14,
            "y": 13,
            "i": [
                "net/http"
            ],
            "o": []
        },
        "net/http/internal": {
            "x": 12,
            "y": 46,
            "i": [
                "bufio",
                "fmt"
            ],
            "o": [
                "net/http"
            ]
        },
        "net/http/pprof": {
            "x": 14,
            "y": 35,
            "i": [
                "html/template",
                "net/http",
                "runtime/pprof",
                "runtime/trace"
            ],
            "o": []
        },
        "net/internal/socktest": {
            "x": 8,
            "y": 14,
            "i": [
                "fmt"
            ],
            "o": []
        },
        "net/mail": {
            "x": 13,
            "y": 23,
            "i": [
                "log",
                "mime",
                "net/textproto"
            ],
            "o": []
        },
        "net/rpc": {
            "x": 14,
            "y": 30,
            "i": [
                "encoding/gob",
                "html/template",
                "net/http"
            ],
            "o": [
                "net/rpc/jsonrpc"
            ]
        },
        "net/rpc/jsonrpc": {
            "x": 15,
            "y": 30,
            "i": [
                "net/rpc"
            ],
            "o": []
        },
        "net/smtp": {
            "x": 13,
            "y": 42,
            "i": [
                "crypto/tls",
                "net/textproto"
            ],
            "o": []
        },
        "net/textproto": {
            "x": 11,
            "y": 35,
            "i": [
                "bufio",
                "io/ioutil",
                "net"
            ],
            "o": [
                "mime/multipart",
                "net/mail",
                "net/smtp"
            ]
        },
        "net/url": {
            "x": 10,
            "y": 24,
            "i": [
                "bytes",
                "fmt",
                "sort",
                "strings"
            ],
            "o": [
                "net/http",
                "text/template"
            ]
        },
        "os": {
            "x": 6,
            "y": 27,
            "i": [
                "io",
                "time"
            ],
            "o": [
                "fmt",
                "os/signal",
                "path/filepath",
                "runtime/debug",
                "vendor/golang_org/x/net/route"
            ]
        },
        "os/exec": {
            "x": 10,
            "y": 18,
            "i": [
                "bytes",
                "context",
                "path/filepath"
            ],
            "o": [
                "crypto/x509",
                "go/internal/gccgoimporter",
                "internal/testenv",
                "internal/trace"
            ]
        },
        "os/signal": {
            "x": 7,
            "y": 19,
            "i": [
                "os"
            ],
            "o": []
        },
        "os/user": {
            "x": 8,
            "y": 38,
            "i": [
                "fmt",
                "strings"
            ],
            "o": []
        },
        "path": {
            "x": 11,
            "y": 15,
            "i": [
                "strings"
            ],
            "o": [
                "archive/tar",
                "archive/zip",
                "debug/dwarf",
                "go/doc",
                "net/http"
            ]
        },
        "path/filepath": {
            "x": 9,
            "y": 24,
            "i": [
                "os",
                "sort",
                "strings"
            ],
            "o": [
                "go/scanner",
                "io/ioutil",
                "os/exec"
            ]
        },
        "reflect": {
            "x": 6,
            "y": 24,
            "i": [
                "strconv",
                "sync"
            ],
            "o": [
                "encoding/binary",
                "fmt"
            ]
        },
        "regexp": {
            "x": 11,
            "y": 11,
            "i": [
                "regexp/syntax"
            ],
            "o": [
                "go/doc",
                "index/suffixarray",
                "net/http/cgi"
            ]
        },
        "regexp/syntax": {
            "x": 10,
            "y": 20,
            "i": [
                "bytes",
                "sort",
                "strconv",
                "strings"
            ],
            "o": [
                "regexp"
            ]
        },
        "runtime": {
            "x": 2,
            "y": 28,
            "i": [
                "runtime/internal/atomic",
                "runtime/internal/sys"
            ],
            "o": [
                "sync"
            ]
        },
        "runtime/cgo": {
            "x": 1,
            "y": 26,
            "i": [
                "unsafe"
            ],
            "o": []
        },
        "runtime/debug": {
            "x": 13,
            "y": 15,
            "i": [
                "os",
                "sort"
            ],
            "o": [
                "testing"
            ]
        },
        "runtime/internal/atomic": {
            "x": 1,
            "y": 28,
            "i": [
                "unsafe"
            ],
            "o": [
                "runtime"
            ]
        },
        "runtime/internal/sys": {
            "x": 1,
            "y": 30,
            "i": [],
            "o": [
                "runtime"
            ]
        },
        "runtime/pprof": {
            "x": 13,
            "y": 30,
            "i": [
                "bufio",
                "fmt",
                "sort",
                "strings",
                "text/tabwriter"
            ],
            "o": [
                "net/http/pprof",
                "testing"
            ]
        },
        "runtime/race": {
            "x": 0,
            "y": 31,
            "i": [],
            "o": []
        },
        "runtime/trace": {
            "x": 13,
            "y": 46,
            "i": [
                "io"
            ],
            "o": [
                "net/http/pprof",
                "testing"
            ]
        },
        "sort": {
            "x": 8,
            "y": 10,
            "i": [],
            "o": [
                "compress/bzip2",
                "compress/flate",
                "container/heap",
                "database/sql",
                "debug/dwarf",
                "encoding/json",
                "encoding/pem",
                "flag",
                "go/token",
                "mime",
                "net",
                "net/url",
                "path/filepath",
                "regexp/syntax",
                "runtime/debug",
                "runtime/pprof"
            ]
        },
        "strconv": {
            "x": 5,
            "y": 26,
            "i": [
                "errors",
                "math",
                "unicode/utf8"
            ],
            "o": [
                "crypto",
                "crypto/aes",
                "crypto/rc4",
                "encoding/ascii85",
                "encoding/base32",
                "encoding/base64",
                "image",
                "io/ioutil",
                "reflect",
                "regexp/syntax"
            ]
        },
        "strings": {
            "x": 7,
            "y": 29,
            "i": [
                "io",
                "unicode",
                "unicode/utf8"
            ],
            "o": [
                "debug/gosym",
                "encoding/base32",
                "encoding/csv",
                "encoding/json",
                "encoding/pem",
                "encoding/xml",
                "html",
                "log/syslog",
                "math/big",
                "mime",
                "net/url",
                "os/user",
                "path",
                "path/filepath",
                "regexp/syntax",
                "runtime/pprof",
                "testing/quick",
                "text/template/parse",
                "vendor/golang_org/x/net/lex/httplex"
            ]
        },
        "sync": {
            "x": 3,
            "y": 28,
            "i": [
                "internal/race",
                "runtime",
                "sync/atomic"
            ],
            "o": [
                "internal/singleflight",
                "io",
                "math/rand",
                "reflect",
                "syscall"
            ]
        },
        "sync/atomic": {
            "x": 2,
            "y": 30,
            "i": [
                "unsafe"
            ],
            "o": [
                "sync"
            ]
        },
        "syscall": {
            "x": 4,
            "y": 27,
            "i": [
                "errors",
                "sync"
            ],
            "o": [
                "time"
            ]
        },
        "testing": {
            "x": 14,
            "y": 28,
            "i": [
                "flag",
                "runtime/debug",
                "runtime/pprof",
                "runtime/trace"
            ],
            "o": [
                "internal/testenv"
            ]
        },
        "testing/iotest": {
            "x": 13,
            "y": 4,
            "i": [
                "log"
            ],
            "o": []
        },
        "testing/quick": {
            "x": 14,
            "y": 26,
            "i": [
                "flag",
                "math/rand",
                "strings"
            ],
            "o": []
        },
        "text/scanner": {
            "x": 13,
            "y": 32,
            "i": [
                "bytes",
                "fmt"
            ],
            "o": [
                "go/internal/gccgoimporter",
                "go/internal/gcimporter"
            ]
        },
        "text/tabwriter": {
            "x": 12,
            "y": 48,
            "i": [
                "bytes"
            ],
            "o": [
                "go/printer",
                "runtime/pprof"
            ]
        },
        "text/template": {
            "x": 11,
            "y": 40,
            "i": [
                "io/ioutil",
                "net/url",
                "text/template/parse"
            ],
            "o": [
                "go/doc",
                "html/template"
            ]
        },
        "text/template/parse": {
            "x": 10,
            "y": 40,
            "i": [
                "bytes",
                "fmt",
                "strings"
            ],
            "o": [
                "text/template"
            ]
        },
        "time": {
            "x": 5,
            "y": 24,
            "i": [
                "syscall"
            ],
            "o": [
                "os"
            ]
        },
        "unicode": {
            "x": 6,
            "y": 33,
            "i": [],
            "o": [
                "bytes",
                "strings"
            ]
        },
        "unicode/utf16": {
            "x": 11,
            "y": 51,
            "i": [],
            "o": [
                "encoding/json"
            ]
        },
        "unicode/utf8": {
            "x": 4,
            "y": 25,
            "i": [],
            "o": [
                "bytes",
                "strconv",
                "strings"
            ]
        },
        "unsafe": {
            "x": 0,
            "y": 27,
            "i": [],
            "o": [
                "internal/race",
                "math",
                "runtime/cgo",
                "runtime/internal/atomic",
                "sync/atomic"
            ]
        },
        "vendor/golang_org/x/net/http2/hpack": {
            "x": 8,
            "y": 40,
            "i": [
                "bytes",
                "fmt"
            ],
            "o": []
        },
        "vendor/golang_org/x/net/lex/httplex": {
            "x": 8,
            "y": 12,
            "i": [
                "strings"
            ],
            "o": []
        },
        "vendor/golang_org/x/net/route": {
            "x": 7,
            "y": 37,
            "i": [
                "os"
            ],
            "o": []
        }
    }
};
