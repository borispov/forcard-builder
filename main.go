package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
)

type Config struct {
	Name       string                 `json:"name"`
	Meta       string                 `json:"meta"`
	Properties map[string]interface{} `json:"properties"`
}

type Site struct {
	Cfg      Config    `json:"site"`
	HtmlBody []Element `json:"elements"`
}

type Element struct {
	Type    string `json:"type"`
	Class   string `json:"class"`
	Content string `json:"content"`
	Props   struct {
		Level      int               `json:"level"`
		Styles     map[string]string `json:"styles"`
		Attributes map[string]string `json:"attributes"`
		Children   []Element         `json:"children"`
	} `json:"props"`
}

func (e Element) Render() string {

	var htmlElement string

	switch e.Type {
	case "heading":
		htmlElement = headingTemplate(e)
	case "paragraph":
		htmlElement = paragraphTemplate(e)
	case "button":
		htmlElement = fmt.Sprintf(`<button class="%s">%s</button>`, e.Class, e.Content)
	case "div":
		htmlElement = divTemplate(e)
	default:
		htmlElement = ""
	}

	return htmlElement
	// return template.HTML(htmlElement)
}

func headingTemplate(e Element) string {
	return fmt.Sprintf(`<h%d class="%s">%s</h%d>`, e.Props.Level, e.Class, e.Content, e.Props.Level)
}

func paragraphTemplate(e Element) string {
	return fmt.Sprintf(`<p class="%s">%s</p>`, e.Class, e.Content)
}

func divTemplate(e Element) string {
	if len(e.Props.Children) == 0 {
		return fmt.Sprintf(`<div class="%s"></div>`, e.Class)
	}

	childrenHTML := "\n"
	for _, c := range e.Props.Children {
		childHTML := c.Render()
		childrenHTML += childHTML + "\n"
	}
	return fmt.Sprintf(`<div class="%s">%s</div>`, e.Class, childrenHTML)
}

func genHTML(elements []Element) (string, error) {

	dynamicElementsFunc := template.FuncMap{
		"element": func(e Element) (template.HTML, error) {
			var htmlElement string

			switch e.Type {
			case "heading":
				htmlElement = headingTemplate(e)
			case "paragraph":
				htmlElement = paragraphTemplate(e)
			case "div":
				htmlElement = divTemplate(e)
			default:
				htmlElement = ""
			}

			return template.HTML(htmlElement), nil
		},
	}

	tmpl := template.Must(template.New("").Funcs(dynamicElementsFunc).Parse(`
		{{ range . }}
			{{ element . }}
		{{ end }}
	`))

	var htmlBytes bytes.Buffer
	err := tmpl.Execute(&htmlBytes, &elements)
	if err != nil {
		return "", err
	}

	htmlOutput := htmlBytes.String()
	return htmlOutput, nil
}

func main() {
	var site Site

	siteFile, err := ioutil.ReadFile("./site.json")
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(siteFile, &site)
	if err != nil {
		fmt.Println(err)
	}

	output, err := genHTML(site.HtmlBody)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(output)
}
