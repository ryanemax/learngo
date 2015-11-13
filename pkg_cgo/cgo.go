package pkgcgo

/*
#include <stdio.h>
#include <stdlib.h>

void myprint(char* s){
	printf("%s",s);
}
*/

import (
	"C"
	"unsafe"
)

func Example() {
	cs := C.CString("Hello from stdio\n")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}

//func main() {
//	Example()
//}
