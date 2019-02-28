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

var KeyPattern = flag.String("keypattern", "", "Pattern for collapsing keys")
var KeyReplace = flag.String("keyreplace", "{key}", "Replacement for collapsed keys")
var RootExp = flag.String("root-exp", "(ROOT)", "Expression for root object")
var PathSeparator = flag.String("path-separator", ".", "Separator for object attribute")
var ArrayPrefix = flag.String("array-prefix", "", "Prefix for array expression")
var ArraySuffix = flag.String("array-suffix", "[]", "Suffix for array expression")
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

	fmt.Fprintf(os.Stderr, "keypattern: %v\n", KeyPattern)
	if KeyPattern != nil && *KeyPattern != "" {
		ptn, err := regexp.Compile(*KeyPattern)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: Invalid regular expression keypattern: %s because of %v\n", *KeyPattern, err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "keypattern: %s => %v\n", *KeyPattern, ptn)
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

	fmt.Fprintf(os.Stderr, "flag,args %v\n", flag.Args())
	for _, path := range flag.Args() {
		fmt.Fprintf(os.Stderr, "Loafing %s\n", path)

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
