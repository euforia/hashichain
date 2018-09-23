package main

import (
	"flag"
	"log"
	"os"

	"github.com/euforia/hashichain/pkg/nomad"
)

var (
	infile = flag.String("i", "", "Input file")
)

func main() {
	flag.Parse()

	// Parse job spec into struct
	job, err := nomad.ReadJobSpecFile(*infile)
	if err != nil {
		log.Fatal(err)
	}

	//
	// The job spec can be mutated here then written out
	// as a valid nomad job spec HCL
	//

	// Write out spec to stdout
	err = nomad.WriteJobSpec(os.Stdout, job)
	if err != nil {
		log.Fatal(err)
	}
}
