package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/akm/summarizejson"
)

var ArrayPrefix = flag.String("array-prefix", "", "Prefix for array expression")
var ArraySuffix = flag.String("array-suffix", "[]", "Suffix for array expression")
var KeyPattern = flag.String("keypattern", "", "Pattern for collapsing keys")
var KeyReplace = flag.String("keyreplace", "{key}", "Replacement for collapsed keys")
var PathSeparator = flag.String("path-separator", ".", "Separator for object attribute")
var RootExp = flag.String("root-exp", "(ROOT)", "Expression for root object")
var TypeSeparator = flag.String("type-separator", "\t", "Separator for type expression")

func init() {
	flag.Usage = func() {
		out := os.Stderr
		fmt.Fprintf(out, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(out, "  $ %s [options] file1[, file2, ....]\n", os.Args[0])
		fmt.Fprintf(out, "\noptions:\n")

		flag.PrintDefaults()
	}
}

func newSummarizer() *summarizejson.Summarizer {
	s := &summarizejson.Summarizer{
		Result: map[string]int{},
	}

	if KeyPattern != nil && *KeyPattern != "" {
		ptn, err := regexp.Compile(*KeyPattern)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: Invalid regular expression keypattern: %s because of %v\n", *KeyPattern, err)
			os.Exit(1)
		}
		s.KeyCollapse = &summarizejson.Replacement{
			Pattern: ptn,
			Replace: *KeyReplace,
		}
	}

	if RootExp != nil {
		s.RootExpression = *RootExp
	}
	if PathSeparator != nil {
		s.PathSeparator = *PathSeparator
	}
	if ArrayPrefix != nil {
		s.ArrayPrefix = *ArrayPrefix
	}
	if ArraySuffix != nil {
		s.ArraySuffix = *ArraySuffix
	}
	if TypeSeparator != nil {
		s.TypeSeparator = *TypeSeparator
	}

	return s
}

func main() {
	flag.Parse()

	s := newSummarizer()

	for _, path := range flag.Args() {
		b, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: Failed to read %s because of %v\n", path, err)
			os.Exit(1)
		}

		var obj interface{}
		{
			err := json.Unmarshal(b, &obj)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: Failed to parse %s as JSON because of %v\n", path, err)
				os.Exit(1)
			}
		}

		s.Load(obj)

		for key, cnt := range s.Result {
			fmt.Fprintf(os.Stdout, "%s\t%d\n", key, cnt)
		}
	}
}
