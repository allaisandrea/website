package main

import "flag"
import "os"
import "fmt"
import "path/filepath"
// import "log"
import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
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
        fmt.Printf("usage: serve path/to/pages\n")
        os.Exit(1)
    }
    pagesPath:= flag.Arg(0)
    pages, err:= getAllPages(pagesPath)
    if err != nil {
        fmt.Printf("Cannot list path %q: %q\n", pagesPath, err)
        os.Exit(1)
    }
    fmt.Printf("Pages: %v\n", pages)
//     http.HandleFunc("/", handler)
//     log.Fatal(http.ListenAndServe(":8080", nil))
}
