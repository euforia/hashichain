package nomad

import (
	"io"

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
	j := hcl.JobSpec(job)
	b, err := hclencoder.Encode(map[string]*structs.Job{
		`job "` + j.Name + `"`: j,
	})
	if err == nil {
		_, err = w.Write(b)
	}
	return err
}
