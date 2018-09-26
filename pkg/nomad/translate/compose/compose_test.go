package compose

import (
	"os"
	"testing"

	"github.com/euforia/cchain/pkg/compose"
	"github.com/euforia/hashichain/pkg/nomad"
)

var testComposeFiles = []string{
	"test-fixtures/service.yml",
	"test-fixtures/db.yml",
}

func Test_Compose(t *testing.T) {
	c, err := compose.NewCompose(".", nil, testComposeFiles...)
	if err != nil {
		t.Fatal(err)
	}

	job, err := translate(c.Config())
	if err != nil {
		t.Fatal(err)
	}

	err = nomad.WriteJobSpec(os.Stdout, job)
	if err != nil {
		t.Fatal(err)
	}
}
