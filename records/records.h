#ifndef __BINDINGS_RECORDS_H
#define __BINDINGS_RECORDS_H
#ifdef __cplusplus
extern "C" {
#endif

#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>

typedef uint8_t imports_f1_t;

#define IMPORTS_F1_A (1 << 0)
#define IMPORTS_F1_B (1 << 1)

typedef struct {
  uint8_t a;
  imports_f1_t b;
} imports_r1_t;

typedef struct {
} imports_tuple0_t;

typedef struct {
  uint8_t f0;
} imports_tuple1_u8_t;

typedef uint8_t records_f1_t;

#define RECORDS_F1_A (1 << 0)
#define RECORDS_F1_B (1 << 1)

typedef struct {
  uint8_t a;
  records_f1_t b;
} records_r1_t;

typedef struct {
} records_tuple0_t;

typedef struct {
  uint8_t f0;
} records_tuple1_u8_t;

// Imported Functions from `imports`
imports_f1_t imports_roundtrip_flags1(imports_f1_t a);
void imports_roundtrip_record1(imports_r1_t *a, imports_r1_t *ret);
void imports_tuple0(imports_tuple0_t *a, imports_tuple0_t *ret);
void imports_tuple1(imports_tuple1_u8_t *a, imports_tuple1_u8_t *ret);

// Exported Functions from `records`
void records_test_imports(void);
records_f1_t records_roundtrip_flags1(records_f1_t a);
void records_roundtrip_record1(records_r1_t *a, records_r1_t *ret);
void records_tuple0(records_tuple0_t *a, records_tuple0_t *ret);
void records_tuple1(records_tuple1_u8_t *a, records_tuple1_u8_t *ret);

#ifdef __cplusplus
}
#endif
#endif
