package hcl

import (
	"bytes"
	"testing"

	"github.com/hashicorp/nomad/jobspec"

	"github.com/euforia/hashichain/pkg/nomad/structs"
	"github.com/euforia/hclencoder"
)

const (
	testFile  = "test-fixtures/example.nomad"
	testFile2 = "test-fixtures/nomad09.nomad"
)

func Test_Translator(t *testing.T) {
	job, err := jobspec.ParseFile(testFile)
	if err != nil {
		t.Fatal(err)
	}

	j := NewJob(job)

	b, err := hclencoder.Encode(map[string]*structs.Job{`job "` + j.Name + `"`: j})
	if err != nil {
		t.Fatal(err)
	}

	_, err = jobspec.Parse(bytes.NewBuffer(b))
	if err != nil {
		t.Fatal(err)
	}

}

func Test_Translator2(t *testing.T) {
	job, err := jobspec.ParseFile(testFile2)
	if err != nil {
		t.Fatal(err)
	}

	j := NewJob(job)

	b, err := hclencoder.Encode(map[string]*structs.Job{`job "` + j.Name + `"`: j})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s\n", b)
	_, err = jobspec.Parse(bytes.NewBuffer(b))
	if err != nil {
		t.Fatal(err)
	}
}
