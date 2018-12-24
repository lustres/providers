package providers

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/google/wire"
)

var ReadFile = wire.NewSet(ioutil.ReadFile, bytes.NewReader, wire.Bind(new(io.Reader), new(bytes.Reader)))
