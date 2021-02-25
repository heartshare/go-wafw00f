package util

import (
	"bytes"
	"go-wafw00f/log"
)

func PrintWafResult(items []string) {
	if len(items) == 0 {
		log.Default("WAF Not Detected")
	} else {
		temp := bytes.Buffer{}
		temp.WriteString("WAF : [")
		for i, v := range items {
			if i == len(items)-1 {
				temp.WriteString(v + "]")
			} else {
				temp.WriteString(v + ",")
			}
		}
		log.Default(temp.String())
	}
}
