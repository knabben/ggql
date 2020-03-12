package parser

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"net/http/httptest"

)


var (
	file = "../tests/schema.json"
)

func TestParseFileSchema(t *testing.T) {
	parser := NewParser(file, "http://")
	schema, err := parser.parseFileSchema()

	assert.Nil(t, err)
	assert.Equal(t, len(schema.Schema.Types), 22)
	assert.Equal(t, schema.Schema.QueryType.Name, "RootQueryType")
}


func TestParseWrongFileSchema(t *testing.T) {
	parser := NewParser("wrong.json", "http://")
	schema, err := parser.parseFileSchema()

	assert.Error(t, err)
	assert.Equal(t, schema.Schema.QueryType.Name, "")
}


func TestWrongParseURLSchema(t *testing.T) {
	parser := NewParser("wrong.json", "http://")
	schema, err := parser.parseURLSchema()

	assert.Equal(t, err.Error(), "Post http:: http: no Host in request URL")
	assert.NotNil(t, schema)
}


func TestParseURLSchema(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			schema, _ := ioutil.ReadFile(file)
			fmt.Fprintln(w, string(schema))
		}),
	)
	defer ts.Close()

	parser := NewParser("wrong.json", ts.URL)
	schema, err := parser.parseURLSchema()

	assert.Nil(t, err)
	assert.Equal(t, len(schema.Schema.Types), 22)
	assert.Equal(t, schema.Schema.QueryType.Name, "RootQueryType")
}