package main

import (
	"log"
	"net/http"
)

// Add a handler for the root path that redirects to /home
func rootHandler(w http.ResponseWriter, r *http.Request) {
    // Only redirect if the path is exactly "/", otherwise it's a 404
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}
	http.NotFound(w, r)
}

// ... (your other page handlers remain the same) ...

func main() {
    // Handle the root path
	http.HandleFunc("/", rootHandler)
    
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/courses", coursePage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/contact", contactPage)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
