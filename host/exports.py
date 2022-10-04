from abc import abstractmethod
import ctypes
from dataclasses import dataclass
from enum import Flag, auto
from typing import Any, Tuple
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

def _load(ty: Any, mem: wasmtime.Memory, store: wasmtime.Storelike, base: int, offset: int) -> Any:
    ptr = (base & 0xffffffff) + offset
    if ptr + ctypes.sizeof(ty) > mem.data_len(store):
        raise IndexError('out-of-bounds store')
    raw_base = mem.data_ptr(store)
    c_ptr = ctypes.POINTER(ty)(
        ty.from_address(ctypes.addressof(raw_base.contents) + ptr)
    )
    return c_ptr[0]
class F1(Flag):
    A = auto()
    B = auto()

@dataclass
class R1:
    a: int
    b: F1

class Exports:
    instance: wasmtime.Instance
    _memory: wasmtime.Memory
    _roundtrip_flags1: wasmtime.Func
    _roundtrip_record1: wasmtime.Func
    _test_imports: wasmtime.Func
    _tuple0: wasmtime.Func
    _tuple1: wasmtime.Func
    def __init__(self, store: wasmtime.Store, linker: wasmtime.Linker, module: wasmtime.Module):
        self.instance = linker.instantiate(store, module)
        exports = self.instance.exports(store)
        
        memory = exports['memory']
        assert(isinstance(memory, wasmtime.Memory))
        self._memory = memory
        
        roundtrip_flags1 = exports['roundtrip-flags1: func(a: flags { a, b }) -> flags { a, b }']
        assert(isinstance(roundtrip_flags1, wasmtime.Func))
        self._roundtrip_flags1 = roundtrip_flags1
        
        roundtrip_record1 = exports['roundtrip-record1: func(a: record { a: u8, b: flags { a, b } }) -> record { a: u8, b: flags { a, b } }']
        assert(isinstance(roundtrip_record1, wasmtime.Func))
        self._roundtrip_record1 = roundtrip_record1
        
        test_imports = exports['test-imports: func() -> ()']
        assert(isinstance(test_imports, wasmtime.Func))
        self._test_imports = test_imports
        
        tuple0 = exports['tuple0: func(a: tuple<>) -> tuple<>']
        assert(isinstance(tuple0, wasmtime.Func))
        self._tuple0 = tuple0
        
        tuple1 = exports['tuple1: func(a: tuple<u8>) -> tuple<u8>']
        assert(isinstance(tuple1, wasmtime.Func))
        self._tuple1 = tuple1
    def test_imports(self, caller: wasmtime.Store) -> None:
        self._test_imports(caller)
    def roundtrip_flags1(self, caller: wasmtime.Store, a: F1) -> F1:
        ret = self._roundtrip_flags1(caller, (a).value)
        assert(isinstance(ret, int))
        return F1(ret)
    def roundtrip_record1(self, caller: wasmtime.Store, a: R1) -> R1:
        memory = self._memory;
        record = a
        field = record.a
        field0 = record.b
        ret = self._roundtrip_record1(caller, _clamp(field, 0, 255), (field0).value)
        assert(isinstance(ret, int))
        load = _load(ctypes.c_uint8, memory, caller, ret, 0)
        load1 = _load(ctypes.c_uint8, memory, caller, ret, 1)
        return R1(_clamp(load, 0, 255), F1(load1))
    def tuple0(self, caller: wasmtime.Store, a: None) -> None:
        self._tuple0(caller)
        return None
    def tuple1(self, caller: wasmtime.Store, a: Tuple[int]) -> Tuple[int]:
        (tuplei,) = a
        ret = self._tuple1(caller, _clamp(tuplei, 0, 255))
        assert(isinstance(ret, int))
        return (_clamp(ret, 0, 255),)
