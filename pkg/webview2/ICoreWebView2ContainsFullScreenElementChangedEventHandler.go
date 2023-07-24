//go:build windows

package webview2

type _ICoreWebView2ContainsFullScreenElementChangedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type ICoreWebView2ContainsFullScreenElementChangedEventHandler struct {
	vtbl *_ICoreWebView2ContainsFullScreenElementChangedEventHandlerVtbl
	impl _ICoreWebView2ContainsFullScreenElementChangedEventHandlerImpl
}

func (i *ICoreWebView2ContainsFullScreenElementChangedEventHandler) AddRef() uintptr {
	return i.AddRef()
}
func _ICoreWebView2ContainsFullScreenElementChangedEventHandlerIUnknownQueryInterface(this *ICoreWebView2ContainsFullScreenElementChangedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ContainsFullScreenElementChangedEventHandlerIUnknownAddRef(this *ICoreWebView2ContainsFullScreenElementChangedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ContainsFullScreenElementChangedEventHandlerIUnknownRelease(this *ICoreWebView2ContainsFullScreenElementChangedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ContainsFullScreenElementChangedEventHandlerInvoke(this *ICoreWebView2ContainsFullScreenElementChangedEventHandler, sender *ICoreWebView2, args *_IUnknown) uintptr {
	return this.impl.ContainsFullScreenElementChanged(sender, args)
}

type _ICoreWebView2ContainsFullScreenElementChangedEventHandlerImpl interface {
	_IUnknownImpl
	ContainsFullScreenElementChanged(sender *ICoreWebView2, args *_IUnknown) uintptr
}

var _ICoreWebView2ContainsFullScreenElementChangedEventHandlerFn = _ICoreWebView2ContainsFullScreenElementChangedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2ContainsFullScreenElementChangedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2ContainsFullScreenElementChangedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2ContainsFullScreenElementChangedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2ContainsFullScreenElementChangedEventHandlerInvoke),
}

func NewICoreWebView2ContainsFullScreenElementChangedEventHandler(impl _ICoreWebView2ContainsFullScreenElementChangedEventHandlerImpl) *ICoreWebView2ContainsFullScreenElementChangedEventHandler {
	return &ICoreWebView2ContainsFullScreenElementChangedEventHandler{
		vtbl: &_ICoreWebView2ContainsFullScreenElementChangedEventHandlerFn,
		impl: impl,
	}
}
