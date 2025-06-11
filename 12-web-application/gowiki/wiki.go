package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

var validPath = regexp.MustCompile("^/(edit|save|view|editTemplate|viewTemplate)/([a-zA-Z0-9]+)$")
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

type Page struct {
	Title string
	Body  []byte
}

func (page *Page) save() error {
	filename := page.Title + ".txt"
	return os.WriteFile(filename, page.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func getTitle(response http.ResponseWriter, request *http.Request) (string, error) {
	validItems := validPath.FindStringSubmatch(request.URL.Path)
	if validItems == nil {
		http.NotFound(response, request)
		return "", errors.New("invalid URL")
	}
	return validItems[2], nil
}

/*
func handler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello! I like %s", request.URL.Path[1:])
	fmt.Fprintf(response, "\n")
	fmt.Fprintf(response, "Hello! I like %s", request.URL.Path)
}
*/

func viewHandler(response http.ResponseWriter, request *http.Request) {
	//title := request.URL.Path[len("/view/"):]
	title, err := getTitle(response, request)
	if err != nil {
		return
	}
	page, err := loadPage(title)
	if err != nil {
		http.Redirect(response, request, "/edit/"+title, http.StatusFound)
		return
	}
	fmt.Fprintf(response, "<h1>%s</h1><div>%s</div>", page.Title, page.Body)
}

func viewTemplateHandler(response http.ResponseWriter, request *http.Request) {
	//title := request.URL.Path[len("/viewTemplate/"):]
	title, err := getTitle(response, request)
	if err != nil {
		return
	}
	page, err := loadPage(title)
	if err != nil {
		http.Redirect(response, request, "/editTemplate/"+title, http.StatusFound)
		return
	}
	renderTemplate(response, "view", page)
}

func editHandler(response http.ResponseWriter, request *http.Request) {
	//title := request.URL.Path[len("/edit/"):]
	title, err := getTitle(response, request)
	if err != nil {
		return
	}
	page, err := loadPage(title)
	if err != nil {
		page = &Page{Title: title}
	}
	fmt.Fprintf(response, "<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</textarea><br>"+
		"<input type=\"submit\" value=\"Guardar\">"+
		"</form>",
		page.Title, page.Title, page.Body)
}

func editTemplateHandler(response http.ResponseWriter, request *http.Request) {
	//title := request.URL.Path[len("/editTemplate/"):]
	title, err := getTitle(response, request)
	if err != nil {
		return
	}
	page, err := loadPage(title)
	if err != nil {
		page = &Page{Title: title}
	}
	renderTemplateCache(response, "edit", page)
}

func renderTemplate(response http.ResponseWriter, templateName string, page *Page) {
	template, err := template.ParseFiles(templateName + ".html")
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	err = template.Execute(response, page)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

func renderTemplateCache(response http.ResponseWriter, templateName string, page *Page) {
	err := templates.ExecuteTemplate(response, templateName+".html", page)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

func saveHandler(response http.ResponseWriter, request *http.Request) {
	//title := request.URL.Path[len("/save/"):]
	title, err := getTitle(response, request)
	if err != nil {
		return
	}
	body := request.FormValue("body")
	page := &Page{Title: title, Body: []byte(body)}
	err = page.save()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(response, request, "/viewTemplate/"+page.Title, http.StatusFound)
}

func main() {
	fmt.Println("*** Starting wiki ***")

	//http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/viewTemplate/", viewTemplateHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/editTemplate/", editTemplateHandler)
	http.HandleFunc("/save/", saveHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

	fmt.Println("*** Ending wiki ***")
}
