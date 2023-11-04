package main

import (
	"bytes"
	"context"
	"html/template"
	"math/rand"
	"strings"
)

var helloNativeTemplate = template.Must(template.ParseFiles("hello.html"))
var tableNativeTemplate = template.Must(template.ParseFiles("table.html"))
var layoutNativeTemplate = template.Must(template.ParseFiles("layout.html", "table.html"))

type HelloData struct {
	Name string
}

type TableItem struct {
	Id      string
	Name    string
	Number  int
	Complex *TableItem
}
type Table struct {
	TableName string
	Items     []TableItem
}

var alphabet []rune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

func randomString(length int) string {
	alphabetSize := len(alphabet)
	var sb strings.Builder

	for i := 0; i < length; i++ {
		ch := alphabet[rand.Intn(alphabetSize)]
		sb.WriteRune(ch)
	}

	s := sb.String()
	return s
}

func generateTestTable() Table {
	var items []TableItem

	for i := 0; i < 1000; i++ {
		items = append(items, TableItem{
			Id:     randomString(36),
			Name:   randomString(rand.Intn(500)),
			Number: rand.Intn(99999),
			Complex: &TableItem{
				Id:     randomString(36),
				Name:   randomString(rand.Intn(500)),
				Number: rand.Intn(99999),
			},
		})
	}
	return Table{
		TableName: "test-table",
		Items:     items,
	}
}

var testTable = generateTestTable()

func helloTempl() string {
	component := helloStruct(HelloData{Name: "world"})
	buf := new(bytes.Buffer)

	component.Render(context.Background(), buf)

	return buf.String()
}

func helloNative() string {
	buf := new(bytes.Buffer)
	helloNativeTemplate.Execute(buf, HelloData{Name: "world"})
	return buf.String()
}

func tableNative() string {
	buf := new(bytes.Buffer)
	tableNativeTemplate.ExecuteTemplate(buf, "table", testTable)
	return buf.String()
}

func tableTempl() string {
	component := table(testTable)
	buf := new(bytes.Buffer)

	component.Render(context.Background(), buf)

	return buf.String()
}

func layoutNative() string {
	buf := new(bytes.Buffer)
	layoutNativeTemplate.ExecuteTemplate(buf, "layout", testTable)
	return buf.String()
}

func layoutTempl() string {
	t := table(testTable)
	component := layout(t)
	buf := new(bytes.Buffer)

	component.Render(context.Background(), buf)

	return buf.String()
}
