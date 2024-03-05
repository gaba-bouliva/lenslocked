package main

import (
	"fmt"
	"net/http"
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

// func pathHandler(w http.ResponseWriter, r *http.Request) {

// 	switch r.URL.Path {
// 		case "/":
// 			homeHandler(w,r)
// 		case "/contact":
// 			contactHandler(w,r)
// 		default:
// 			// TODO handle the page not found error
// 			pageNotFoundHandler(w,r)

// 	}

// }

type Router struct {}

func (router Router)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w,r)
	case "/contact":
		contactHandler(w,r)
	case "/faq":
		faqHandler(w,r)
	default:
		// TODO handle the page not found error
		pageNotFoundHandler(w,r)

	}
}

func main() {
	// http.HandleFunc("/", pathHandler)
	
	var router Router

	fmt.Println("Starting the server on :3000...")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err)
	}
}

