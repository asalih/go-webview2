//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2AcceleratorKeyPressedEventArgsVtbl struct {
	_IUnknownVtbl
	GetKeyEventKind      ComProc
	GetVirtualKey        ComProc
	GetKeyEventLParam    ComProc
	GetPhysicalKeyStatus ComProc
	GetHandled           ComProc
	PutHandled           ComProc
}

type ICoreWebView2AcceleratorKeyPressedEventArgs struct {
	vtbl *_ICoreWebView2AcceleratorKeyPressedEventArgsVtbl
}

func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) GetKeyEventKind() (*COREWEBVIEW2_KEY_EVENT_KIND, error) {
	var err error

	var keyEventKind *COREWEBVIEW2_KEY_EVENT_KIND

	_, _, err = i.vtbl.GetKeyEventKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&keyEventKind)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return keyEventKind, nil
}

func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) GetVirtualKey() (uint, error) {
	var err error

	var virtualKey uint

	_, _, err = i.vtbl.GetVirtualKey.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&virtualKey)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return virtualKey, nil
}

func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) GetKeyEventLParam() (int, error) {
	var err error

	var lParam int

	_, _, err = i.vtbl.GetKeyEventLParam.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&lParam)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}
	return lParam, nil
}

func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) GetPhysicalKeyStatus() (*COREWEBVIEW2_PHYSICAL_KEY_STATUS, error) {
	var err error

	var physicalKeyStatus *COREWEBVIEW2_PHYSICAL_KEY_STATUS

	_, _, err = i.vtbl.GetPhysicalKeyStatus.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&physicalKeyStatus)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return physicalKeyStatus, nil
}

func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) GetHandled() (bool, error) {
	var err error

	var handled bool

	_, _, err = i.vtbl.GetHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&handled)),
	)
	if err != windows.ERROR_SUCCESS {
		return false, err
	}
	return handled, nil
}

func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) PutHandled(handled bool) error {
	var err error

	_, _, err = i.vtbl.PutHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&handled)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}
