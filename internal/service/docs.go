package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sycret-test-task/internal/model"
)

type DocsGeneratorService struct {
}

func NewDocsGeneratorService() *DocsGeneratorService {
	return &DocsGeneratorService{}
}

func (s *DocsGeneratorService) Generate(input model.Input) (model.Output, error) {
	// get doc by link in input
	// read doc in xml
	// get slice of tags
	// call requests for all tags
	// modify doc
	// save it
	// create response with path

	doc := doSimpleGetRequest(input.URLTemplate)
	fileName := parseAndChange(doc, input)

	return model.Output{URLWord: fileName}, nil
}

func parseAndChange(input string, obj model.Input) string {
	//var fileName string
	rex, err := regexp.Compile(`<\w*?:text field=\".*?\">`)
	if err != nil {
		log.Printf("%s", err.Error())
		return "ti loh"
	}
	res := rex.FindString(input)
	res = getStringInBetween(res, "\"", "\"")
	result, _ := doGetRequest(fmt.Sprintf("https://sycret.ru/service/apigendoc/apigendoc?text=%s&recordid=%d", res, obj.RecordId))
	log.Printf("result: %v", result)
	indexStart := strings.Index(input, res)
	log.Println(indexStart)

	//rexStart, err := regexp.Compile(`<w:t>_</w:t>`)
	//rexEnd, err := regexp.Compile(`</w:t>`)
	res = searchTagAndReplaceWithData(input, indexStart, result.ResultData)
	return res
}

func searchTagAndReplaceWithData(inputStr string, indexStart int, data string) string {
	rex, _ := regexp.Compile(`</ns1:text>`)
	tmpStr := inputStr[indexStart:rex.FindStringIndex(inputStr)[0]]
	log.Println(tmpStr)
	return strings.ReplaceAll(tmpStr, "<w:t>_</w:t>", fmt.Sprintf("<w:t>%s</w:t>", data))
}

func getStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return result
	}
	newS := str[s+len(start):]
	e := strings.Index(newS, end)
	if e == -1 {
		return result
	}
	result = newS[:e]
	return result
}

func doSimpleGetRequest(url string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header = http.Header{
		"User-Agent": {"Mozilla/5.0"},
	}
	resp, _ := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(bytes)
}

func doGetRequest(url string) (model.ReturnedData, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header = http.Header{
		"User-Agent": {"Mozilla/5.0"},
	}
	resp, _ := client.Do(req)
	if err != nil {
		return model.ReturnedData{}, err
	}
	defer resp.Body.Close()

	var data model.ReturnedData
	json.NewDecoder(resp.Body).Decode(&data)

	return data, nil
}
