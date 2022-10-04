package exports

// #include "exports.h"
import "C"

var exports Exports = nil

type Exports interface {
	TestImports()
	RoundtripFlags1(a F1) F1
	RoundtripRecord1(a R1) R1
	Tuple0(a Tuple0T) Tuple0T
	Tuple1(a Tuple1Uint8T) Tuple1Uint8T
}

func SetExports(e Exports) {
	exports = e
}

type F1 uint8

const (
	F1_A F1 = 1 << iota
	F1_B
)

type R1 struct {
	A uint8
	B F1
}

type Tuple0T struct{}
type Tuple1Uint8T struct {
	F0 uint8
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

//export exports_tuple0
func Tuple0(a *C.exports_tuple0_t, ret *C.exports_tuple0_t) {
	param := Tuple0T{}
	exports.Tuple0(param)
}

//export exports_tuple1
func Tuple1(a *C.exports_tuple1_u8_t, ret *C.exports_tuple1_u8_t) {
	param := Tuple1Uint8T{
		F0: uint8(a.f0),
	}
	b := exports.Tuple1(param)
	ret.f0 = C.uint8_t(b.F0)
}

// FIXME: how do I resolve name conflicts?

// FIXME: The problem with adding generic types to tuples/result/option types is that
// the type will be bubbled up to the top level Exports interface. This means that
// the interface will have to be generic.
