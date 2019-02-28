package summarizejson

import (
	"fmt"
	"regexp"
)

type Replacement struct {
	Pattern *regexp.Regexp
	Replace string
}

func (r *Replacement) Do(s string) string {
	return r.Pattern.ReplaceAllString(s, r.Replace)
}

type Summarizer struct {
	Result      map[string]int
	KeyCollapse *Replacement
}

func (s *Summarizer) Load(obj interface{}) {
	s.Walk("", obj)
}

func (s *Summarizer) Walk(path string, obj interface{}) {
	s.Result[fmt.Sprintf("%s$%T", path, obj)] += 1
	switch val := obj.(type) {
	case map[string]interface{}:
		for k, v := range val {
			var newKey string
			if s.KeyCollapse != nil {
				newKey = s.KeyCollapse.Do(k)
			} else {
				newKey = k
			}
			s.Walk(fmt.Sprintf("%s.%s", path, newKey), v)
		}
	case []interface{}:
		key := fmt.Sprintf("%s[]", path)
		for _, v := range val {
			s.Walk(key, v)
		}
	}
}
