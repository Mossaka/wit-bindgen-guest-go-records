#ifndef __BINDINGS_IMPORTS_H
#define __BINDINGS_IMPORTS_H
#ifdef __cplusplus
extern "C"
{
  #endif
  
  #include <stdint.h>
  #include <stdbool.h>
  typedef uint8_t imports_f1_t;
  #define IMPORTS_F1_A (1 << 0)
  #define IMPORTS_F1_B (1 << 1)
  typedef struct {
    uint8_t a;
    imports_f1_t b;
  } imports_r1_t;
  imports_f1_t imports_roundtrip_flags1(imports_f1_t a);
  void imports_roundtrip_record1(imports_r1_t *a, imports_r1_t *ret0);
  #ifdef __cplusplus
}
#endif
#endif
