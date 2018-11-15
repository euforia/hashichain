package nomad

import (
	"bytes"
	"io"
	"strings"

	"github.com/aryann/difflib"
	"github.com/hashicorp/nomad/api"
	"github.com/hashicorp/nomad/jobspec"

	"github.com/euforia/hashichain/pkg/nomad/structs"
	"github.com/euforia/hashichain/pkg/nomad/translate/hcl"
	"github.com/euforia/hclencoder"
)

// ReadJobSpecFile reads and parses a nomad job spec from the
// given file
func ReadJobSpecFile(filename string) (*api.Job, error) {
	return jobspec.ParseFile(filename)
}

// ReadJobSpec parses a job spec
func ReadJobSpec(r io.Reader) (*api.Job, error) {
	return jobspec.Parse(r)
}

// WriteJobSpec translates and writes out the job in HCL.  This is useful to
// write out a valid job spec
func WriteJobSpec(w io.Writer, job *api.Job) error {
	// Wrap to our job
	j := hcl.NewJob(job)
	b, err := hclencoder.Encode(map[string]*structs.Job{
		`job "` + j.Name + `"`: j,
	})
	if err == nil {
		_, err = w.Write(b)
	}
	return err
}

// DiffJobSpec returns the diff between to jobs
func DiffJobSpec(orig, job *api.Job) []difflib.DiffRecord {
	oldBuff := bytes.NewBuffer(nil)
	newBuff := bytes.NewBuffer(nil)
	WriteJobSpec(oldBuff, orig)
	WriteJobSpec(newBuff, job)
	t1 := strings.Split(string(oldBuff.Bytes()), "\n")
	t2 := strings.Split(string(newBuff.Bytes()), "\n")
	return difflib.Diff(t1, t2)
}
