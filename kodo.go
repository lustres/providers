package providers

import (
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

func BucketManagerDefault(ak string, sk []byte) *storage.BucketManager {
	return BucketManager(ak, sk, nil)
}

func BucketManager(ak string, sk []byte, cfg *storage.Config) *storage.BucketManager {
	return BucketManagerWithMac(qbox.NewMac(ak, string(sk)), cfg)
}

var BucketManagerWithMac = storage.NewBucketManager
