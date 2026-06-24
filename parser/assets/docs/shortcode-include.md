---
title: Code Include
tags: [features, demo]
---
Embed source files directly.

## Local File (Full)

{{ include src="examples/hello.go" lang="go" }}

Usage:
`{{` `include src="examples/hello.go" lang="go"` `}}`

## Local File (Selected Lines)

Here is just the main function (lines 5-7):

{{ include src="examples/hello.go" lines="5-7" lang="go" }}

Usage:
`{{` `include src="examples/hello.go" lines="5-7" lang="go"` `}}`

## Remote File

Includes code from a URL (e.g. Go example).

{{ include src="https://raw.githubusercontent.com/golang/example/master/hello/hello.go" lang="go" }}

Usage:
`{{` `include src="https://..." lang="go"` `}}`
