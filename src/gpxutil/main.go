package main

import (
	"log"
	"flag"
	"io/ioutil"
	"os"
	"github.com/tkrajina/gpxgo/gpx"
	"fmt"
	"errors"
	"gpxutil/context"
	"gpxutil/terminal"
)

var (
	version = "0.1"
)

func welcomeToTerminal() {
	fmt.Printf("Welcome to console! Version: %s\n", version)
}

func tryParseFilename() (string, error) {
	filename := flag.String("f", "", "filename")
	flag.Parse()

	if *filename == "" {
		return "", errors.New("cannot parse flag -f")
	}

	return *filename, nil
}

func getGPX(filename string) ([]*gpx.GPX, error) {
	r, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("bad file %s, error %v", filename, err)
	}

	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("cannot read file %s, error %v", filename, err)
	}

	gpxFile, err := gpx.ParseBytes(bytes)
	if err != nil {
		return nil, fmt.Errorf("bad file %s, error %v", filename, err)
	}

	return []*gpx.GPX{gpxFile}, nil
}

func main() {
	filename, err := tryParseFilename()
	if err != nil {
		log.Panicf("%v", err)
	}

	gpxs, err := getGPX(filename)
	if err != nil {
		log.Panicf("%v", err)
	}

	gctx, err := context.InitFromGPXS(gpxs)
	if err != nil {
		log.Panicf("%v", err)
	}

	errs := make(chan error)

	welcomeToTerminal()
	go func() {
		errs <- terminal.InitTerminal(gctx)
	}()

	for err := range errs {
		log.Panicf("error from terminal, %v", err)
	}
}
