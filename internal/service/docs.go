package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sycret-test-task/internal/model"
	"time"
)

type DocsGeneratorService struct {
}

func NewDocsGeneratorService() *DocsGeneratorService {
	return &DocsGeneratorService{}
}

func (s *DocsGeneratorService) Generate(input model.Input) (model.Output, error) {
	// getting xml file as string
	doc := doSimpleGetRequest(input.URLTemplate)

	// creating regular expression for searching desired tags
	rex, err := regexp.Compile(`<\w*?:text field=\".*?\">.*?</\w*?:text>`)
	if err != nil {
		log.Printf("%s", err.Error())
	}

	// creating array of all <text>...</text> entries
	tagEntries := getTagsArray(doc, rex)

	// getting modified array where <w:t>_</w:t> replaced with <w:t>some new text from API etc.</w:t>
	modifiedTags := modifyTags(tagEntries, input.RecordId)

	// modifying doc
	modifiedDoc := modifyDoc(doc, tagEntries, modifiedTags)

	// saving doc
	filePath := saveAsMSDoc(modifiedDoc)

	return model.Output{URLWord: filePath}, nil
}

func saveAsMSDoc(doc string) string {
	filePath := "C:/test/" + time.Now().Format("2006-01-02 15-04-05") + ".doc"
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Fprintf(file, "%v", doc)
	file.Close()
	return filePath
}

func modifyDoc(doc string, old []string, new []string) string {
	for i, str := range old {
		doc = strings.ReplaceAll(doc, str, new[i])
	}
	return doc
}

func modifyTags(tags []string, recordId int64) []string {
	var modifiedTags []string

	rex, err := regexp.Compile(`<\w*?:text field=\".*?\">`)
	if err != nil {
		log.Printf("%s", err.Error())
		return nil
	}

	for _, tag := range tags {
		field := getStringInBetween(rex.FindString(tag), "\"", "\"")
		data := getDataByField(field, recordId)
		t := strings.ReplaceAll(tag, "<w:t>_</w:t>", fmt.Sprintf("<w:t>%s</w:t>", data))
		modifiedTags = append(modifiedTags, t)
	}

	return modifiedTags
}

func getDataByField(field string, recordId int64) string {
	result, err := doGetRequest(fmt.Sprintf("https://sycret.ru/service/apigendoc/apigendoc?text=%s&recordid=%d", field, recordId))
	if err != nil {
		log.Printf("%s", err.Error())
	}
	return result.ResultData
}

func getTagsArray(str string, rex *regexp.Regexp) []string {
	return rex.FindAllString(str, -1)
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
