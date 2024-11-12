// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024 Hajime Hoshi

//go:generate go run gen.go

package libvorbis

// #include <stdlib.h>
// #include "vorbis_codec.h"
import "C"

import (
	"fmt"
	"runtime"
	"unsafe"
)

// https://xiph.org/vorbis/doc/libvorbis/reference.html

type Error C.int

const (
	ErrFalse     Error = C.OV_FALSE
	ErrHole      Error = C.OV_HOLE
	ErrRead      Error = C.OV_EREAD
	ErrFault     Error = C.OV_EFAULT
	ErrImpl      Error = C.OV_EIMPL
	ErrInval     Error = C.OV_EINVAL
	ErrNotVorbis Error = C.OV_ENOTVORBIS
	ErrBadHeader Error = C.OV_EBADHEADER
	ErrVersion   Error = C.OV_EVERSION
	ErrNotAudio  Error = C.OV_ENOTAUDIO
	ErrBadPacket Error = C.OV_EBADPACKET
	ErrBadLink   Error = C.OV_EBADLINK
	ErrNoSeek    Error = C.OV_ENOSEEK
)

func (e Error) Error() string {
	switch e {
	case ErrFalse:
		return "OV_FALSE"
	case ErrHole:
		return "OV_HOLE"
	case ErrRead:
		return "OV_EREAD"
	case ErrFault:
		return "OV_EFAULT"
	case ErrImpl:
		return "OV_EIMPL"
	case ErrInval:
		return "OV_EINVAL"
	case ErrNotVorbis:
		return "OV_ENOTVORBIS"
	case ErrBadHeader:
		return "OV_EBADHEADER"
	case ErrVersion:
		return "OV_EVERSION"
	case ErrNotAudio:
		return "OV_ENOTAUDIO"
	case ErrBadPacket:
		return "OV_EBADPACKET"
	case ErrBadLink:
		return "OV_EBADLINK"
	case ErrNoSeek:
		return "OV_ENOSEEK"
	default:
		return fmt.Sprintf("Error(%d)", e)
	}
}

type OggPacket struct {
	Packet     []byte
	BOS        bool
	EOS        bool
	GranulePos int64
	PacketNo   int64
}

func (o *OggPacket) c() *C.ogg_packet {
	cPacket := C.CBytes(o.Packet)
	c := &C.ogg_packet{
		packet:     (*C.uchar)(cPacket),
		bytes:      C.long(len(o.Packet)),
		b_o_s:      C.long(btoi(o.BOS)),
		e_o_s:      C.long(btoi(o.EOS)),
		granulepos: C.ogg_int64_t(o.GranulePos),
		packetno:   C.ogg_int64_t(o.PacketNo),
	}
	runtime.SetFinalizer(o, func(o *OggPacket) {
		C.free(unsafe.Pointer(cPacket))
	})
	return c
}

type Block struct {
	c *C.vorbis_block
}

type Comment struct {
	c C.vorbis_comment
}

func (c *Comment) UserComments() []string {
	cUserComments := unsafe.Slice((**C.char)(unsafe.Pointer(c.c.user_comments)), c.c.comments)
	commentLengths := unsafe.Slice((*C.int)(unsafe.Pointer(c.c.comment_lengths)), c.c.comments)

	var userComments []string
	for i := range c.c.comments {
		userComments = append(userComments, C.GoStringN(cUserComments[i], commentLengths[i]))
	}

	return userComments
}

func (c *Comment) Vendor() string {
	return C.GoString(c.c.vendor)
}

type DspState struct {
	c *C.vorbis_dsp_state
}

type Info struct {
	c C.vorbis_info
}

func (i *Info) Channels() int {
	return int(i.c.channels)
}

func (i *Info) Rate() int {
	return int(i.c.rate)
}

func InfoInit() *Info {
	var cInfo C.vorbis_info
	C.vorbis_info_init(&cInfo)
	return &Info{c: cInfo}
}

func Synthesis(vb *Block, op *OggPacket) error {
	cOp := op.c()
	defer runtime.KeepAlive(vb)
	defer runtime.KeepAlive(op)
	if ret := C.vorbis_synthesis(vb.c, cOp); ret != 0 {
		return Error(ret)
	}
	return nil
}

func SynthesisBlockin(vd *DspState, vb *Block) error {
	defer runtime.KeepAlive(vd)
	defer runtime.KeepAlive(vb)
	if ret := C.vorbis_synthesis_blockin(vd.c, vb.c); ret != 0 {
		return Error(ret)
	}
	return nil
}

func SynthesisHeaderin(vi *Info, vc *Comment, op *OggPacket) error {
	cOp := op.c()
	defer runtime.KeepAlive(vi)
	defer runtime.KeepAlive(vc)
	defer runtime.KeepAlive(op)
	if ret := C.vorbis_synthesis_headerin(&vi.c, &vc.c, cOp); ret != 0 {
		return Error(ret)
	}
	return nil
}

func SynthesisInit(vi *Info) (*DspState, error) {
	cDspState := (*C.vorbis_dsp_state)(C.calloc(1, C.size_t(unsafe.Sizeof(C.vorbis_dsp_state{}))))
	d := &DspState{c: cDspState}
	runtime.SetFinalizer(d, func(d *DspState) {
		// TODO: Call C.vorbis_dsp_clear(d.c)?
		C.free(unsafe.Pointer(d.c))
	})

	defer runtime.KeepAlive(vi)
	if ret := C.vorbis_synthesis_init(cDspState, &vi.c); ret != 0 {
		return nil, Error(ret)
	}
	return d, nil
}

func SynthesisPcmout(vd *DspState) [][]float32 {
	var cPCM **C.float
	defer runtime.KeepAlive(vd)
	n := C.vorbis_synthesis_pcmout(vd.c, &cPCM)
	if n == 0 {
		return nil
	}

	cPCMPtrs := unsafe.Slice(cPCM, int(vd.c.vi.channels))
	pcms := make([][]float32, len(cPCMPtrs))
	for i, cPCMPtr := range cPCMPtrs {
		pcms[i] = make([]float32, n)
		copy(pcms[i], unsafe.Slice((*float32)(unsafe.Pointer(cPCMPtr)), int(n)))
	}
	return pcms
}

func SynthesisRead(vd *DspState, samples int) error {
	defer runtime.KeepAlive(vd)
	if ret := C.vorbis_synthesis_read(vd.c, C.int(samples)); ret != 0 {
		return Error(ret)
	}
	return nil
}

func BlockInit(vd *DspState) (*Block, error) {
	cBlock := (*C.vorbis_block)(C.calloc(1, C.size_t(unsafe.Sizeof(C.vorbis_block{}))))
	b := &Block{c: cBlock}
	runtime.SetFinalizer(b, func(b *Block) {
		// TODO: Call C.vorbis_block_clear(b.c)?
		C.free(unsafe.Pointer(b.c))
	})

	if ret := C.vorbis_block_init(vd.c, cBlock); ret != 0 {
		return nil, Error(ret)
	}
	return b, nil
}

func CommentInit() *Comment {
	var cComment C.vorbis_comment
	C.vorbis_comment_init(&cComment)
	return &Comment{c: cComment}
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
