package main

import (
	"github.com/mossaka/go-wit-bindgen-records/records"
)

func init() {
	records.SetExports(RecordsExportImpl{})
}

type RecordsExportImpl struct{}

func (i RecordsExportImpl) RecordsTestImports() {
	records.RoundtripFlags1(records.F1_A)
	records.RoundtripRecord1(records.R1{})
	records.Tuple0(records.Tuple0T{})
	records.Tuple1(records.Tuple1Uint8T{
		F0: 0,
	})
}

func (i RecordsExportImpl) RecordsRoundtripFlags1(a records.F1) records.F1 {
	return a
}

func (i RecordsExportImpl) RecordsRoundtripRecord1(a records.R1) records.R1 {
	return a
}

func (i RecordsExportImpl) RecordsTuple0(a records.Tuple0T) records.Tuple0T {
	return a
}

func (i RecordsExportImpl) RecordsTuple1(a records.Tuple1Uint8T) records.Tuple1Uint8T {
	return records.Tuple1Uint8T{F0: a.F0}
}

func main() {
}
