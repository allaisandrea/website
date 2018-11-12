package main

import "flag"
import "os"
import "fmt"
import "path"
import "path/filepath"
import "log"
import "net/http"
import "io/ioutil"
import "html/template"

type ServerData struct {
    ClientPath string
    PageTemplate *template.Template
}

type Page struct {
    Title string
    Body []byte
}


func loadPage(filename string) (*Page, error) {
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: "TODO", Body: body}, nil
}


func handleIndex(w http.ResponseWriter, r *http.Request, serverData ServerData) {
    page, err := loadPage(path.Join(serverData.ClientPath, "pages", "index.html"))

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    serverData.PageTemplate.Execute(w, page)

}


func getAllPages(path string) ([]string, error) {
    var pages []string
    err:= filepath.Walk(
        path,
        func(path string, info os.FileInfo, err error) error {
            if err != nil {
                return err
            }
            if !info.IsDir() {
                pages = append(pages, path)
            }
            return nil
        })
    return pages, err
}


func main() {
    flag.Parse()
    if flag.NArg() != 1 {
        fmt.Printf("usage: serve path/to/client\n")
        os.Exit(1)
    }

    clientPath := flag.Arg(0)

    pageTemplate, err := template.ParseFiles(path.Join(clientPath, "templates", "page.html"))
    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }

    serverData := ServerData {
        ClientPath: clientPath,
        PageTemplate: pageTemplate}

    http.HandleFunc(
        "/",
        func(w http.ResponseWriter, r *http.Request) {
            handleIndex(w, r, serverData)
        })

    log.Fatal(http.ListenAndServe(":8080", nil))
}
