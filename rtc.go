package providers

import (
	"github.com/google/wire"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/rtc"
)

var RTCManager = wire.NewSet(RTCManagerWithMac, QBoxMac)

func RTCManagerWithMac(mac *qbox.Mac) *rtc.Manager {
	return rtc.NewManager(mac)
}
