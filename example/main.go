package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/euforia/hashichain/pkg/nomad"
	"github.com/euforia/hashichain/pkg/nomad/translate/compose"
)

var (
	infile = flag.String("i", "", "Input file or directory")
)

func getDirFiles(dir string) ([]string, error) {
	list, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	out := make([]string, 0)
	for _, l := range list {
		name := l.Name()
		ext := filepath.Ext(name)
		switch ext[1:] {
		case "yml", "yaml":
			out = append(out, filepath.Join(dir, name))

		}
	}
	return out, nil
}

func getInputFiles() []string {
	stat, err := os.Stat(*infile)
	if err != nil {
		log.Fatal(err)
	}
	var filelist []string
	if stat.IsDir() {
		filelist, err = getDirFiles(*infile)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		filelist = []string{*infile}
	}

	return filelist
}

func main() {
	flag.Parse()

	filelist := getInputFiles()

	// Convert compose file into a new nomad job
	job, err := compose.NewJob(".", nil, filelist...)
	if err != nil {
		log.Fatal(err)
	}

	// Write out spec to stdout
	err = nomad.WriteJobSpec(os.Stdout, job)
	if err != nil {
		log.Fatal(err)
	}
}
