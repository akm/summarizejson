package summarizejson

import (
	"fmt"
	"regexp"
)

// Replacer replaces string with something
type Replacer interface {
	Do(string) string
}

type replacerNullObject struct{}

func (r *replacerNullObject) Do(s string) string {
	return s
}

// Replacement is a kind of Replacer.
// Replaces a string which matches `Pattern` with `Replace`.
type Replacement struct {
	Pattern *regexp.Regexp
	Replace string
}

// Do returns a string from given string `s`.
func (r *Replacement) Do(s string) string {
	return r.Pattern.ReplaceAllString(s, r.Replace)
}

// Summarizer counts for each `path` in data.
// `path` means a route from the root to a node.
type Summarizer struct {
	// Result is a mamory to count up.
	Result map[string]int

	// KeyCollapse is a Replacer to summarize dynamic attribute names.
	KeyCollapse Replacer

	// RootExpression is used as root of `path`. default is ""
	RootExpression string

	// PathSeparator is separator between node names in `path`. default is "."
	PathSeparator string

	// ArrayPrefix is prefix for array object([]interface{}) in `path`. default is ""
	ArrayPrefix string

	// ArrayPrefix is suffix for array object([]interface{}) in `path`. default is "[]"
	ArraySuffix string

	// TypeSeparator is separator between `path` and object type. default is "$"
	TypeSeparator string
}

func (s *Summarizer) Fulfill() {
	if s.KeyCollapse == nil {
		s.KeyCollapse = &replacerNullObject{}
	}
	// if s.RootExpression == "" {
	// 	s.RootExpression = ""
	// }
	if s.PathSeparator == "" {
		s.PathSeparator = "."
	}
	// if s.ArrayPrefix == "" {
	// 	s.ArrayPrefix = ""
	// }
	if s.ArraySuffix == "" {
		s.ArraySuffix = "[]"
	}
	if s.TypeSeparator == "" {
		s.TypeSeparator = "$"
	}
}

// Load starts count up
func (s *Summarizer) Load(obj interface{}) {
	s.Fulfill()
	s.Walk(s.RootExpression, obj)
}

func (s *Summarizer) Walk(path string, obj interface{}) {
	s.Result[fmt.Sprintf("%s%s%T", path, s.TypeSeparator, obj)] += 1
	switch val := obj.(type) {
	case map[string]interface{}:
		for k, v := range val {
			newKey := s.KeyCollapse.Do(k)
			s.Walk(fmt.Sprintf("%s%s%s", path, s.PathSeparator, newKey), v)
		}
	case []interface{}:
		key := fmt.Sprintf("%s%s%s", s.ArrayPrefix, path, s.ArraySuffix)
		for _, v := range val {
			s.Walk(key, v)
		}
	}
}
