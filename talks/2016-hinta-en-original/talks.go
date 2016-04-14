package webapp

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"code.google.com/p/go.tools/present"
)

var basePath = "."
var presPath = "present"

// Maps the presentable extensions to the templates to be executed.
var extensions = map[string]string{
	".slide":   "slides.tmpl",
	".article": "article.tmpl",
}

// Sets the talk handler for slides and static files.
func init() {
	http.HandleFunc("/", talkHandler)
}

// Serves either a talk or an associated resource file.
func talkHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(basePath, r.URL.Path)

	// Check the requested resource's type
	if file, err := os.Open(path); err == nil {
		defer file.Close()
		if stat, err := file.Stat(); err == nil {
			// Accessing directory listings is forbidden (just no point really)
			if stat.IsDir() {
				http.Error(w, "Forbidden", 403)
				return
			}
			// If the request is for a slide, generate it
			if _, ok := extensions[filepath.Ext(path)]; ok {
				err := renderDoc(w, presPath, path, file)
				if err != nil {
					http.Error(w, err.Error(), 500)
				}
				return
			}
		}
	} else {
		log.Println(err)
	}
	// Nothing interesting, serve by plain file server
	http.FileServer(http.Dir(basePath)).ServeHTTP(w, r)
}

// Reads the present file, builds its template representation, and executes the
// template, sending output to w.
func renderDoc(w io.Writer, base, name string, pres *os.File) error {
	// Read the input and build the doc structure
	doc, err := present.Parse(pres, name, 0)
	if err != nil {
		return err
	}
	// Find which template should be executed
	contentTmpl, ok := extensions[filepath.Ext(name)]
	if !ok {
		return fmt.Errorf("no template for extension %v", filepath.Ext(name))
	}
	// Locate the template file
	actionTmpl := filepath.Join(presPath, "templates/action.tmpl")
	contentTmpl = filepath.Join(presPath, "templates", contentTmpl)

	// Read and parse the input
	tmpl := present.Template()
	if _, err := tmpl.ParseFiles(actionTmpl, contentTmpl); err != nil {
		return err
	}
	// Execute the template.
	return doc.Render(w, tmpl)
}
