// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package streams represents the imported interface "wasi:io/streams@0.2.0".
package streams

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
	ioerror "github.com/wasmcloud/wasmcloud/examples/golang/components/http-client-tinygo/gen/wasi/io/error"
	"github.com/wasmcloud/wasmcloud/examples/golang/components/http-client-tinygo/gen/wasi/io/poll"
)

// StreamError represents the imported variant "wasi:io/streams@0.2.0#stream-error".
//
//	variant stream-error {
//		last-operation-failed(error),
//		closed,
//	}
type StreamError cm.Variant[uint8, ioerror.Error, ioerror.Error]

// StreamErrorLastOperationFailed returns a [StreamError] of case "last-operation-failed".
func StreamErrorLastOperationFailed(data ioerror.Error) StreamError {
	return cm.New[StreamError](0, data)
}

// LastOperationFailed returns a non-nil *[ioerror.Error] if [StreamError] represents the variant case "last-operation-failed".
func (self *StreamError) LastOperationFailed() *ioerror.Error {
	return cm.Case[ioerror.Error](self, 0)
}

// StreamErrorClosed returns a [StreamError] of case "closed".
func StreamErrorClosed() StreamError {
	var data struct{}
	return cm.New[StreamError](1, data)
}

// Closed returns true if [StreamError] represents the variant case "closed".
func (self *StreamError) Closed() bool {
	return self.Tag() == 1
}

// InputStream represents the imported resource "wasi:io/streams@0.2.0#input-stream".
//
//	resource input-stream
type InputStream cm.Resource

// ResourceDrop represents the imported resource-drop for resource "input-stream".
//
// Drops a resource handle.
//
//go:nosplit
func (self InputStream) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_InputStreamResourceDrop((uint32)(self0))
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [resource-drop]input-stream
//go:noescape
func wasmimport_InputStreamResourceDrop(self0 uint32)

// BlockingRead represents the imported method "blocking-read".
//
//	blocking-read: func(len: u64) -> result<list<u8>, stream-error>
//
//go:nosplit
func (self InputStream) BlockingRead(len_ uint64) (result cm.Result[cm.List[uint8], cm.List[uint8], StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	len0 := (uint64)(len_)
	wasmimport_InputStreamBlockingRead((uint32)(self0), (uint64)(len0), &result)
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]input-stream.blocking-read
//go:noescape
func wasmimport_InputStreamBlockingRead(self0 uint32, len0 uint64, result *cm.Result[cm.List[uint8], cm.List[uint8], StreamError])

// BlockingSkip represents the imported method "blocking-skip".
//
//	blocking-skip: func(len: u64) -> result<u64, stream-error>
//
//go:nosplit
func (self InputStream) BlockingSkip(len_ uint64) (result cm.Result[uint64, uint64, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	len0 := (uint64)(len_)
	wasmimport_InputStreamBlockingSkip((uint32)(self0), (uint64)(len0), &result)
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]input-stream.blocking-skip
//go:noescape
func wasmimport_InputStreamBlockingSkip(self0 uint32, len0 uint64, result *cm.Result[uint64, uint64, StreamError])

// Read represents the imported method "read".
//
//	read: func(len: u64) -> result<list<u8>, stream-error>
//
//go:nosplit
func (self InputStream) Read(len_ uint64) (result cm.Result[cm.List[uint8], cm.List[uint8], StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	len0 := (uint64)(len_)
	wasmimport_InputStreamRead((uint32)(self0), (uint64)(len0), &result)
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]input-stream.read
//go:noescape
func wasmimport_InputStreamRead(self0 uint32, len0 uint64, result *cm.Result[cm.List[uint8], cm.List[uint8], StreamError])

// Skip represents the imported method "skip".
//
//	skip: func(len: u64) -> result<u64, stream-error>
//
//go:nosplit
func (self InputStream) Skip(len_ uint64) (result cm.Result[uint64, uint64, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	len0 := (uint64)(len_)
	wasmimport_InputStreamSkip((uint32)(self0), (uint64)(len0), &result)
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]input-stream.skip
//go:noescape
func wasmimport_InputStreamSkip(self0 uint32, len0 uint64, result *cm.Result[uint64, uint64, StreamError])

// Subscribe represents the imported method "subscribe".
//
//	subscribe: func() -> pollable
//
//go:nosplit
func (self InputStream) Subscribe() (result poll.Pollable) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_InputStreamSubscribe((uint32)(self0))
	result = cm.Reinterpret[poll.Pollable]((uint32)(result0))
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]input-stream.subscribe
//go:noescape
func wasmimport_InputStreamSubscribe(self0 uint32) (result0 uint32)

// OutputStream represents the imported resource "wasi:io/streams@0.2.0#output-stream".
//
//	resource output-stream
type OutputStream cm.Resource

// ResourceDrop represents the imported resource-drop for resource "output-stream".
//
// Drops a resource handle.
//
//go:nosplit
func (self OutputStream) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_OutputStreamResourceDrop((uint32)(self0))
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [resource-drop]output-stream
//go:noescape
func wasmimport_OutputStreamResourceDrop(self0 uint32)

// BlockingFlush represents the imported method "blocking-flush".
//
//	blocking-flush: func() -> result<_, stream-error>
//
//go:nosplit
func (self OutputStream) BlockingFlush() (result cm.Result[StreamError, struct{}, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_OutputStreamBlockingFlush((uint32)(self0), &result)
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.blocking-flush
//go:noescape
func wasmimport_OutputStreamBlockingFlush(self0 uint32, result *cm.Result[StreamError, struct{}, StreamError])

// BlockingSplice represents the imported method "blocking-splice".
//
//	blocking-splice: func(src: borrow<input-stream>, len: u64) -> result<u64, stream-error>
//
//go:nosplit
func (self OutputStream) BlockingSplice(src InputStream, len_ uint64) (result cm.Result[uint64, uint64, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	src0 := cm.Reinterpret[uint32](src)
	len0 := (uint64)(len_)
	wasmimport_OutputStreamBlockingSplice((uint32)(self0), (uint32)(src0), (uint64)(len0), &result)
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.blocking-splice
//go:noescape
func wasmimport_OutputStreamBlockingSplice(self0 uint32, src0 uint32, len0 uint64, result *cm.Result[uint64, uint64, StreamError])

// BlockingWriteAndFlush represents the imported method "blocking-write-and-flush".
//
//	blocking-write-and-flush: func(contents: list<u8>) -> result<_, stream-error>
//
//go:nosplit
func (self OutputStream) BlockingWriteAndFlush(contents cm.List[uint8]) (result cm.Result[StreamError, struct{}, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	contents0, contents1 := cm.LowerList(contents)
	wasmimport_OutputStreamBlockingWriteAndFlush((uint32)(self0), (*uint8)(contents0), (uint32)(contents1), &result)
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.blocking-write-and-flush
//go:noescape
func wasmimport_OutputStreamBlockingWriteAndFlush(self0 uint32, contents0 *uint8, contents1 uint32, result *cm.Result[StreamError, struct{}, StreamError])

// BlockingWriteZeroesAndFlush represents the imported method "blocking-write-zeroes-and-flush".
//
//	blocking-write-zeroes-and-flush: func(len: u64) -> result<_, stream-error>
//
//go:nosplit
func (self OutputStream) BlockingWriteZeroesAndFlush(len_ uint64) (result cm.Result[StreamError, struct{}, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	len0 := (uint64)(len_)
	wasmimport_OutputStreamBlockingWriteZeroesAndFlush((uint32)(self0), (uint64)(len0), &result)
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.blocking-write-zeroes-and-flush
//go:noescape
func wasmimport_OutputStreamBlockingWriteZeroesAndFlush(self0 uint32, len0 uint64, result *cm.Result[StreamError, struct{}, StreamError])

// CheckWrite represents the imported method "check-write".
//
//	check-write: func() -> result<u64, stream-error>
//
//go:nosplit
func (self OutputStream) CheckWrite() (result cm.Result[uint64, uint64, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_OutputStreamCheckWrite((uint32)(self0), &result)
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.check-write
//go:noescape
func wasmimport_OutputStreamCheckWrite(self0 uint32, result *cm.Result[uint64, uint64, StreamError])

// Flush represents the imported method "flush".
//
//	flush: func() -> result<_, stream-error>
//
//go:nosplit
func (self OutputStream) Flush() (result cm.Result[StreamError, struct{}, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_OutputStreamFlush((uint32)(self0), &result)
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.flush
//go:noescape
func wasmimport_OutputStreamFlush(self0 uint32, result *cm.Result[StreamError, struct{}, StreamError])

// Splice represents the imported method "splice".
//
//	splice: func(src: borrow<input-stream>, len: u64) -> result<u64, stream-error>
//
//go:nosplit
func (self OutputStream) Splice(src InputStream, len_ uint64) (result cm.Result[uint64, uint64, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	src0 := cm.Reinterpret[uint32](src)
	len0 := (uint64)(len_)
	wasmimport_OutputStreamSplice((uint32)(self0), (uint32)(src0), (uint64)(len0), &result)
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.splice
//go:noescape
func wasmimport_OutputStreamSplice(self0 uint32, src0 uint32, len0 uint64, result *cm.Result[uint64, uint64, StreamError])

// Subscribe represents the imported method "subscribe".
//
//	subscribe: func() -> pollable
//
//go:nosplit
func (self OutputStream) Subscribe() (result poll.Pollable) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_OutputStreamSubscribe((uint32)(self0))
	result = cm.Reinterpret[poll.Pollable]((uint32)(result0))
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.subscribe
//go:noescape
func wasmimport_OutputStreamSubscribe(self0 uint32) (result0 uint32)

// Write represents the imported method "write".
//
//	write: func(contents: list<u8>) -> result<_, stream-error>
//
//go:nosplit
func (self OutputStream) Write(contents cm.List[uint8]) (result cm.Result[StreamError, struct{}, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	contents0, contents1 := cm.LowerList(contents)
	wasmimport_OutputStreamWrite((uint32)(self0), (*uint8)(contents0), (uint32)(contents1), &result)
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.write
//go:noescape
func wasmimport_OutputStreamWrite(self0 uint32, contents0 *uint8, contents1 uint32, result *cm.Result[StreamError, struct{}, StreamError])

// WriteZeroes represents the imported method "write-zeroes".
//
//	write-zeroes: func(len: u64) -> result<_, stream-error>
//
//go:nosplit
func (self OutputStream) WriteZeroes(len_ uint64) (result cm.Result[StreamError, struct{}, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	len0 := (uint64)(len_)
	wasmimport_OutputStreamWriteZeroes((uint32)(self0), (uint64)(len0), &result)
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.write-zeroes
//go:noescape
func wasmimport_OutputStreamWriteZeroes(self0 uint32, len0 uint64, result *cm.Result[StreamError, struct{}, StreamError])
