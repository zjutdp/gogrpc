package parser

import (
	"regexp"

	"github.com/zjutdp/crawler/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// ParseCityList ...
func ParseCityList(
	contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)

	result := engine.ParserResult{}

	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Items = append(
			result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			URL:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
		// fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
	}
	return result
}
