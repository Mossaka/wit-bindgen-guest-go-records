#ifndef __BINDINGS_EXPORTS_H
#define __BINDINGS_EXPORTS_H
#ifdef __cplusplus
extern "C"
{
  #endif
  
  #include <stdint.h>
  #include <stdbool.h>
  typedef uint8_t exports_f1_t;
  #define EXPORTS_F1_A (1 << 0)
  #define EXPORTS_F1_B (1 << 1)
  typedef struct {
    uint8_t a;
    exports_f1_t b;
  } exports_r1_t;
  typedef struct {
  } exports_tuple0_t;
  typedef struct {
    uint8_t f0;
  } exports_tuple1_u8_t;
  void exports_test_imports(void);
  exports_f1_t exports_roundtrip_flags1(exports_f1_t a);
  void exports_roundtrip_record1(exports_r1_t *a, exports_r1_t *ret0);
  void exports_tuple0(exports_tuple0_t *a, exports_tuple0_t *ret0);
  void exports_tuple1(exports_tuple1_u8_t *a, exports_tuple1_u8_t *ret0);
  #ifdef __cplusplus
}
#endif
#endif
