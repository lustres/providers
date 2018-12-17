package providers

import (
	"net/http"

	"github.com/pili-engineering/pili-sdk-go.v2/pili"
)

func PiliClientDefault(ak string, sk []byte) *pili.Client {
	return PiliClient(ak, sk, http.DefaultTransport)
}

func PiliClient(ak string, sk []byte, tran http.RoundTripper) *pili.Client {
	return PiliClientWithMac(&pili.MAC{ak, sk}, tran)
}

var PiliClientWithMac = pili.New
