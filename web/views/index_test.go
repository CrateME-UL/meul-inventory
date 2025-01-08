package templates

import (
	"bytes"
	"encoding/xml"
	"html/template"
	"io"
	"meul/inventory/internal/infrastructures/drivers/postgres/models"
	"meul/inventory/test/fixtures"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func renderTemplate(templateName string, model any) bytes.Buffer {
	templ := template.Must(template.ParseFiles(templateName))
	var buf bytes.Buffer
	err := templ.Execute(&buf, model)
	if err != nil {
		panic(err)
	}
	return buf
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

func text(node *html.Node) string {
	// A little mess due to the fact that goquery has
	// a .Text() method on Selection but not on html.Node
	sel := goquery.Selection{Nodes: []*html.Node{node}}
	return strings.TrimSpace(sel.Text())
}
func Test_wellFormedHtml(t *testing.T) {
	// Create mock items
	mockItems := []models.Item{
		{ItemID: 1, ItemNumber: uuid.New(), Name: "that item"},
		{ItemID: 2, ItemNumber: uuid.New(), Name: "another item"},
	}

	// Mock the DAO
	// mockDAO := &models.MockItemDAO{}
	// mockDAO.On("GetAllItems").Return(mockItems, nil)

	buf := renderTemplate("index.html", mockItems)

	assertWellFormedHtml(t, buf)
}

func Test_todoItemsAreShown(t *testing.T) {
	// Create mock items
	mockItems := []models.Item{
		{ItemID: 1, ItemNumber: uuid.New(), Name: "that item"},
		{ItemID: 2, ItemNumber: uuid.New(), Name: "another item"},
	}

	// Mock the DAO
	// mockDAO := &models.MockItemDAO{}
	// mockDAO.On("GetAllItems").Return(mockItems, nil)

	// Render the template with items from the DAO
	buf := renderTemplate("index.html", mockItems)

	document := parseHtml(t, buf)
	// Assert there are two <li> elements inside the <ul class="todo-list">
	selection := document.Find("ul.todo-list li")
	assert.Equal(t, 2, selection.Length())

	// Assert the first <li> text is "that item"
	assert.Equal(t, "that item", text(selection.Nodes[0]))

	// Assert the second <li> text is "another item"
	assert.Equal(t, "another item", text(selection.Nodes[1]))

	// Verify mock expectations
	// mockDAO.AssertExpectations(t)
}

func Test_selectedItemsGetSelectedClass(t *testing.T) {
	mockItems := []models.Item{
		fixtures.NewItemFixture(
			fixtures.WithName("that item"),
			fixtures.WithSelected(true),
		),
		fixtures.NewItemFixture(
			fixtures.WithName("another item"),
			fixtures.WithSelected(false),
		),
	}

	buf := renderTemplate("index.html", mockItems)

	document := parseHtml(t, buf)
	selection := document.Find("ul.todo-list li.selected")
	assert.Equal(t, 1, selection.Size())
	assert.Equal(t, "that item", text(selection.Nodes[0]))
}
