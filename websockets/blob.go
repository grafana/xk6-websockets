package websockets

import (
	"bytes"
	"fmt"

	"github.com/grafana/sobek"

	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules/k6/experimental/streams"
)

type blob struct {
	typ  string
	data bytes.Buffer
}

func (b *blob) text() string {
	return b.data.String()
}

func (r *WebSocketsAPI) blob(call sobek.ConstructorCall) *sobek.Object {
	rt := r.vu.Runtime()

	var blobParts []interface{}
	if len(call.Arguments) > 0 {
		if err := rt.ExportTo(call.Arguments[0], &blobParts); err != nil {
			common.Throw(rt, fmt.Errorf("failed to process [blobParts]: %w", err))
		}
	}

	b := &blob{}
	if len(blobParts) > 0 {
		for _, part := range blobParts {
			var err error
			switch v := part.(type) {
			case []uint8:
				_, err = b.data.Write(v)
			case string:
				_, err = b.data.WriteString(v)
			}
			if err != nil {
				common.Throw(rt, fmt.Errorf("failed to process [blobParts]: %w", err))
			}
		}
	}

	obj := rt.NewObject()
	must(rt, obj.DefineAccessorProperty("type", rt.ToValue(func() sobek.Value {
		return rt.ToValue(b.typ)
	}), nil, sobek.FLAG_FALSE, sobek.FLAG_TRUE))
	must(rt, obj.DefineAccessorProperty("size", rt.ToValue(func() sobek.Value {
		return rt.ToValue(b.data.Len())
	}), nil, sobek.FLAG_FALSE, sobek.FLAG_TRUE))
	must(rt, obj.Set("text", func(_ sobek.FunctionCall) sobek.Value {
		return rt.ToValue(b.text())
	}))
	must(rt, obj.Set("stream", func(_ sobek.FunctionCall) sobek.Value {
		return rt.ToValue(streams.NewReadableStreamFromReader(r.vu, &b.data))
	}))
	must(rt, obj.Set("arrayBuffer", func(_ sobek.FunctionCall) sobek.Value {
		return rt.ToValue(rt.NewArrayBuffer(b.data.Bytes()))
	}))

	proto := call.This.Prototype()
	must(rt, proto.Set("toString", func(_ sobek.FunctionCall) sobek.Value {
		return rt.ToValue("[object Blob]")
	}))
	must(rt, obj.SetPrototype(proto))

	return obj
}
