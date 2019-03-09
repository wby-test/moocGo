package engine

import (
	"fmt"
	"log"
	"moocGo/craw/singleTask/fetcher"
)

type SimpleEngine struct {

}


func (SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		/* 单任务版本
		log.Printf("Fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
			continue
		}

		parseResult := r.ParserFunc(body)
		*/

		//并发版本
		parseResult, err := worker(r)
		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)//...将切片所有内容append

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}

	fmt.Println("finish")
}

//并发版
func  worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, nil
	}

	return r.ParserFunc(body), nil
}