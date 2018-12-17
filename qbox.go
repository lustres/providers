package providers

import (
	"github.com/qiniu/api.v7/auth/qbox"
)

func QBoxMac(ak string, sk []byte) *qbox.Mac {
	return qbox.NewMac(ak, string(sk))
}
