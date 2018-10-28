package compose

type datastores map[string]struct{}

// IsSupported returns true if the datastore is supported. The is the
// image name only and not the tag
func (d datastores) IsSupported(name string) bool {
	_, ok := d[name]
	return ok
}

// Datastores holds the supported data stores
var Datastores = datastores{
	"postgres":              struct{}{},
	"mysql":                 struct{}{},
	"cockroachdb/cockroach": struct{}{},
	"elasticsearch":         struct{}{},
	"consul":                struct{}{},
	"etcd":                  struct{}{},
	"redis":                 struct{}{},
}
