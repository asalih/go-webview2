//go:build windows

package webview2

type _ICoreWebView2FrameDOMContentLoadedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2FrameDOMContentLoadedEventHandler struct {
	vtbl *_ICoreWebView2FrameDOMContentLoadedEventHandlerVtbl
	impl _ICoreWebView2FrameDOMContentLoadedEventHandlerImpl
}

func (i *ICoreWebView2FrameDOMContentLoadedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2FrameDOMContentLoadedEventHandlerIUnknownQueryInterface(this *ICoreWebView2FrameDOMContentLoadedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2FrameDOMContentLoadedEventHandlerIUnknownAddRef(this *ICoreWebView2FrameDOMContentLoadedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2FrameDOMContentLoadedEventHandlerIUnknownRelease(this *ICoreWebView2FrameDOMContentLoadedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2FrameDOMContentLoadedEventHandlerInvoke(this *ICoreWebView2FrameDOMContentLoadedEventHandler, sender *ICoreWebView2Frame, args *ICoreWebView2DOMContentLoadedEventArgs) uintptr {
	return this.impl.FrameDOMContentLoaded(sender, args)
}

type _ICoreWebView2FrameDOMContentLoadedEventHandlerImpl interface {
	_IUnknownImpl
	FrameDOMContentLoaded(sender *ICoreWebView2Frame, args *ICoreWebView2DOMContentLoadedEventArgs) uintptr
}

var _ICoreWebView2FrameDOMContentLoadedEventHandlerFn = _ICoreWebView2FrameDOMContentLoadedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2FrameDOMContentLoadedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2FrameDOMContentLoadedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2FrameDOMContentLoadedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2FrameDOMContentLoadedEventHandlerInvoke),
}

func NewICoreWebView2FrameDOMContentLoadedEventHandler(impl _ICoreWebView2FrameDOMContentLoadedEventHandlerImpl) *ICoreWebView2FrameDOMContentLoadedEventHandler {
	return &ICoreWebView2FrameDOMContentLoadedEventHandler{
		vtbl: &_ICoreWebView2FrameDOMContentLoadedEventHandlerFn,
		impl: impl,
	}
}
