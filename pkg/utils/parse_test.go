package utils

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testKVS = []byte(`
key=value
 key1=value1

jsonvalue={"some":"json"}

withequal=name=foo
quote = "value"


`)

func Test_ParseKV(t *testing.T) {
	kvs, err := ParseKeyValuePairs(bytes.NewBuffer(testKVS))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 5, len(kvs))
	assert.Equal(t, `"value"`, kvs["quote"])
	assert.Equal(t, "value", kvs["key"])
	assert.Equal(t, "name=foo", kvs["withequal"])
	assert.Equal(t, "value1", kvs["key1"])
	assert.Equal(t, `{"some":"json"}`, kvs["jsonvalue"])

	t.Log(kvs)

	_, err = ParseKeyValuePairs(bytes.NewBuffer([]byte("foo")))
	assert.NotNil(t, err)
}
