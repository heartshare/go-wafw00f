package waf

import (
	jsoniter "github.com/json-iterator/go"
	"go-wafw00f/log"
	"go-wafw00f/model"
	"go-wafw00f/util"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var Waf model.Waf

func ResolveWafLib() model.Waf {
	filename := "./waf/json/waf.json"
	_, err := os.Stat(filename)
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	if err == nil {
		var ret model.Waf
		log.Info("WAF Json Exist")
		data, _ := ioutil.ReadFile(filename)
		err = json.Unmarshal(data, &ret)
		if err != nil {
			log.Error("WAF Unmarshal Error")
		}
		Waf = ret
		return ret
	} else {
		filenames, _ := util.GetAllFile("./waf/lib")
		for _, v := range filenames {
			content, _ := ioutil.ReadFile(v)
			if strings.HasSuffix(v, ".lib") {
				doParse(string(content))
			}
		}
		data, _ := json.Marshal(Waf)
		log.Info("Create WAF Json File")
		file, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0777)
		if file != nil {
			defer file.Close()
			_, _ = file.Write(data)
		}
		return Waf
	}
}

func doParse(input string) {
	var result model.WafItems

	nameRe := regexp.MustCompile(`.*?NAME[\s]=[\s]'(.*?)'.*?`)
	nameMatch := nameRe.FindAllStringSubmatch(input, -1)
	name := nameMatch[0][1]
	result.Name = name

	headerRe := regexp.MustCompile(`.*?self.matchHeader\(\((.*?)\)\).*?`)
	headerMatch := headerRe.FindAllStringSubmatch(input, -1)
	for _, v := range headerMatch {
		header := v[1]
		key := strings.TrimLeft(strings.Split(header, "',")[0], "'")
		key = strings.TrimSpace(key)
		value := strings.TrimSpace(strings.Split(header, "',")[1])
		value = strings.TrimLeft(value, "r")
		value = strings.Trim(value, "'")
		headerStruct := model.WafHeader{Key: key, ReValue: value}
		result.ReHeaders = append(result.ReHeaders, headerStruct)
	}

	contentRe := regexp.MustCompile(`.*?self.matchContent\(r'(.*?)'\).*?`)
	contentMatch := contentRe.FindAllStringSubmatch(input, -1)
	for _, v := range contentMatch {
		content := v[1]
		result.ReContent = append(result.ReContent, content)
	}

	cookieRe := regexp.MustCompile(`self.matchCookie\(r'(.*?)'\).*?`)
	cookieMatch := cookieRe.FindAllStringSubmatch(input, -1)
	for _, v := range cookieMatch {
		cookie := v[1]
		cookie = strings.TrimLeft(cookie, "r")
		cookie = strings.Trim(cookie, "'")
		result.ReCookies = append(result.ReContent, cookie)
	}

	Waf.Items = append(Waf.Items, result)
}
