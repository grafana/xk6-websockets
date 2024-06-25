package websockets

import (
	"testing"

	"github.com/grafana/sobek"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBlob_text(t *testing.T) {
	t.Parallel()
	ts := newTestState(t)
	val, err := ts.runtime.RunOnEventLoop(`
		const blob = new Blob(["P", "A", "SS"]);
		blob.text();
	`)
	require.NoError(t, err)
	require.Equal(t, "PASS", val.String())
}

func TestBlob_arrayBuffer(t *testing.T) {
	t.Parallel()
	ts := newTestState(t)
	val, err := ts.runtime.RunOnEventLoop(`
		const blob = new Blob(["P", "A", "SS"]);
		blob.arrayBuffer();
	`)
	require.NoError(t, err)

	ab, ok := val.Export().(sobek.ArrayBuffer)
	require.True(t, ok)
	require.Equal(t, "PASS", string(ab.Bytes()))
}

func TestBlob_stream(t *testing.T) {
	t.Parallel()
	ts := newTestState(t)
	val, err := ts.runtime.RunOnEventLoop(`
		(async () => {
		  const blob = new Blob(["P", "A", "SS"]);
		  const reader = blob.stream().getReader();
		  const {value} = await reader.read();
		  return value;
		})()
	`)
	require.NoError(t, err)

	p, ok := val.Export().(*sobek.Promise)
	require.True(t, ok)
	assert.Equal(t, "PASS", p.Result().String())
}
