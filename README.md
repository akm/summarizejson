# summarizejson

Summarize JSON file structure.

---
[![Build Status](https://travis-ci.org/akm/summarizejson.svg?branch=master)](https://travis-ci.org/akm/summarizejson)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/akm/summarizejson)
[![Go Report Card](https://goreportcard.com/badge/github.com/akm/summarizejson)](https://goreportcard.com/report/github.com/akm/summarizejson)


## Install

```
go get github.com/akm/summarizejson/cmd/summarizejson
```

## Usage

```
$ $GOPATH/bin/summarizejson -h
Usage of $GOPATH/bin/summarizejson:
  $ $GOPATH/bin/summarizejson [options] file1[, file2, ....]

options:
  -array-prefix string
    	Prefix for array expression
  -array-suffix string
    	Suffix for array expression (default "[]")
  -key-pattern string
    	Pattern for collapsing keys
  -key-replace string
    	Replacement for collapsed keys (default "{key}")
  -no-header
    	Hide header
  -path-separator string
    	Separator for object attribute (default ".")
  -root-exp string
    	Expression for root object (default "(ROOT)")
  -type-separator string
    	Separator for type expression (default "\t")
```


### Examples

```
$ curl -L https://github.com/nlohmann/json/raw/develop/benchmarks/data/jeopardy/jeopardy.json -o jeopardy.json
$ summarizejson jeopardy.json
PATH	TYPE	COUNT
(ROOT)	[]interface {}	1
(ROOT)[]	map[string]interface {}	216930
(ROOT)[].air_date	string	216930
(ROOT)[].answer	string	216930
(ROOT)[].category	string	216930
(ROOT)[].question	string	216930
(ROOT)[].round	string	216930
(ROOT)[].show_number	string	216930
(ROOT)[].value	<nil>	3634
(ROOT)[].value	string	213296
```

### TYPE in Golang

| type name | meaning |
|----------|----------|
| `interface {}` | Any type of data |
| `map[string]interface {}` | type of object which has named attributes |

### For dynamic attribute names like ID

```
$ curl -L https://github.com/nlohmann/json/raw/develop/benchmarks/data/nativejson-benchmark/citm_catalog.json -o citm_catalog.json
$ summarizejson citm_catalog.json
PATH	TYPE	COUNT
(ROOT)	map[string]interface {}	1
(ROOT).areaNames	map[string]interface {}	1
(ROOT).areaNames.205705993	string	1
(ROOT).areaNames.205705994	string	1
(ROOT).areaNames.205705995	string	1
(ROOT).areaNames.205705996	string	1
(ROOT).areaNames.205705998	string	1
(ROOT).areaNames.205705999	string	1
(ROOT).areaNames.205706000	string	1
(ROOT).areaNames.205706001	string	1
(ROOT).areaNames.205706002	string	1
(ROOT).areaNames.205706003	string	1
(ROOT).areaNames.205706004	string	1
(ROOT).areaNames.205706005	string	1
(ROOT).areaNames.205706006	string	1
(ROOT).areaNames.205706007	string	1
(ROOT).areaNames.205706008	string	1
(ROOT).areaNames.205706009	string	1
(ROOT).areaNames.342752287	string	1
(ROOT).audienceSubCategoryNames	map[string]interface {}	1
(ROOT).audienceSubCategoryNames.337100890	string	1
(ROOT).blockNames	map[string]interface {}	1
(ROOT).events	map[string]interface {}	1
(ROOT).events.138586341	map[string]interface {}	1
(ROOT).events.138586341.description	<nil>	1
(ROOT).events.138586341.id	float64	1

...
```

This file includes lots of numeric ID used as attribute name.
You can summarize dynamic attribute names with `-key-pattern` and `-key-replace` like this:

```
$ summarizejson -key-pattern='\A\d+\z' -key-replace='{ID}' citm_catalog.json
PATH	TYPE	COUNT
(ROOT)	map[string]interface {}	1
(ROOT).areaNames	map[string]interface {}	1
(ROOT).areaNames.{ID}	string	17
(ROOT).audienceSubCategoryNames	map[string]interface {}	1
(ROOT).audienceSubCategoryNames.{ID}	string	1
(ROOT).blockNames	map[string]interface {}	1
(ROOT).events	map[string]interface {}	1
(ROOT).events.{ID}	map[string]interface {}	184
(ROOT).events.{ID}.description	<nil>	184
(ROOT).events.{ID}.id	float64	184
(ROOT).events.{ID}.logo	<nil>	90
(ROOT).events.{ID}.logo	string	94
(ROOT).events.{ID}.name	string	184
(ROOT).events.{ID}.subTopicIds	[]interface {}	184
(ROOT).events.{ID}.subTopicIds[]	float64	611
(ROOT).events.{ID}.subjectCode	<nil>	184
(ROOT).events.{ID}.subtitle	<nil>	184
(ROOT).events.{ID}.topicIds	[]interface {}	184
(ROOT).events.{ID}.topicIds[]	float64	536
(ROOT).performances	[]interface {}	1
(ROOT).performances[]	map[string]interface {}	243
(ROOT).performances[].eventId	float64	243
(ROOT).performances[].id	float64	243
(ROOT).performances[].logo	<nil>	135
(ROOT).performances[].logo	string	108
(ROOT).performances[].name	<nil>	243
(ROOT).performances[].prices	[]interface {}	243
(ROOT).performances[].prices[]	map[string]interface {}	907
(ROOT).performances[].prices[].amount	float64	907
(ROOT).performances[].prices[].audienceSubCategoryId	float64	907
(ROOT).performances[].prices[].seatCategoryId	float64	907
(ROOT).performances[].seatCategories	[]interface {}	243
(ROOT).performances[].seatCategories[]	map[string]interface {}	907
(ROOT).performances[].seatCategories[].areas	[]interface {}	907
(ROOT).performances[].seatCategories[].areas[]	map[string]interface {}	8685
(ROOT).performances[].seatCategories[].areas[].areaId	float64	8685
(ROOT).performances[].seatCategories[].areas[].blockIds	[]interface {}	8685
(ROOT).performances[].seatCategories[].seatCategoryId	float64	907
(ROOT).performances[].seatMapImage	<nil>	243
(ROOT).performances[].start	float64	243
(ROOT).performances[].venueCode	string	243
(ROOT).seatCategoryNames	map[string]interface {}	1
(ROOT).seatCategoryNames.{ID}	string	64
(ROOT).subTopicNames	map[string]interface {}	1
(ROOT).subTopicNames.{ID}	string	19
(ROOT).subjectNames	map[string]interface {}	1
(ROOT).topicNames	map[string]interface {}	1
(ROOT).topicNames.{ID}	string	4
(ROOT).topicSubTopics	map[string]interface {}	1
(ROOT).topicSubTopics.{ID}	[]interface {}	4
(ROOT).topicSubTopics.{ID}[]	float64	19
(ROOT).venueNames	map[string]interface {}	1
(ROOT).venueNames.PLEYEL_PLEYEL	string	1
```

Passing a regular expreesion `\A\d+\z` to `-key-pattern`.
See https://github.com/google/re2/wiki/Syntax for more detail about regular expression.

#### 180 MB JSON file

```
$ curl -L https://github.com/zemirco/sf-city-lots-json/raw/master/citylots.json -o citylots.json
$ time summarizejson citylots.json
PATH	TYPE	COUNT
(ROOT)	map[string]interface {}	1
(ROOT).features	[]interface {}	1
(ROOT).features[]	map[string]interface {}	206560
(ROOT).features[].geometry	<nil>	6
(ROOT).features[].geometry	map[string]interface {}	206554
(ROOT).features[].geometry.coordinates	[]interface {}	206554
(ROOT).features[].geometry.coordinates[]	[]interface {}	206966
(ROOT).features[].geometry.coordinates[][]	[]interface {}	2568513
(ROOT).features[].geometry.coordinates[][][]	[]interface {}	56825
(ROOT).features[].geometry.coordinates[][][]	float64	7704714
(ROOT).features[].geometry.coordinates[][][][]	float64	170475
(ROOT).features[].geometry.type	string	206554
(ROOT).features[].properties	map[string]interface {}	206560
(ROOT).features[].properties.BLKLOT	string	206560
(ROOT).features[].properties.BLOCK_NUM	string	206560
(ROOT).features[].properties.FROM_ST	<nil>	10053
(ROOT).features[].properties.FROM_ST	string	196507
(ROOT).features[].properties.LOT_NUM	string	206560
(ROOT).features[].properties.MAPBLKLOT	string	206560
(ROOT).features[].properties.ODD_EVEN	<nil>	10053
(ROOT).features[].properties.ODD_EVEN	string	196507
(ROOT).features[].properties.STREET	<nil>	10053
(ROOT).features[].properties.STREET	string	196507
(ROOT).features[].properties.ST_TYPE	<nil>	14118
(ROOT).features[].properties.ST_TYPE	string	192442
(ROOT).features[].properties.TO_ST	<nil>	10053
(ROOT).features[].properties.TO_ST	string	196507
(ROOT).features[].type	string	206560
(ROOT).type	string	1
./summarizejson citylots.json  13.54s user 0.59s system 120% cpu 11.683 total
```

## Run test

```
make
```

## Contributing
Bug reports and pull requests are welcome on GitHub at https://github.com/akm/summarizejson.
This project is intended to be a safe, welcoming space for collaboration, and contributors are
expected to adhere to the [Contributor Covenant](http://contributor-covenant.org/) code of conduct.

## License
The gem is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).
