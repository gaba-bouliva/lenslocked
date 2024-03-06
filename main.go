package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)


func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contect-Type", "text/html charset-UTF-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Contect-Type", "text/html charset-UTF-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:jon@calhoun.io\">jon@calhoun.io</a>.")
}

func faqHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html charset-UTF-8")
	fmt.Fprint(w, ` <h1> FAQ Page </h1>
	<p>Q: Is there a free version ? </p> 
	<p>A: Yes! We offoer a free trial for 30 days on paid plans. </p></br>
	<p>Q: What are your support hours?</p>
	<p>A: We have support staff answering emails 24/7, though response times may be a bit slower on weekends.</p></br>
	<p>Q: How do I contact support?</p>
	<p>A: Email us - support@lenslcoked.com</p>
	`)
}

func pageNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contect-Type", "text/html charset-UTF-8")
	http.Error(w, "Page not found", http.StatusNotFound)

}

func main() {
	// http.HandleFunc("/", pathHandler)
	
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func (w http.ResponseWriter, r *http.Request)  {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}

