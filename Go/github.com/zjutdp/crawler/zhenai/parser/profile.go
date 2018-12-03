package parser

import (
	"regexp"
	"strconv"

	"github.com/zjutdp/crawler/engine"
	"github.com/zjutdp/crawler/model"
)

var ageRe = regexp.MustCompile(
	`<dt><span class="label">年龄: </span>([\d]+)岁</td>`)

var marriageRe = regexp.MustCompile(
	`<dt><span class="label">婚况: </span>([^<]+)</td>`)

func ParseProfile(
	contents []byte) engine.ParseResult {

	profile := model.Profile{}
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}

	profile.Marriage = extractString(contents, marriageRe)

	return engine.ParseResult{
		Items: []interface{}{profile},
	}
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
