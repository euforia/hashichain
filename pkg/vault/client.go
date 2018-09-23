package vault

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/hashicorp/vault/api"
)

const defaultTimeout = 3 * time.Second

var (
	vmap = map[uint8]map[string]string{
		1: map[string]string{
			"get":  "",
			"list": "",
		},
		2: map[string]string{
			"get":  "data",
			"list": "metadata",
		},
	}
)

// CopyOptions are key value pair copy options
type CopyOptions struct {
	KeysOnly bool
}

// Client is a vault client with helper functions
type Client struct {
	// vault mount point
	prefix    string
	kvVersion uint8
	*api.Client
}

// NewClient returns a new Client instance
func NewClient(conf *api.Config, version uint8, prefix string) (*Client, error) {
	conf.Timeout = defaultTimeout
	client, err := api.NewClient(conf)
	if err != nil {
		return nil, err
	}

	c := &Client{
		prefix:    prefix,
		kvVersion: version,
		Client:    client,
	}

	return c, nil
}

// Copy recursively copies the source path to the destination path
func (c *Client) Copy(src, dst string, opt ...CopyOptions) error {
	mm, err := c.RecursiveGet(src)
	if err != nil {
		return err
	}

	if len(opt) > 0 && opt[0].KeysOnly {
		return c.copyKeys(src, dst, mm)
	}

	for k, v := range mm {
		baseKey := strings.TrimPrefix(k, src)
		newKey := filepath.Join(dst, baseKey)
		err = c.Set(newKey, v)
		if err != nil {
			break
		}
	}

	return err
}

// Delete deletes a key. If v2 then only the data is removed
func (c *Client) Delete(key string) error {
	vlt := c.Client.Logical()
	path := c.getGetPath(key)
	_, err := vlt.Delete(path)
	return err
}

func (c *Client) copyKeys(srcpath, dst string, mm map[string]map[string]interface{}) error {
	var err error

	for k, kvs := range mm {
		if len(kvs) == 0 {
			continue
		}

		baseKey := strings.TrimPrefix(k, srcpath)
		newKey := filepath.Join(dst, baseKey)

		for k := range kvs {
			kvs[k] = ""
		}
		err = c.Set(newKey, kvs)
		if err != nil {
			break
		}
	}

	return err
}

// RecursiveGet recursively gets all key value pairs under the startPath
func (c *Client) RecursiveGet(startPath string) (map[string]map[string]interface{}, error) {
	return c.recursiveGet(startPath)
}

func (c *Client) recursiveGet(startPath string) (map[string]map[string]interface{}, error) {
	out := make(map[string]map[string]interface{})
	kv, err := c.Get(startPath)
	if err == nil {
		out[startPath] = kv
	}

	keys, err := c.List(startPath + "/")
	if err != nil {
		fmt.Println(err)
		return out, nil
	}

	for _, key := range keys {
		k := filepath.Join(startPath, key)

		if key[len(key)-1] != '/' {
			kv, err := c.Get(k)
			if err != nil {
				fmt.Println(err)
				continue
			}
			out[k] = kv

		} else {
			kvs, err := c.recursiveGet(k)
			if err != nil {
				fmt.Println(err)
				continue
			}
			for k, v := range kvs {
				out[k] = v
			}
		}
	}

	return out, nil
}

// Set sets the key to the given value
func (c *Client) Set(key string, value map[string]interface{}) error {
	var (
		path = c.getGetPath(key)
		val  map[string]interface{}
	)

	if c.kvVersion == 2 {
		val = map[string]interface{}{"data": value}
	} else {
		val = value
	}

	vlt := c.Client.Logical()
	_, err := vlt.Write(path, val)
	return err
}

// Update update the key with the value.  It only considers the first level adds/overrides level1
// keys only
func (c *Client) Update(key string, value map[string]interface{}) (map[string]interface{}, error) {
	data, err := c.Get(key)
	if err != nil {
		return nil, err
	}
	for k, v := range value {
		data[k] = v
	}
	err = c.Set(key, data)
	return data, err
}

// Get returns the value for the given key
func (c *Client) Get(key string) (map[string]interface{}, error) {
	vlt := c.Client.Logical()
	resp, err := vlt.Read(c.getGetPath(key))
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, fmt.Errorf("key not found: %s", key)
	}

	if c.kvVersion == 1 {
		return resp.Data, nil
	}

	// Assume v2
	v := resp.Data["data"]
	if v == nil {
		// May need to re-visit
		return nil, fmt.Errorf("key not found: %s", key)
	}

	d, ok := v.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("data not a map: %v", v)
	}

	return d, nil
}

// List recursively lists all keys under a path
func (c *Client) List(prefix string) ([]string, error) {
	vlt := c.Client.Logical()
	resp, err := vlt.List(c.getListPath(prefix))
	if err != nil {
		return nil, err
	}

	list, _ := extractListData(resp)
	return list, nil
}

func (c *Client) getGetPath(key string) string {
	p := filepath.Join(c.prefix, vmap[c.kvVersion]["get"], key)
	if key[len(key)-1] == '/' {
		p += "/"
	}
	return p
}

func (c *Client) getListPath(key string) string {
	return filepath.Join(c.prefix, vmap[c.kvVersion]["list"], key)
}

func extractListData(secret *api.Secret) ([]string, bool) {
	if secret == nil || secret.Data == nil {
		return nil, false
	}

	k, ok := secret.Data["keys"]
	if !ok || k == nil {
		return nil, false
	}

	i, ok := k.([]interface{})
	if !ok {
		return nil, false
	}

	out := make([]string, len(i))
	for j, v := range i {
		out[j] = v.(string)
	}
	return out, true
}
