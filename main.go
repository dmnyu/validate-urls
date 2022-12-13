package main

import (
	"github.com/antchfx/xmlquery"
	"io"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {
	root := os.Args[1]
	xmlFile, err := os.Open(root)
	if err != nil {
		log.Fatal(err)
	}
	xmlBytes, err := io.ReadAll(xmlFile)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := xmlquery.Parse(strings.NewReader(string(xmlBytes)))
	if err != nil {
		panic(err)
	}

	daos := xmlquery.Find(doc, "//dao")
	if len(daos) > 0 {
		for _, dao := range daos {
			href := dao.SelectAttr("xlink:href")
			_, err := url.ParseRequestURI(href)
			if err != nil {
				log.Fatal(err)
			}

		}
	}
	log.Println("No invalid URLS")
}
