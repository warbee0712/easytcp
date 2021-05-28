package packet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultCodec_Encode(t *testing.T) {
	c := &DefaultCodec{}
	b, err := c.Encode("hello")
	assert.NoError(t, err)
	assert.Equal(t, []byte("hello"), b)
}

func TestDefaultCodec_Decode(t *testing.T) {
	c := &DefaultCodec{}
	data := []byte("hello")
	v, err := c.Decode(data)
	assert.NoError(t, err)
	assert.Equal(t, string(data), v)
}
