//go:build windows && cgo

// Package fitzstub provides a MinGW-compatible symbol for MuPDF static libs
// that were compiled against MSVC setjmp intrinsics.
package fitzstub

/*
#include <setjmp.h>

// Some prebuilt MuPDF objects reference MSVC's __intrinsic_setjmpex.
// When linking with MinGW/GCC, define a thin wrapper around setjmp.
int __intrinsic_setjmpex(jmp_buf env) {
	return setjmp(env);
}
*/
import "C"
