#include "records.h"


__attribute__((import_module("imports"), import_name("roundtrip-flags1")))
int32_t __wasm_import_imports_roundtrip_flags1(int32_t);

__attribute__((import_module("imports"), import_name("roundtrip-record1")))
void __wasm_import_imports_roundtrip_record1(int32_t, int32_t, int32_t);

__attribute__((import_module("imports"), import_name("tuple0")))
void __wasm_import_imports_tuple0(void);

__attribute__((import_module("imports"), import_name("tuple1")))
int32_t __wasm_import_imports_tuple1(int32_t);

__attribute__((weak, export_name("cabi_realloc")))
void *cabi_realloc(void *ptr, size_t orig_size, size_t org_align, size_t new_size) {
  void *ret = realloc(ptr, new_size);
  if (!ret) abort();
  return ret;
}

// Component Adapters

__attribute__((aligned(1)))
static uint8_t RET_AREA[2];

imports_f1_t imports_roundtrip_flags1(imports_f1_t a) {
  int32_t ret = __wasm_import_imports_roundtrip_flags1(a);
  return ret;
}

void imports_roundtrip_record1(imports_r1_t *a, imports_r1_t *ret) {
  __attribute__((aligned(1)))
  uint8_t ret_area[2];
  int32_t ptr = (int32_t) &ret_area;
  __wasm_import_imports_roundtrip_record1((int32_t) ((*a).a), (*a).b, ptr);
  *ret = (imports_r1_t) {
    (uint8_t) ((int32_t) (*((uint8_t*) (ptr + 0)))),
    (int32_t) (*((uint8_t*) (ptr + 1))),
  };
}

void imports_tuple0(imports_tuple0_t *a, imports_tuple0_t *ret) {
  __wasm_import_imports_tuple0();
  *ret = (imports_tuple0_t) {
  };
}

void imports_tuple1(imports_tuple1_u8_t *a, imports_tuple1_u8_t *ret) {
  int32_t ret0 = __wasm_import_imports_tuple1((int32_t) ((*a).f0));
  *ret = (imports_tuple1_u8_t) {
    (uint8_t) (ret0),
  };
}

__attribute__((export_name("test-imports")))
void __wasm_export_records_test_imports(void) {
  records_test_imports();
}

__attribute__((export_name("roundtrip-flags1")))
int32_t __wasm_export_records_roundtrip_flags1(int32_t arg) {
  records_f1_t ret = records_roundtrip_flags1(arg);
  return ret;
}

__attribute__((export_name("roundtrip-record1")))
int32_t __wasm_export_records_roundtrip_record1(int32_t arg, int32_t arg0) {
  records_r1_t arg1 = (records_r1_t) {
    (uint8_t) (arg),
    arg0,
  };
  records_r1_t ret;
  records_roundtrip_record1(&arg1, &ret);
  int32_t ptr = (int32_t) &RET_AREA;
  *((int8_t*)(ptr + 0)) = (int32_t) ((ret).a);
  *((int8_t*)(ptr + 1)) = (ret).b;
  return ptr;
}

__attribute__((export_name("tuple0")))
void __wasm_export_records_tuple0(void) {
  records_tuple0_t arg = (records_tuple0_t) {
  };
  records_tuple0_t ret;
  records_tuple0(&arg, &ret);
}

__attribute__((export_name("tuple1")))
int32_t __wasm_export_records_tuple1(int32_t arg) {
  records_tuple1_u8_t arg0 = (records_tuple1_u8_t) {
    (uint8_t) (arg),
  };
  records_tuple1_u8_t ret;
  records_tuple1(&arg0, &ret);
  return (int32_t) ((ret).f0);
}

extern void __component_type_object_force_link_records(void);
void __component_type_object_force_link_records_public_use_in_this_compilation_unit(void) {
  __component_type_object_force_link_records();
}
