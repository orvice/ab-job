package main

import "github.com/orvice/ab-job/biz"

var tgs = make([]string, 0)

func main() {
	go biz.Init()
	biz.Web()
}
