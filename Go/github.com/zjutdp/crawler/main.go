package main

import (
	"github.com/zjutdp/crawler/engine"
	"github.com/zjutdp/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		URL:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
