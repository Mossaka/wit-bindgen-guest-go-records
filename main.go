package main

import (
	"github.com/mossaka/go-wit-bindgen-records/exports"
	"github.com/mossaka/go-wit-bindgen-records/imports"
)

func init() {
	exports.SetExports(RecordsExportImpl{})
}

type RecordsExportImpl struct{}

func (i RecordsExportImpl) TestImports() {
	imports.RoundtripFlags1(imports.F1_A)
	imports.RoundtripRecord1(imports.R1{})
	imports.Tuple0(imports.Tuple0T{})
	imports.Tuple1(imports.Tuple1Uint8T{
		F0: 0,
	})
}

func (i RecordsExportImpl) RoundtripFlags1(a exports.F1) exports.F1 {
	return a
}

func (i RecordsExportImpl) RoundtripRecord1(a exports.R1) exports.R1 {
	return a
}

func (i RecordsExportImpl) Tuple0(a exports.Tuple0T) exports.Tuple0T {
	return a
}

func (i RecordsExportImpl) Tuple1(a exports.Tuple1Uint8T) exports.Tuple1Uint8T {
	return exports.Tuple1Uint8T{F0: a.F0}
}

func main() {
}
