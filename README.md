# env-template

This tool parses environment variables and makes the values available to be used in a [go template](https://pkg.go.dev/text/template). It supports reading the template from either stdin or a file, making it flexible for various use cases.

## Description

The env-template is a command-line tool written in Go that simplifies the process of rendering templates with environment variables. It accepts environment variables in key-value pairs and populates a template with their values. Additionally, it can handle JSON values within environment variables, parsing them into maps for more complex configurations.

## Features

- Parses environment variables and extracts their values.
- Supports JSON values within environment variables.
- Renders templates with environment variable values.
- Flexible input options: reads templates from either stdin or a file.

## Usage

### Command-Line Arguments

```sh
./env-template <template_file>
```

- `<template_file>`: Path to the template file. Use `-` to read the template from stdin.

### Example

Assuming you have a template file named `example_template.yaml`:

```yaml
# example_template.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: my-config
data:
  {{ range $key, $value := . -}}
  {{ $key }}: "{{ $value }}"
  {{ end -}}
```

You can render the template using the env-template as follows:

```sh
./env-template example_template.yaml
```

Or, you can provide the template via stdin:

```sh
echo "apiVersion: v1\nkind: ConfigMap\nmetadata:\n  name: my-config\ndata:\n  {{ range $key, $value := . }}\n  {{ $key }}: {{ $value }}\n  {{ end }}" | ./env-template -
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
