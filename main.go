package main

import (
	"go-wafw00f/input"
	"go-wafw00f/log"
	"go-wafw00f/util"
	"go-wafw00f/waf"
	"strings"
)

func main() {
	util.PrintLogo()
	params := input.ParseInput()
	if strings.TrimSpace(params.Url) == "" {
		log.Error("Need Url")
		return
	}
	log.Default("url : " + params.Url)
	log.Default("Start Go-wafw00f")
	waf.ResolveWafLib()
	res := waf.DetectWaf("http://www.lxxcpx.com/t")
	util.PrintWafResult(res)
}
