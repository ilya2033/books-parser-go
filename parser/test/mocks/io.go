package mocks

import (
	"bytes"
	"io"
)

type ReadCloserMock struct {
	Data   *bytes.Reader
	Error  error
	Closed bool
}

func (this *ReadCloserMock) Read(p []byte) (int, error) {
	if this.Error != nil {
		return 0, this.Error
	}

	if this.Data == nil {
		return 0, io.EOF
	}

	return this.Data.Read(p)
}

func (this *ReadCloserMock) Close() error {
	this.Closed = true
	return nil
}
