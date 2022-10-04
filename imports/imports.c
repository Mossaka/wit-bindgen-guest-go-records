#include <stdlib.h>
#include <imports.h>

__attribute__((weak, export_name("cabi_realloc")))
void *cabi_realloc(
void *ptr,
size_t orig_size,
size_t org_align,
size_t new_size
) {
  void *ret = realloc(ptr, new_size);
  if (!ret)
  abort();
  return ret;
}

__attribute__((aligned(1)))
static uint8_t RET_AREA[2];
__attribute__((import_module("imports"), import_name("roundtrip-flags1: func(a: flags { a, b }) -> flags { a, b }")))
int32_t __wasm_import_imports_roundtrip_flags1(int32_t);
imports_f1_t imports_roundtrip_flags1(imports_f1_t a) {
  int32_t ret = __wasm_import_imports_roundtrip_flags1(a);
  return ret;
}
__attribute__((import_module("imports"), import_name("roundtrip-record1: func(a: record { a: u8, b: flags { a, b } }) -> record { a: u8, b: flags { a, b } }")))
void __wasm_import_imports_roundtrip_record1(int32_t, int32_t, int32_t);
void imports_roundtrip_record1(imports_r1_t *a, imports_r1_t *ret0) {
  int32_t ptr = (int32_t) &RET_AREA;
  __wasm_import_imports_roundtrip_record1((int32_t) ((*a).a), (*a).b, ptr);
  *ret0 = (imports_r1_t) {
    (uint8_t) ((int32_t) (*((uint8_t*) (ptr + 0)))),
    (int32_t) (*((uint8_t*) (ptr + 1))),
  };
}
__attribute__((import_module("imports"), import_name("tuple0: func(a: tuple<>) -> tuple<>")))
void __wasm_import_imports_tuple0(void);
void imports_tuple0(imports_tuple0_t *a, imports_tuple0_t *ret0) {
  __wasm_import_imports_tuple0();
  *ret0 = (imports_tuple0_t) {
  };
}
__attribute__((import_module("imports"), import_name("tuple1: func(a: tuple<u8>) -> tuple<u8>")))
int32_t __wasm_import_imports_tuple1(int32_t);
void imports_tuple1(imports_tuple1_u8_t *a, imports_tuple1_u8_t *ret0) {
  int32_t ret = __wasm_import_imports_tuple1((int32_t) ((*a).f0));
  *ret0 = (imports_tuple1_u8_t) {
    (uint8_t) (ret),
  };
}
