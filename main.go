package main

import (
    "net/http"
    "./routes"
)

//Function for Serving the Content
func serveSingleResource(pattern string, filename string) {
    http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, filename)
    })
}


func main() {
	//All Dynamic Routes are Assing in main.go come from routes.go
    http.HandleFunc("/", routes.RootHandler)

    // Mandatory root-based resources
    serveSingleResource("/sitemap.xml", "./sitemap.xml")
    serveSingleResource("/favicon.ico", "./favicon.png")
    serveSingleResource("/robots.txt", "./robots.txt")

    // Normal Static Resources
    http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))


    http.ListenAndServe(":8080", nil)
}