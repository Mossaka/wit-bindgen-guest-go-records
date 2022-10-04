#include <stdlib.h>
#include <exports.h>

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
__attribute__((export_name("test-imports: func() -> ()")))
void __wasm_export_exports_test_imports(void) {
  exports_test_imports();
}
__attribute__((export_name("roundtrip-flags1: func(a: flags { a, b }) -> flags { a, b }")))
int32_t __wasm_export_exports_roundtrip_flags1(int32_t arg) {
  exports_f1_t ret = exports_roundtrip_flags1(arg);
  return ret;
}
__attribute__((export_name("roundtrip-record1: func(a: record { a: u8, b: flags { a, b } }) -> record { a: u8, b: flags { a, b } }")))
int32_t __wasm_export_exports_roundtrip_record1(int32_t arg, int32_t arg0) {
  exports_r1_t arg1 = (exports_r1_t) {
    (uint8_t) (arg),
    arg0,
  };
  exports_r1_t ret;
  exports_roundtrip_record1(&arg1, &ret);
  int32_t ptr = (int32_t) &RET_AREA;
  *((int8_t*)(ptr + 0)) = (int32_t) ((ret).a);
  *((int8_t*)(ptr + 1)) = (ret).b;
  return ptr;
}
__attribute__((export_name("tuple0: func(a: tuple<>) -> tuple<>")))
void __wasm_export_exports_tuple0(void) {
  exports_tuple0_t arg = (exports_tuple0_t) {
  };
  exports_tuple0_t ret;
  exports_tuple0(&arg, &ret);
}
__attribute__((export_name("tuple1: func(a: tuple<u8>) -> tuple<u8>")))
int32_t __wasm_export_exports_tuple1(int32_t arg) {
  exports_tuple1_u8_t arg0 = (exports_tuple1_u8_t) {
    (uint8_t) (arg),
  };
  exports_tuple1_u8_t ret;
  exports_tuple1(&arg0, &ret);
  return (int32_t) ((ret).f0);
}
