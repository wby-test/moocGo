package parser

import (
	"moocGo/craw/singleTask/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<th><a href="(http://album.zhenai.com/u/[0-9]+)" [^>]+>([^<]+)</a></th>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParseResult {
	matchs := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	//为什么直接使用m[2] 会出现解析出来的用户名一样
	for _, m := range matchs {
		name := string(m[2])
		result.Items = append(result.Items, "user  " + name)
		result.Requests = append(result.Requests, engine.Request{string(m[1]), func(c []byte) engine.ParseResult {
			return ParseProfile(c, name)
		},
		})
	}

	matchs = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matchs {
		result.Requests = append(result.Requests, engine.Request{
			Url:		string(m[1]),
			ParserFunc:	ParseCity,
		})
	}
	return result
}