package imports

//#include "imports.h"
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
