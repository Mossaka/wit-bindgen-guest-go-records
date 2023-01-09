from host import Descriptor, Records, F1, R1, RecordsImports, WasiStream, imports, wasi_filesystem, wasi_logging, wasi_poll
from typing import Tuple
import sys
import wasmtime
from host.imports.wasi_filesystem import Errno, Filesize, Size

from host.imports.wasi_logging import Level
from host.imports.wasi_poll import StreamError
from host.types import Result

class MyImports(imports.Imports):
    def roundtrip_flags1(self, a: F1) -> F1:
        return a

    def roundtrip_record1(self, a: R1) -> R1:
        return a

    def tuple0(self, a: None) -> None:
        pass

    def tuple1(self, a: Tuple[int]) -> Tuple[int]:
        return (a[0],)

class Logging(wasi_logging.WasiLogging):
    def log(self, level: Level, context: str, message: str) -> None:
        print(f"{message}")

class Filesystem(wasi_filesystem.WasiFilesystem):
    def write_via_stream(self, fd: Descriptor, offset: Filesize) -> Result[WasiStream, Errno]:
        raise NotImplementedError

class Poll(wasi_poll.WasiPoll):
    def write_stream(self, stream: WasiStream, buf: bytes) -> Result[Size, StreamError]:
        raise NotImplementedError
 
def run() -> None:
    store = wasmtime.Store()
    wasm = Records(store, RecordsImports(MyImports(), Logging(), Filesystem(), Poll()))
    
    wasm.test_imports(store)
    assert(wasm.roundtrip_flags1(store, F1.A) == F1.A)
    
    r = wasm.roundtrip_record1(store, R1(8, F1(0)))
    assert(r.a == 8)
    assert(r.b == F1(0))

    r = wasm.roundtrip_record1(store, R1(a=0, b=F1.A | F1.B))
    assert(r.a == 0)
    assert(r.b == (F1.A | F1.B))

    wasm.tuple0(store, None)
    assert(wasm.tuple1(store, (1,)) == (1,))
if __name__ == '__main__':
    run()
