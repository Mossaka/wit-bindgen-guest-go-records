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
}

func (i RecordsExportImpl) RoundtripFlags1(a exports.F1) exports.F1 {
	return a
}

func (i RecordsExportImpl) RoundtripRecord1(a exports.R1) exports.R1 {
	return a
}

func main() {
}
