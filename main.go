package main

import (
	"github.com/mossaka/go-wit-bindgen-records/records"
)

func init() {
	records.SetRecords(RecordsExportImpl{})
}

type RecordsExportImpl struct{}

func (i RecordsExportImpl) TestImports() {
	records.ImportsRoundtripFlags1(records.ImportsF1_A)
	records.ImportsRoundtripRecord1(records.ImportsR1{})
	records.ImportsTuple0(records.ImportsTuple0T{})
	records.ImportsTuple1(records.ImportsTuple1U8T{
		F0: 0,
	})
}

func (i RecordsExportImpl) RoundtripFlags1(a records.RecordsF1) records.RecordsF1 {
	return a
}

func (i RecordsExportImpl) RoundtripRecord1(a records.RecordsR1) records.RecordsR1 {
	return a
}

func (i RecordsExportImpl) Tuple0(a records.RecordsTuple0T) records.RecordsTuple0T {
	return a
}

func (i RecordsExportImpl) Tuple1(a records.RecordsTuple1U8T) records.RecordsTuple1U8T {
	return records.RecordsTuple1U8T{F0: a.F0}
}

func main() {
}
