//go:build windows

package webview2

type COREWEBVIEW2_BROWSING_DATA_KINDS uint32

const (
	COREWEBVIEW2_BROWSING_DATA_KINDS_FILE_SYSTEMS      = 1 << 0
	COREWEBVIEW2_BROWSING_DATA_KINDS_INDEXED_DB        = 1 << 1
	COREWEBVIEW2_BROWSING_DATA_KINDS_LOCAL_STORAGE     = 1 << 2
	COREWEBVIEW2_BROWSING_DATA_KINDS_WEB_SQL           = 1 << 3
	COREWEBVIEW2_BROWSING_DATA_KINDS_CACHE_STORAGE     = 1 << 4
	COREWEBVIEW2_BROWSING_DATA_KINDS_ALL_DOM_STORAGE   = 1 << 5
	COREWEBVIEW2_BROWSING_DATA_KINDS_COOKIES           = 1 << 6
	COREWEBVIEW2_BROWSING_DATA_KINDS_ALL_SITE          = 1 << 7
	COREWEBVIEW2_BROWSING_DATA_KINDS_DISK_CACHE        = 1 << 8
	COREWEBVIEW2_BROWSING_DATA_KINDS_DOWNLOAD_HISTORY  = 1 << 9
	COREWEBVIEW2_BROWSING_DATA_KINDS_GENERAL_AUTOFILL  = 1 << 10
	COREWEBVIEW2_BROWSING_DATA_KINDS_PASSWORD_AUTOSAVE = 1 << 11
	COREWEBVIEW2_BROWSING_DATA_KINDS_BROWSING_HISTORY  = 1 << 12
	COREWEBVIEW2_BROWSING_DATA_KINDS_SETTINGS          = 1 << 13
	COREWEBVIEW2_BROWSING_DATA_KINDS_ALL_PROFILE       = 1 << 14
)
