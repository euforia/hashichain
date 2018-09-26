package hcl

import (
	"bytes"
	"testing"

	"github.com/hashicorp/nomad/jobspec"

	"github.com/euforia/hashichain/pkg/nomad/structs"
	"github.com/euforia/hclencoder"
)

const testFile = "test-fixtures/example.nomad"

func Test_Translator(t *testing.T) {
	job, err := jobspec.ParseFile(testFile)
	if err != nil {
		t.Fatal(err)
	}

	j := JobSpec(job)

	b, err := hclencoder.Encode(map[string]*structs.Job{`job "` + j.Name + `"`: j})
	if err != nil {
		t.Fatal(err)
	}

	_, err = jobspec.Parse(bytes.NewBuffer(b))
	if err != nil {
		t.Fatal(err)
	}

}
