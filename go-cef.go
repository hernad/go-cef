package main

/*
  #cgo CXXFLAGS: -DGOLANG
  #cgo CXXFLAGS: -I/home/bringout/dev/cef/cef_bin
  #cgo CXXFLAGS: -I/home/bringout/dev/cef/cef_bin/include
  #cgo LDFLAGS: -L/home/bringout/dev/cef/cef_bin/Release -lcef
  #cgo LDFLAGS: -L/home/bringout/dev/cef/cef_bin/out/Release/obj.target -lcef_dll_wrapper
  #cgo LDFLAGS: -lstdc++ -lm -ldl
  #cgo pkg-config: gtk+-2.0
  #include "go-cef.h"
  #include <stdlib.h>
*/
import "C"

import (
  "fmt"
  "os"
  "unsafe"
)

func main() {
   fmt.Println("cef poziv 1")
   go_main_cef()

   fmt.Println("cef poziv 2")
   go_main_cef()


}
func go_main_cef() {

  argc := C.int(len(os.Args)) 
  argv := make([]*C.char, argc) 
  // napuni matricu argv
  for i, arg := range os.Args { 
       argv[i] = C.CString(arg) 
  } 

  fmt.Println("go-chromium embedded framework")

  C.cef_main(argc, &argv[0])

  // oslobodi argv
  for _, arg := range argv { 
        C.free(unsafe.Pointer(arg)) 
  }

  fmt.Println("---- kraj ----") 
}


