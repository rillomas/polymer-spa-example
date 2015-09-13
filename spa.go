package spa

import (
	"appengine"
	"fmt"
	"html/template"
	"net/http"
)

const (
	leftTemplateDelimiter  string = "[["
	rightTemplateDelimiter string = "]]"
	rootDirectory          string = "web/web"
	indexFileName          string = "index.html"
)

// Create a template with a custom delimiter
// because the default delimiter interferes with polymer's templating
var xt *template.Template = template.Must(
	template.New("spa").
		Delims(leftTemplateDelimiter, rightTemplateDelimiter).
		ParseFiles(
		fmt.Sprintf("%s/%s", rootDirectory, indexFileName)))

// setup handlers
func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	err := xt.ExecuteTemplate(w, indexFileName, nil)
	if err != nil {
		c.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
