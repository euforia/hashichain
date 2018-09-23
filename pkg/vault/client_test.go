package vault

import (
	"testing"

	"github.com/hashicorp/vault/api"
	"github.com/stretchr/testify/assert"
)

const (
	testVaultAddr  = "http://127.0.0.1:8200"
	testVaultToken = "myroot"
)

var testKVS = map[string]map[string]interface{}{
	"key": map[string]interface{}{
		"key1": "value",
	},
	"key/sub": map[string]interface{}{
		"key2": "value",
	},
	"key/level1/sub": map[string]interface{}{
		"key3": "value",
	},
	"key/level1/level2": map[string]interface{}{
		"key4": "value",
	},
}

func Test_Client(t *testing.T) {
	conf := api.DefaultConfig()
	client, err := NewClient(conf, 2, "secret")
	if err != nil {
		t.Fatal(err)
	}
	client.SetAddress(testVaultAddr)
	client.SetToken(testVaultToken)

	for k, v := range testKVS {
		err = client.Set(k, v)
		assert.Nil(t, err)
	}

	cp, err := client.RecursiveGet("key")
	assert.Nil(t, err)
	assert.Equal(t, 4, len(cp))

	list, err := client.List("")
	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(list))

	list, err = client.List("key/level1")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(list))

	err = client.Copy("key", "key_copy")
	assert.Nil(t, err)

	cp, err = client.RecursiveGet("key_copy")
	assert.Nil(t, err)
	assert.Equal(t, 4, len(cp))
}
