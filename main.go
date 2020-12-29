package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

var (
	flagSilent              = flag.Bool("silent", false, "Do not print the files repalced")
	flagSilentShort         = flag.Bool("s", false, "Do not print the files repalced")
	flagDryRun              = flag.Bool("dry-run", false, "Only print the renames")
	flagDryRunShort         = flag.Bool("p", false, "Only print the renames")
	flagContinueOnDupe      = flag.Bool("continue-on-dupe", false, "Replace files that won't produce a duplicate, even when others will")
	flagContinueOnDupeShort = flag.Bool("c", false, "Replace files that won't produce a duplicate, even when others will")
)

func main() {
	flag.Parse()
	silent := *flagSilent || *flagSilentShort
	dryRun := *flagDryRun || *flagDryRunShort
	continueOnDupe := *flagContinueOnDupe || *flagContinueOnDupeShort
	if flag.NArg() < 2 {
		fmt.Println("regex find and regex replace are missing")
		flag.PrintDefaults()
		os.Exit(1)
	}
	expFind := flag.Arg(0)
	expReplace := flag.Arg(1)

	reg, err := regexp.Compile(expFind)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	files, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var replaces []replacePair
	for _, f := range files {
		pair := replacePair{
			f.Name(),
			reg.ReplaceAllString(f.Name(), expReplace),
		}
		checkName(pair.Replacment)
		if pair.Original != pair.Replacment {
			replaces = append(replaces, pair)
		}
	}

	if duplicateFileNames && !continueOnDupe {
		fmt.Println("No files renamed. Change regex or use -continue-on-dupe (-c) to proceed with non-duplicated names.")
		os.Exit(1)
	}

	for _, r := range replaces {
		if !silent {
			fmt.Printf("'%s' -> '%s'\n", r.Original, r.Replacment)
		}
		if !dryRun {
			os.Rename(r.Original, r.Replacment)
		}
	}

}

var duplicateFileNames = false
var fileNames = make(map[string]bool)

func checkName(name string) {
	if fileNames[name] {
		duplicateFileNames = true
		fmt.Printf("Duplicate file name created: '%s'\n", name)
	}
	fileNames[name] = true
}

type replacePair struct {
	Original   string
	Replacment string
}
