from exports import Exports
from imports import add_imports_to_linker, Imports
from typing import Tuple
import exports as e
import imports as i
import sys
import wasmtime

class MyImports:    
    def roundtrip_flags1(self, a: i.F1) -> i.F1:
        return a

    def roundtrip_record1(self, a: i.R1) -> i.R1:
        return a

def run(wasm_file: str) -> None:
    store = wasmtime.Store()
    module = wasmtime.Module.from_file(store.engine, wasm_file)
    linker = wasmtime.Linker(store.engine)
    linker.define_wasi()
    wasi = wasmtime.WasiConfig()
    wasi.inherit_stdout()
    wasi.inherit_stderr()
    store.set_wasi(wasi)

    imports = MyImports()
    add_imports_to_linker(linker, store, imports)
    wasm = Exports(store, linker, module)

    wasm.test_imports(store)
    assert(wasm.roundtrip_flags1(store, e.F1.A) == e.F1.A)
    
    r = wasm.roundtrip_record1(store, e.R1(8, e.F1(0)))
    assert(r.a == 8)
    assert(r.b == e.F1(0))

    r = wasm.roundtrip_record1(store, e.R1(a=0, b=e.F1.A | e.F1.B))
    assert(r.a == 0)
    assert(r.b == (e.F1.A | e.F1.B))

if __name__ == '__main__':
    run(sys.argv[1])
