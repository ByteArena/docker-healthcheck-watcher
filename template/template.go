package template

import (
	"bytes"
	"io/ioutil"
	"log"
	"text/template"

	"github.com/bytearena/docker-healthcheck-watcher/types"
)

func getTemplateContent(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func MakeTemplate(s types.ErrorMessage) string {

	//create a new template with some name
	tmpl := template.New("test")
	content, getTemplateErr := getTemplateContent("../../template/error")

	if getTemplateErr != nil {
		log.Panicln(getTemplateErr)
	}

	//parse some content and generate a template
	tmpl, err := tmpl.Parse(string(content))
	if err != nil {
		log.Panicln("Parse: ", err)
	}

	var res bytes.Buffer

	//merge template 'tmpl' with content of 's'
	err1 := tmpl.Execute(&res, s)
	if err1 != nil {
		log.Panicln("Execute: ", err1)
	}

	return res.String()
}
