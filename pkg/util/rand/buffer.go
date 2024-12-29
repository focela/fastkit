// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package random provides high-performance random bytes/number/string generation functionality.
package random

import (
	"crypto/rand"

	"github.com/focela/loom/pkg/errors"
	"github.com/focela/loom/pkg/errors/code"
)

const (
	// bufferChanSize defines the size of the random byte buffer channel.
	bufferChanSize = 10000
	// bufferChunkSize defines the size of each random byte chunk.
	bufferChunkSize = 1024
	// stepSize defines the step increment for slicing the random byte buffer.
	stepSize = 4
)

// bufferChan serves as a channel to store random byte chunks for high-performance access.
var bufferChan = make(chan []byte, bufferChanSize)

// init starts a goroutine to buffer random bytes for high-performance random generation.
func init() {
	// Start asynchronous production of random bytes.
	go produceRandomBufferBytesAsync()
}

// produceRandomBufferBytesAsync continuously generates random bytes and stores them in bufferChan.
// This approach avoids repeated expensive system calls to fetch random data.
func produceRandomBufferBytesAsync() {
	for {
		buffer := make([]byte, bufferChunkSize)
		n, err := rand.Read(buffer)
		if err != nil {
			panic(errors.WrapCode(code.CodeInternalError, err, "error reading random buffer from system"))
		}

		// Efficiently distribute random bytes into bufferChan using defined step size.
		for i := 0; i <= n-stepSize; i += stepSize {
			bufferChan <- buffer[i : i+stepSize]
		}
	}
}
