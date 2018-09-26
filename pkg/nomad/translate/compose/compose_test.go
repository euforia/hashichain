package compose

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/hashicorp/nomad/jobspec"

	"github.com/euforia/hashichain/pkg/nomad"
)

var testComposeFiles = []string{
	"test-fixtures/service.yml",
	"test-fixtures/db.yml",
}

func Test_Compose(t *testing.T) {
	job, err := NewJob(".", nil, testComposeFiles...)
	if err != nil {
		t.Fatal(err)
	}

	buf := bytes.NewBuffer(nil)
	err = nomad.WriteJobSpec(buf, job)
	if err != nil {
		t.Fatal(err)
	}

	out := make([]byte, buf.Len())
	copy(out, buf.Bytes())

	_, err = jobspec.Parse(buf)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%s\n", out)
}
