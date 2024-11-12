// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024 Hajime Hoshi

//go:generate go run gen.go

package libopus

// #cgo CFLAGS: -DOPUS_BUILD -DUSE_ALLOCA -DHAVE_LRINT -DHAVE_LRINTF
//
// #include "opus.h"
import "C"

import (
	"fmt"
	"unsafe"
)

type Error C.int

const (
	ErrBadArg         Error = C.OPUS_BAD_ARG
	ErrBufferTooSmall Error = C.OPUS_BUFFER_TOO_SMALL
	ErrInternalError  Error = C.OPUS_INTERNAL_ERROR
	ErrInvalidPacket  Error = C.OPUS_INVALID_PACKET
	ErrUnimplemented  Error = C.OPUS_UNIMPLEMENTED
	ErrInvalidState   Error = C.OPUS_INVALID_STATE
	ErrAllocFail      Error = C.OPUS_ALLOC_FAIL
)

func (e Error) Error() string {
	switch e {
	case ErrBadArg:
		return "OPUS_BAD_ARG"
	case ErrBufferTooSmall:
		return "OPUS_BUFFER_TOO_SMALL"
	case ErrInternalError:
		return "OPUS_INTERNAL_ERROR"
	case ErrInvalidPacket:
		return "OPUS_INVALID_PACKET"
	case ErrUnimplemented:
		return "OPUS_UNIMPLEMENTED"
	case ErrInvalidState:
		return "OPUS_INVALID_STATE"
	case ErrAllocFail:
		return "OPUS_ALLOC_FAIL"
	default:
		return fmt.Sprintf("Error(%d)", e)
	}
}

type Decoder struct {
	decoder *C.OpusDecoder
}

func DecoderCreate(Fs int, channels int) (*Decoder, error) {
	var err C.int
	d := C.opus_decoder_create(C.opus_int32(Fs), C.int(channels), &err)
	if err != C.OPUS_OK {
		return nil, Error(err)
	}
	return &Decoder{
		decoder: d,
	}, nil
}

func (d *Decoder) DecodeFloat(data []byte, pcm []float32, decodeFec int) int {
	n := C.opus_decode_float(
		d.decoder,
		(*C.uchar)(unsafe.Pointer(unsafe.SliceData(data))),
		C.opus_int32(len(data)),
		(*C.float)(unsafe.Pointer(unsafe.SliceData(pcm))),
		C.int(len(pcm)),
		C.int(decodeFec))
	return int(n)
}
