package websockets

import (
	"fmt"
	"net/http"

	"github.com/dop251/goja"
)

// wsParams represent the parameters bag for websocket
type wsParams struct {
	headers http.Header
}

// parseWSParams parses the params from the constructor call or returns an error
func parseWSParams(rt *goja.Runtime, raw goja.Value) (*wsParams, error) {
	parsed := &wsParams{
		headers: make(http.Header),
	}

	if raw == nil || goja.IsUndefined(raw) {
		return parsed, nil
	}

	params := raw.ToObject(rt)
	for _, k := range params.Keys() {
		switch k {
		case "headers":
			headersV := params.Get(k)
			if goja.IsUndefined(headersV) || goja.IsNull(headersV) {
				continue
			}
			headersObj := headersV.ToObject(rt)
			if headersObj == nil {
				continue
			}
			for _, key := range headersObj.Keys() {
				parsed.headers.Set(key, headersObj.Get(key).String())
			}
		// TODO: more params
		default:
			return nil, fmt.Errorf("unknown option %s", k)
		}
	}

	return parsed, nil
}
