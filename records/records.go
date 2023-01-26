package records

// #include "records.h"
import "C"

// imports
type ImportsF1 uint8

const (
	ImportsF1_A ImportsF1 = 1 << iota
	ImportsF1_B
)

type ImportsR1 struct {
	A uint8
	B ImportsR2
}

type ImportsR2 struct {
	C uint8
	D uint8
}

func ImportsRoundtripFlags1(a ImportsF1) ImportsF1 {

	var lower_a C.imports_f1_t
	lower_a = C.uint8_t(a)

	result := C.imports_roundtrip_flags1(lower_a)
	var lift_result ImportsF1

	lift_result = ImportsF1(result)
	return lift_result
}

func ImportsRoundtripRecord1(a ImportsR1) ImportsR1 {

	var lower_a C.imports_r1_t
	lower_a = C.imports_r1_t{}
	lower_a.a = C.uint8_t(a.A)
	lower_a.b = C.imports_r2_t{}
	lower_a.b.c = C.uint8_t(a.B.C)
	lower_a.b.d = C.uint8_t(a.B.D)

	var result C.imports_r1_t
	C.imports_roundtrip_record1(&lower_a, &result)
	var lift_result ImportsR1

	var lift_result_A uint8
	lift_result_A = uint8(result.a)
	lift_result.A = lift_result_A
	var lift_result_B ImportsR2

	var lift_result_B_C uint8
	lift_result_B_C = uint8(result.b.c)
	lift_result_B.C = lift_result_B_C
	var lift_result_B_D uint8
	lift_result_B_D = uint8(result.b.d)
	lift_result_B.D = lift_result_B_D
	lift_result.B = lift_result_B
	return lift_result
}

func ImportsTuple0(a ImportsTuple0T) ImportsTuple0T {

	var lower_a C.imports_tuple0_t
	lower_a = C.imports_tuple0_t{}

	var result C.imports_tuple0_t
	C.imports_tuple0(&lower_a, &result)
	var lift_result ImportsTuple0T

	return lift_result
}

func ImportsTuple1(a ImportsTuple1U8T) ImportsTuple1U8T {

	var lower_a C.imports_tuple1_u8_t
	lower_a = C.imports_tuple1_u8_t{}
	lower_a.f0 = C.uint8_t(a.F0)

	var result C.imports_tuple1_u8_t
	C.imports_tuple1(&lower_a, &result)
	var lift_result ImportsTuple1U8T

	var lift_result_F0 uint8
	lift_result_F0 = uint8(result.f0)
	lift_result.F0 = lift_result_F0
	return lift_result
}

type ImportsTuple0T struct {
}

type ImportsTuple1U8T struct {
	F0 uint8
}

// default records
type RecordsF1 uint8

const (
	RecordsF1_A RecordsF1 = 1 << iota
	RecordsF1_B
)

type RecordsR1 struct {
	A uint8
	B RecordsF1
}

type RecordsTuple0T struct {
}

type RecordsTuple1U8T struct {
	F0 uint8
}

var records Records = nil

func SetRecords(i Records) {
	records = i
}

type Records interface {
	TestImports()
	RoundtripFlags1(a RecordsF1) RecordsF1
	RoundtripRecord1(a RecordsR1) RecordsR1
	Tuple0(a RecordsTuple0T) RecordsTuple0T
	Tuple1(a RecordsTuple1U8T) RecordsTuple1U8T
}

//export records_test_imports
func RecordsTestImports() {
	records.TestImports()

}

//export records_roundtrip_flags1
func RecordsRoundtripFlags1(a C.records_f1_t) C.records_f1_t {
	var lift_a RecordsF1

	lift_a = RecordsF1(a)
	result := records.RoundtripFlags1(lift_a)

	var lower_result C.records_f1_t
	lower_result = C.uint8_t(result)

	return lower_result

}

//export records_roundtrip_record1
func RecordsRoundtripRecord1(a *C.records_r1_t, ret *C.records_r1_t) {
	var lift_a RecordsR1

	var lift_a_A uint8
	lift_a_A = uint8(a.a)
	lift_a.A = lift_a_A
	var lift_a_B RecordsF1

	lift_a_B = RecordsF1(a.b)
	lift_a.B = lift_a_B
	result := records.RoundtripRecord1(lift_a)

	var lower_result C.records_r1_t
	lower_result = C.records_r1_t{}
	lower_result.a = C.uint8_t(result.A)
	lower_result.b = C.uint8_t(result.B)

	*ret = lower_result

}

//export records_tuple0
func RecordsTuple0(a *C.records_tuple0_t, ret *C.records_tuple0_t) {
	var lift_a RecordsTuple0T

	records.Tuple0(lift_a)

	var lower_result C.records_tuple0_t
	lower_result = C.records_tuple0_t{}

	*ret = lower_result

}

//export records_tuple1
func RecordsTuple1(a *C.records_tuple1_u8_t, ret *C.records_tuple1_u8_t) {
	var lift_a RecordsTuple1U8T

	var lift_a_F0 uint8
	lift_a_F0 = uint8(a.f0)
	lift_a.F0 = lift_a_F0
	result := records.Tuple1(lift_a)

	var lower_result C.records_tuple1_u8_t
	lower_result = C.records_tuple1_u8_t{}
	lower_result.f0 = C.uint8_t(result.F0)

	*ret = lower_result

}
