package websockets

import (
	"bytes"

	"github.com/grafana/sobek"

	"go.k6.io/k6/js/common"
)

type Blob struct {
	typ  string
	data bytes.Buffer
}

func (b *Blob) text() string {
	return b.data.String()
}

func (r *WebSocketsAPI) blob(call sobek.ConstructorCall) *sobek.Object {
	var blobParts []interface{}
	if len(call.Arguments) > 0 {
		if parts, ok := call.Arguments[0].Export().([]interface{}); ok {
			blobParts = parts
		}
	}

	return newBlob(r.vu.Runtime(), blobParts)
}

func newBlob(rt *sobek.Runtime, blobParts []interface{}) *sobek.Object {
	b := &Blob{}
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
				common.Throw(rt, err)
			}
		}
	}

	obj := rt.NewObject()

	if err := obj.DefineAccessorProperty("type", rt.ToValue(func() sobek.Value {
		return rt.ToValue(b.typ)
	}), nil, sobek.FLAG_FALSE, sobek.FLAG_TRUE); err != nil {
		common.Throw(rt, err)
	}

	if err := obj.DefineAccessorProperty("size", rt.ToValue(func() sobek.Value {
		return rt.ToValue(b.data.Len())
	}), nil, sobek.FLAG_FALSE, sobek.FLAG_TRUE); err != nil {
		common.Throw(rt, err)
	}

	if err := obj.Set("text", func(call sobek.FunctionCall) sobek.Value {
		return rt.ToValue(b.text())
	}); err != nil {
		common.Throw(rt, err)
	}

	if err := obj.Set("arrayBuffer", func(call sobek.FunctionCall) sobek.Value {
		return rt.ToValue(rt.NewArrayBuffer(b.data.Bytes()))
	}); err != nil {
		common.Throw(rt, err)
	}

	return obj
}
