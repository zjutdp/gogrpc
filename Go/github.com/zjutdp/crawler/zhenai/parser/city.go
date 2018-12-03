package parser

import (
	"regexp"

	"github.com/zjutdp/crawler/engine"
)

const cityRe = ``

func ParseCity(
	contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)

	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		result.Items = append(
			result.Items, "User "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			URL:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}

	return result
}
