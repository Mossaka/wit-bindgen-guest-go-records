package exports

// #include "exports.h"
import "C"

var exports Exports = nil

type Exports interface {
	TestImports()
	RoundtripFlags1(a F1) F1
	RoundtripRecord1(a R1) R1
}

func SetExports(e Exports) {
	exports = e
}

type F1 uint8

const (
	A F1 = 1 << iota
	B
)

type R1 struct {
	A uint8
	B F1
}

//export exports_test_imports
func TestImports() {
	exports.TestImports()
}

//export exports_roundtrip_flags1
func RoundtripFlags1(a C.uint8_t) C.uint8_t {
	return C.uint8_t(exports.RoundtripFlags1(F1(a)))
}

//export exports_roundtrip_record1
func RoundtripRecord1(param *C.exports_r1_t, ret *C.exports_r1_t) {
	a := R1{
		A: uint8(param.a),
		B: F1(param.b),
	}
	b := exports.RoundtripRecord1(a)
	ret.a = C.uint8_t(b.A)
	ret.b = C.uint8_t(b.B)
}
