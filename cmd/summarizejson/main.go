package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"

	"github.com/akm/summarizejson"
)

var arrayPrefix = flag.String("array-prefix", "", "Prefix for array expression")
var arraySuffix = flag.String("array-suffix", "[]", "Suffix for array expression")
var keyPattern = flag.String("key-pattern", "", "Pattern for collapsing keys")
var keyReplace = flag.String("key-replace", "{key}", "Replacement for collapsed keys")
var noHeader = flag.Bool("no-header", false, "Hide header")
var pathSeparator = flag.String("path-separator", ".", "Separator for object attribute")
var rootExp = flag.String("root-exp", "(ROOT)", "Expression for root object")
var typeSeparator = flag.String("type-separator", "\t", "Separator for type expression")

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

	if keyPattern != nil && *keyPattern != "" {
		ptn, err := regexp.Compile(*keyPattern)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: Invalid regular expression keypattern: %s because of %v\n", *keyPattern, err)
			os.Exit(1)
		}
		s.KeyCollapse = &summarizejson.Replacement{
			Pattern: ptn,
			Replace: *keyReplace,
		}
	}

	if rootExp != nil {
		s.RootExpression = *rootExp
	}
	if pathSeparator != nil {
		s.PathSeparator = *pathSeparator
	}
	if arrayPrefix != nil {
		s.ArrayPrefix = *arrayPrefix
	}
	if arraySuffix != nil {
		s.ArraySuffix = *arraySuffix
	}
	if typeSeparator != nil {
		s.TypeSeparator = *typeSeparator
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

		r := s.Run(obj)

		if noHeader == nil || !(*noHeader) {
			fmt.Fprintf(os.Stdout, "%s%s%s\t%s\n", "PATH", s.TypeSeparator, "TYPE", "COUNT")
		}

		keys := make([]string, len(r))
		i := 0
		for k := range r {
			keys[i] = k
			i += 1
		}
		sort.Strings(keys)

		for _, key := range keys {
			cnt := r[key]
			fmt.Fprintf(os.Stdout, "%s\t%d\n", key, cnt)
		}
	}
}
