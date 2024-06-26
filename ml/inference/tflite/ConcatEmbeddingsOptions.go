// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ConcatEmbeddingsOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsConcatEmbeddingsOptions(buf []byte, offset flatbuffers.UOffsetT) *ConcatEmbeddingsOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ConcatEmbeddingsOptions{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsConcatEmbeddingsOptions(buf []byte, offset flatbuffers.UOffsetT) *ConcatEmbeddingsOptions {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ConcatEmbeddingsOptions{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ConcatEmbeddingsOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ConcatEmbeddingsOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ConcatEmbeddingsOptions) NumChannels() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ConcatEmbeddingsOptions) MutateNumChannels(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *ConcatEmbeddingsOptions) NumColumnsPerChannel(j int) int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *ConcatEmbeddingsOptions) NumColumnsPerChannelLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ConcatEmbeddingsOptions) MutateNumColumnsPerChannel(j int, n int32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *ConcatEmbeddingsOptions) EmbeddingDimPerChannel(j int) int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *ConcatEmbeddingsOptions) EmbeddingDimPerChannelLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ConcatEmbeddingsOptions) MutateEmbeddingDimPerChannel(j int, n int32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func ConcatEmbeddingsOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func ConcatEmbeddingsOptionsAddNumChannels(builder *flatbuffers.Builder, numChannels int32) {
	builder.PrependInt32Slot(0, numChannels, 0)
}
func ConcatEmbeddingsOptionsAddNumColumnsPerChannel(builder *flatbuffers.Builder, numColumnsPerChannel flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(numColumnsPerChannel), 0)
}
func ConcatEmbeddingsOptionsStartNumColumnsPerChannelVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ConcatEmbeddingsOptionsAddEmbeddingDimPerChannel(builder *flatbuffers.Builder, embeddingDimPerChannel flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(embeddingDimPerChannel), 0)
}
func ConcatEmbeddingsOptionsStartEmbeddingDimPerChannelVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ConcatEmbeddingsOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
