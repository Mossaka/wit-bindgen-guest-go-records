package records

//#include "records.h"
import "C"

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

// Exports
var exports Exports = nil

type Exports interface {
	RecordsTestImports()
	RecordsRoundtripFlags1(a F1) F1
	RecordsRoundtripRecord1(a R1) R1
	RecordsTuple0(a Tuple0T) Tuple0T
	RecordsTuple1(a Tuple1Uint8T) Tuple1Uint8T
}

func SetExports(e Exports) {
	exports = e
}

//export records_test_imports
func RecordsTestImports() {
	exports.RecordsTestImports()
}

//export records_roundtrip_flags1
func RecordsRoundtripFlags1(a C.uint8_t) C.uint8_t {
	return C.uint8_t(exports.RecordsRoundtripFlags1(F1(a)))
}

//export records_roundtrip_record1
func RecordsRoundtripRecord1(param *C.records_r1_t, ret *C.records_r1_t) {
	a := R1{
		A: uint8(param.a),
		B: F1(param.b),
	}
	b := exports.RecordsRoundtripRecord1(a)
	ret.a = C.uint8_t(b.A)
	ret.b = C.uint8_t(b.B)
}

//export records_tuple0
func RecordsTuple0(a *C.records_tuple0_t, ret *C.records_tuple0_t) {
	param := Tuple0T{}
	exports.RecordsTuple0(param)
}

//export records_tuple1
func RecordsTuple1(a *C.records_tuple1_u8_t, ret *C.records_tuple1_u8_t) {
	param := Tuple1Uint8T{
		F0: uint8(a.f0),
	}
	b := exports.RecordsTuple1(param)
	ret.f0 = C.uint8_t(b.F0)
}

// FIXME: how do I resolve name conflicts?

// FIXME: The problem with adding generic types to tuples/result/option types is that
// the type will be bubbled up to the top level Exports interface. This means that
// the interface will have to be generic.

// Imports
func RoundtripFlags1(a F1) F1 {
	return F1(C.imports_roundtrip_flags1(C.uint8_t(a)))
}

func RoundtripRecord1(a R1) R1 {
	param := C.imports_r1_t{
		a: C.uint8_t(a.A),
		b: C.uint8_t(a.B),
	}
	ret := C.imports_r1_t{}
	C.imports_roundtrip_record1(&param, &ret)
	return R1{
		A: uint8(ret.a),
		B: F1(ret.b),
	}
}

func Tuple0(a Tuple0T) Tuple0T {
	return Tuple0T{}
}

func Tuple1(a Tuple1Uint8T) Tuple1Uint8T {
	param := C.imports_tuple1_u8_t{
		f0: C.uint8_t(a.F0),
	}
	ret := C.imports_tuple1_u8_t{}
	C.imports_tuple1(&param, &ret)
	return Tuple1Uint8T{
		F0: uint8(ret.f0),
	}
}
