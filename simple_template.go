package main

import "flag"
import "os"
import "strings"
import "text/template"

func main() {
	flag.Parse()
	tmpl := template.Must(template.ParseFiles(flag.Args()...))

	err := tmpl.Execute(os.Stdout, envs())
	if err != nil {
		panic(err)
	}
}

func envs() map[string]string {
	var result = make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		result[pair[0]] = pair[1]
	}
	return result
}
