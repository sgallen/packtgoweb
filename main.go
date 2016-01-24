package main

import (
    "fmt"
    "net/http"
    "html/template"
)

func main() {
    templates := template.Must(template.ParseGlob("templates/*"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if err := templates.ExecuteTemplate(w, "index", nil); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    })

    fmt.Println(http.ListenAndServe(":8080", nil))
}
