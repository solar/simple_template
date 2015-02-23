package main

import (
	"flag"
	"os"
	"strings"
	"text/template"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var values = make(map[string]interface{})

func main() {
	var conf *string = flag.String("yaml", "", "yaml config file path")
	flag.Parse()

	addenvs()
	loadYaml(*conf)

	tmpl := template.Must(template.ParseFiles(flag.Args()...))
	err := tmpl.Execute(os.Stdout, values)
	if err != nil {
		panic(err)
	}
}

func loadYaml(conf string) {
	if conf != "" {
		source, err := ioutil.ReadFile(conf)
		if err != nil {
			panic(err)
		}
		yaml.Unmarshal(source, &values)
	}
}

func addenvs() {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if values[pair[0]] == nil {
			values[pair[0]] = pair[1]
		}
	}
}
