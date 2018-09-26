package main

import (
	"flag"
	"log"
	"os"
)

var (
	infile = flag.String("i", "", "Input file")
)

func main() {
	flag.Parse()

	// Parse job spec into struct
	job, err := translate.ReadJobSpecFile(*infile)
	if err != nil {
		log.Fatal(err)
	}

	//
	// The job spec can be mutated here then written out
	// as a valid nomad job spec HCL
	//

	// Write out spec to stdout
	err = translate.WriteJobSpec(os.Stdout, job)
	if err != nil {
		log.Fatal(err)
	}
}
