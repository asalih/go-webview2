//go:build windows

package webview2

type COREWEBVIEW2_NON_CLIENT_REGION_KIND uint32

const (
	COREWEBVIEW2_NON_CLIENT_REGION_KIND_NOWHERE  = 0
	COREWEBVIEW2_NON_CLIENT_REGION_KIND_CLIENT   = 1
	COREWEBVIEW2_NON_CLIENT_REGION_KIND_CAPTION  = 2
	COREWEBVIEW2_NON_CLIENT_REGION_KIND_MINIMIZE = 8
	COREWEBVIEW2_NON_CLIENT_REGION_KIND_MAXIMIZE = 9
	COREWEBVIEW2_NON_CLIENT_REGION_KIND_CLOSE    = 20
)
