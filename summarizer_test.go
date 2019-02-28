package summarizejson

import (
	"encoding/json"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

var TestData1 = map[string]interface{}{
	"foo": []interface{}{
		map[string]interface{}{
			"key0": nil,
			"key1": false,
			"key2": true,
		},
		map[string]interface{}{
			"key3": 1,
			"key4": 2,
			"key5": 3.0,
		},
	},
	"bar": []interface{}{
		map[string]interface{}{
			"key6": "",
			"key7": "bar",
		},
	},
}

var ExpectedForTestData1 = map[string]int{
	"$map[string]interface {}":       1,
	".foo$[]interface {}":            1,
	".foo[]$map[string]interface {}": 2,
	".foo[].(key)$<nil>":             1,
	".foo[].(key)$bool":              2,
	".foo[].(key)$int":               2,
	".foo[].(key)$float64":           1,
	".bar$[]interface {}":            1,
	".bar[]$map[string]interface {}": 1,
	".bar[].(key)$string":            2,
}

func NewSummarizer1() *Summarizer {
	return &Summarizer{
		Result: map[string]int{},
		KeyCollapse: &Replacement{
			Pattern: regexp.MustCompile(`\Akey\d+`),
			Replace: "(key)",
		},
	}
}

func TestSummarizerTestData1(t *testing.T) {
	s := NewSummarizer1()
	assert.Equal(t, ExpectedForTestData1, s.Run(TestData1))
}

func TestSummarizerTestData1ViaJSON(t *testing.T) {
	// The number from JSON  is float64 instead of int
	expected := ExpectedForTestData1
	expected[".foo[].(key)$float64"] += expected[".foo[].(key)$int"]
	delete(expected, ".foo[].(key)$int")

	b, err := json.Marshal(TestData1)
	if assert.NoError(t, err) {
		var obj interface{}
		if assert.NoError(t, json.Unmarshal(b, &obj)) {
			s := NewSummarizer1()
			assert.Equal(t, expected, s.Run(obj))
		}
	}
}
