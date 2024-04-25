package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

/*
Parse environment variables and extract their values,
attempt to unmarshal JSON strings into maps if applicable, and return
them as a map[string]interface{}.
*/
func parseEnvVars() map[string]interface{} {
	variables := make(map[string]interface{})

	for _, env := range os.Environ() {
		key, value, found := strings.Cut(env, "=")

		if found {
			var parsedJSON map[string]string
			err := json.Unmarshal([]byte(value), &parsedJSON)

			if err == nil {
				variables[key] = parsedJSON
				continue
			}

			variables[key] = value
		}
	}

	return variables
}

func exitIfError(err error, reason string, args ...any) {
	if err != nil {
		fmt.Fprintf(os.Stderr, reason, args...)
		os.Exit(1)
	}
}

func main() {
	var templateFilePath = "-"

	if len(os.Args) >= 2 {
		templateFilePath = os.Args[1]
	}

	var tmpl *template.Template
	var err error

	if templateFilePath == "-" {
		// Read template from stdin if no file is specified.
		reader := bufio.NewReader(os.Stdin)
		templateBytes, err := io.ReadAll(reader)
		exitIfError(err, "Error reading template from stdin: %v\n", err)

		tmpl, err = template.New("stdin").Parse(string(templateBytes))
		exitIfError(err, "Error parsing template from stdin: %v\n", err)
	} else {
		// Read template from file
		templateBytes, err := os.ReadFile(templateFilePath)
		exitIfError(err, "Error reading template file: %v\n", err)

		tmpl, err = template.New(templateFilePath).Parse(string(templateBytes))
		exitIfError(err, "Error parsing template file: %v\n", err)
	}

	variables := parseEnvVars()

	err = tmpl.Execute(os.Stdout, variables)
	exitIfError(err, "Error rendering template: %v\n", err)
}
