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

// Element.Render method generates an HTML element. It exists as
// a method to allow usage inside genHTML func
func (e Element) Render() string {

	var htmlElement string

	// The formatting extracted to a function to allow for future
	// modifications based on various arguments, element's children nodes,
	// handlers, styles, etc.
	switch e.Type {
	case "heading":
		htmlElement = headingTemplate(e)
	case "paragraph":
		htmlElement = paragraphTemplate(e)
	case "button":
		htmlElement = buttonTemplate(e)
	case "div":
		htmlElement = divTemplate(e)
	default:
		htmlElement = ""
	}

	return htmlElement
}

func headingTemplate(e Element) string {
	return fmt.Sprintf(`<h%d class="%s">%s</h%d>`, e.Props.Level, e.Class, e.Content, e.Props.Level)
}

func paragraphTemplate(e Element) string {
	return fmt.Sprintf(`<p class="%s">%s</p>`, e.Class, e.Content)
}

func buttonTemplate(e Element) string {
	return fmt.Sprintf(`<button class="%s">%s</button>`, e.Class, e.Content)
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

func genHTML(site Site) (string, error) {

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
				htmlElement = divTemplate(e)
			}

			return template.HTML(htmlElement), nil
		},
	}

	tmpl := template.Must(template.New("base").Funcs(dynamicElementsFunc).ParseFiles("tpl/base.html"))

	var htmlBytes bytes.Buffer
	err := tmpl.Execute(&htmlBytes, &site)
	if err != nil {
		return "", err
	}

	htmlOutput := htmlBytes.String()

	htmlOutput += ""
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

	output, err := genHTML(site)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(output)
}
