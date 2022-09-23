from abc import abstractmethod
import ctypes
from dataclasses import dataclass
from enum import Flag, auto
from typing import Any, cast
import wasmtime

try:
    from typing import Protocol
except ImportError:
    class Protocol: # type: ignore
        pass


def _clamp(i: int, min: int, max: int) -> int:
    if i < min or i > max:
        raise OverflowError(f'must be between {min} and {max}')
    return i

def _store(ty: Any, mem: wasmtime.Memory, store: wasmtime.Storelike, base: int, offset: int, val: Any) -> None:
    ptr = (base & 0xffffffff) + offset
    if ptr + ctypes.sizeof(ty) > mem.data_len(store):
        raise IndexError('out-of-bounds store')
    raw_base = mem.data_ptr(store)
    c_ptr = ctypes.POINTER(ty)(
        ty.from_address(ctypes.addressof(raw_base.contents) + ptr)
    )
    c_ptr[0] = val
class F1(Flag):
    A = auto()
    B = auto()

@dataclass
class R1:
    a: int
    b: F1

class Imports(Protocol):
    @abstractmethod
    def roundtrip_flags1(self, a: F1) -> F1:
        raise NotImplementedError
    @abstractmethod
    def roundtrip_record1(self, a: R1) -> R1:
        raise NotImplementedError

def add_imports_to_linker(linker: wasmtime.Linker, store: wasmtime.Store, host: Imports) -> None:
    ty = wasmtime.FuncType([wasmtime.ValType.i32()], [wasmtime.ValType.i32()])
    def roundtrip_flags1(caller: wasmtime.Caller, arg0: int) -> int:
        ret = host.roundtrip_flags1(F1(arg0))
        return (ret).value
    linker.define('imports', 'roundtrip-flags1: func(a: flags { a, b }) -> flags { a, b }', wasmtime.Func(store, ty, roundtrip_flags1, access_caller = True))
    ty = wasmtime.FuncType([wasmtime.ValType.i32(), wasmtime.ValType.i32(), wasmtime.ValType.i32()], [])
    def roundtrip_record1(caller: wasmtime.Caller, arg0: int, arg1: int, arg2: int) -> None:
        m = caller["memory"]
        assert(isinstance(m, wasmtime.Memory))
        memory = cast(wasmtime.Memory, m)
        ret = host.roundtrip_record1(R1(_clamp(arg0, 0, 255), F1(arg1)))
        record = ret
        field = record.a
        field0 = record.b
        _store(ctypes.c_uint8, memory, caller, arg2, 0, _clamp(field, 0, 255))
        _store(ctypes.c_uint8, memory, caller, arg2, 1, (field0).value)
    linker.define('imports', 'roundtrip-record1: func(a: record { a: u8, b: flags { a, b } }) -> record { a: u8, b: flags { a, b } }', wasmtime.Func(store, ty, roundtrip_record1, access_caller = True))
