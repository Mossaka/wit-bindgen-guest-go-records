#[allow(clippy::all)]
mod imports {
  wit_bindgen_guest_rust::bitflags::bitflags! {
    pub struct F1: u8 {
      const A = 1 << 0;
      const B = 1 << 1;
    }
  }
  impl F1 {
        /// Convert from a raw integer, preserving any unknown bits. See
        /// <https://github.com/bitflags/bitflags/issues/263#issuecomment-957088321>
        pub fn from_bits_preserve(bits: u8) -> Self {
              Self { bits }
        }
  }
  #[repr(C)]
  #[derive(Copy, Clone)]
  pub struct R1 {
    pub a: u8,
    pub b: F1,
  }
  impl core::fmt::Debug for R1 {
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
      f.debug_struct("R1").field("a", &self.a).field("b", &self.b).finish()}
  }
  pub fn roundtrip_flags1(a: F1,) -> F1{
    unsafe {
      let flags0 = a;
      #[link(wasm_import_module = "imports")]
      extern "C" {
        #[cfg_attr(target_arch = "wasm32", link_name = "roundtrip-flags1: func(a: flags { a, b }) -> flags { a, b }")]
        #[cfg_attr(not(target_arch = "wasm32"), link_name = "imports_roundtrip-flags1: func(a: flags { a, b }) -> flags { a, b }")]
        fn wit_import(_: i32, ) -> i32;
      }
      let ret = wit_import((flags0.bits() >> 0) as i32);
      F1::empty() | F1::from_bits_preserve(((ret as u8) << 0) as _)
    }
  }
  pub fn roundtrip_record1(a: R1,) -> R1{
    unsafe {
      let R1{ a:a0, b:b0, } = a;
      let flags1 = b0;
      let ptr2 = __IMPORTS_RET_AREA.0.as_mut_ptr() as i32;
      #[link(wasm_import_module = "imports")]
      extern "C" {
        #[cfg_attr(target_arch = "wasm32", link_name = "roundtrip-record1: func(a: record { a: u8, b: flags { a, b } }) -> record { a: u8, b: flags { a, b } }")]
        #[cfg_attr(not(target_arch = "wasm32"), link_name = "imports_roundtrip-record1: func(a: record { a: u8, b: flags { a, b } }) -> record { a: u8, b: flags { a, b } }")]
        fn wit_import(_: i32, _: i32, _: i32, );
      }
      wit_import(wit_bindgen_guest_rust::rt::as_i32(a0), (flags1.bits() >> 0) as i32, ptr2);
      R1{a:i32::from(*((ptr2 + 0) as *const u8)) as u8, b:F1::empty() | F1::from_bits_preserve(((i32::from(*((ptr2 + 1) as *const u8)) as u8) << 0) as _), }
    }
  }
  
  #[repr(align(1))]
  struct __ImportsRetArea([u8; 2]);
  static mut __IMPORTS_RET_AREA: __ImportsRetArea = __ImportsRetArea([0; 2]);
}
