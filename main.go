package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Link struct {
	Href string
	Text string
}

func main() {
	file, err := os.Open("ex4.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		panic(err)
	}
	var linkParsed []Link
	for d := range doc.Descendants() {
		if d.DataAtom == atom.A {
			link := Link{}
			for _, attr := range d.Attr {
				if attr.Key == "href" {
					link.Href = attr.Val
				}
			}
			for a := range d.Descendants() {
				if a.Type == html.TextNode {
					link.Text = fmt.Sprint(link.Text, strings.TrimSpace(a.Data))
				}
			}
			linkParsed = append(linkParsed, link)
		}
	}

	fmt.Println(linkParsed)
}
