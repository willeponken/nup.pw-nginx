package main

import (
	"html/template"
	"os"
)

type TmplErrors []struct {
	StatusCode  string
	Description string
}

func main() {

	errorTemplate, _ := template.ParseFiles("templates/error.tmpl.html")

	errors := TmplErrors{
		{
			StatusCode:  "301",
			Description: "Moved Permanently",
		},
		{
			StatusCode:  "302",
			Description: "Found",
		},
		{
			StatusCode:  "303",
			Description: "See Other",
		},
		{
			StatusCode:  "307",
			Description: "Temporary Redirect",
		},
		{
			StatusCode:  "400",
			Description: "Bad Request",
		},
		{
			StatusCode:  "401",
			Description: "Authorization Required",
		},
		{
			StatusCode:  "402",
			Description: "Payment Required",
		},
		{
			StatusCode:  "403",
			Description: "Forbidden",
		},
		{
			StatusCode:  "404",
			Description: "Not Found",
		},
		{
			StatusCode:  "405",
			Description: "Not Allowed",
		},
		{
			StatusCode:  "406",
			Description: "Not Acceptable",
		},
		{
			StatusCode:  "408",
			Description: "Request Time-out",
		},
		{
			StatusCode:  "409",
			Description: "Conflict",
		},
		{
			StatusCode:  "410",
			Description: "Gone",
		},
		{
			StatusCode:  "411",
			Description: "Length Required",
		},
		{
			StatusCode:  "412",
			Description: "Precondition Failed",
		},
		{
			StatusCode:  "413",
			Description: "Request Entity Too Large",
		},
		{
			StatusCode:  "414",
			Description: "Request-URI Too Large",
		},
		{
			StatusCode:  "415",
			Description: "Unsupported Media Type",
		},
		{
			StatusCode:  "416",
			Description: "Requested Range Not Satisfiable",
		},
		{
			StatusCode:  "494",
			Description: "Request Header Or Cookie Too Large",
		},
		{
			StatusCode:  "495",
			Description: "Certificate Error",
		},
		{
			StatusCode:  "496",
			Description: "No Certificate",
		},
		{
			StatusCode:  "497",
			Description: "Plain HTTP request to HTTPS port",
		},
		{
			StatusCode:  "500",
			Description: "Internal Server Error",
		},
		{
			StatusCode:  "501",
			Description: "Not Implemented",
		},
		{
			StatusCode:  "502",
			Description: "Bad Gateway",
		},
		{
			StatusCode:  "503",
			Description: "Service Temporarily Unavailable",
		},
		{
			StatusCode:  "504",
			Description: "Gateway Time-out",
		},
		{
			StatusCode:  "507",
			Description: "Insufficient Storage",
		},
	}

	for _, err := range errors {
		file, _ := os.Create("internal/error/" + err.StatusCode + ".html")

		errorTemplate.Execute(file, &err)
	}
}
