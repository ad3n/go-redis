package proto

import (
	"bufio"
	"sync"
)

var (
	readerPool = sync.Pool{
		New: func() any {
			return bufio.NewReaderSize(nil, 4096)
		},
	}

	bytePool = sync.Pool{
		New: func() any {
			b := make([]byte, 64)
			return &b
		},
	}
)
