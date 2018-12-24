package providers

import (
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/rtc"
)

func RTCManager(ak string, sk []byte) *rtc.Manager {
	return rtc.NewManager(qbox.NewMac(ak, string(sk)))
}

func RTCManagerWithMac(mac *qbox.Mac) *rtc.Manager {
	return rtc.NewManager(mac)
}
