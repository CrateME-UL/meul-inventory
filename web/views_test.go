package web

import (
	"bytes"
	"encoding/xml"
	"html/template"
	"io"
	"meul/inventory/internal/infrastructures/drivers/postgres/models"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/google/uuid"
	"golang.org/x/net/html"
	"gotest.tools/assert"
)

func renderIndexView(templateName string, model any) bytes.Buffer {
	templ := template.Must(template.ParseFS(
		templatesFS,
		"templates/views/"+templateName,
		"templates/partials/head.html",
		"templates/partials/navbar.html",
	))

	var buf bytes.Buffer
	err := templ.Execute(&buf, model)
	if err != nil {
		panic(err)
	}
	return buf
}

func renderItemsView(templateName string, model any) bytes.Buffer {
	templ := template.Must(template.ParseFS(
		templatesFS,
		"templates/views/"+templateName,
		"templates/partials/head.html",
		"templates/partials/navbar.html",
		"templates/components/item.html",
	))

	var buf bytes.Buffer
	err := templ.Execute(&buf, model)
	if err != nil {
		panic(err)
	}
	return buf
}

func assertWellFormedHtml(t *testing.T, buf bytes.Buffer) {
	decoder := xml.NewDecoder(bytes.NewReader(buf.Bytes()))
	decoder.Strict = false
	decoder.AutoClose = xml.HTMLAutoClose
	decoder.Entity = xml.HTMLEntity
	for {
		_, err := decoder.Token()
		switch err {
		case io.EOF:
			return // We're done, it's valid!
		case nil:
			// do nothing
		default:
			t.Fatalf("Error parsing html: %s", err)
		}
	}
}

func parseHtml(t *testing.T, buf bytes.Buffer) *goquery.Document {
	assertWellFormedHtml(t, buf)
	document, err := goquery.NewDocumentFromReader(bytes.NewReader(buf.Bytes()))
	if err != nil {
		// if parsing fails, we stop the test here with t.FatalF
		t.Fatalf("Error rendering template %s", err)
	}
	return document
}

func text(node *html.Node) string {
	// A little mess due to the fact that goquery has
	// a .Text() method on Selection but not on html.Node
	sel := goquery.Selection{Nodes: []*html.Node{node}}
	return strings.TrimSpace(sel.Text())
}

func Test_whenParseIndex_thenHTMLIsWellFormed(t *testing.T) {
	buf := renderIndexView("index.html", nil)

	assertWellFormedHtml(t, buf)
}

func Test_whenParseNoItems_thenHTMLIsWellFormed(t *testing.T) {
	buf := renderItemsView("items.html", nil)

	assertWellFormedHtml(t, buf)
}

type ViewData struct {
	Items []models.Item
}

func Test_whenParseWithItems_thenHTMLIsWellFormed(t *testing.T) {
	mockItems := []models.Item{
		{ItemID: 1, ItemNumber: uuid.New(), Name: "that item"},
		{ItemID: 2, ItemNumber: uuid.New(), Name: "another item"},
	}
	viewData := ViewData{
		Items: mockItems,
	}

	buf := renderItemsView("items.html", viewData)

	document := parseHtml(t, buf)
	selection := document.Find("ul.items-list li")
	assert.Equal(t, 2, selection.Length())
	assert.Equal(t, "that item", text(selection.Nodes[0]))
	assert.Equal(t, "another item", text(selection.Nodes[1]))
}
