package summarizejson

import (
	"fmt"
	"regexp"
)

type Replacer interface {
	Do(string) string
}

type replacerNullObject struct{}

func (r *replacerNullObject) Do(s string) string {
	return s
}

type Replacement struct {
	Pattern *regexp.Regexp
	Replace string
}

func (r *Replacement) Do(s string) string {
	return r.Pattern.ReplaceAllString(s, r.Replace)
}

type Summarizer struct {
	Result         map[string]int
	KeyCollapse    Replacer
	RootExpression string
	PathSeparator  string
	ArrayPrefix    string
	ArraySuffix    string
	TypeSeparator  string
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
