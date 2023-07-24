//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2Controller2Vtbl struct {
	_IUnknownVtbl
	GetDefaultBackgroundColor ComProc
	PutDefaultBackgroundColor ComProc
}

type ICoreWebView2Controller2 struct {
	vtbl *_ICoreWebView2Controller2Vtbl
}

func (i *ICoreWebView2Controller2) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Controller2) GetDefaultBackgroundColor() (*COREWEBVIEW2_COLOR, error) {
	var err error

	var backgroundColor *COREWEBVIEW2_COLOR

	_, _, err = i.vtbl.GetDefaultBackgroundColor.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&backgroundColor)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return backgroundColor, nil
}

func (i *ICoreWebView2Controller2) PutDefaultBackgroundColor(backgroundColor COREWEBVIEW2_COLOR) error {
	var err error

	_, _, err = i.vtbl.PutDefaultBackgroundColor.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&backgroundColor)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
