package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include "window.h"
#include "application.h"
*/
import "C"
import (
    "unsafe"
)

type Window struct {
    ptr unsafe.Pointer
}

func CreateWindow() *Window {
    w := new(Window)

    title := C.CString("School 21")
    defer C.free(unsafe.Pointer(title))

    w.ptr = C.Window_Create(500, 500, 200, 300, title)
    return w
}

func (w *Window) MakeKeyAndOrderFront() {
    C.Window_MakeKeyAndOrderFront(w.ptr)
}

func main() {
    w := CreateWindow()
    w.MakeKeyAndOrderFront()
    C.InitApplication()
    C.RunApplication()
}
