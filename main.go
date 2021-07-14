package main

import (
	"net/http"

	"src/github.com/gorilla/mux/views"

	"github.com/gorilla/mux"
)

var (
	homeView    *views.View
	contactView *views.View
	faqView     *views.View
	meView      *views.View
	signUpView  *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(faqView.Render(w, nil))

}

func me(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(meView.Render(w, nil))
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(signUpView.Render(w, nil))
}

//A helper function that panics on any error
func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.html")
	contactView = views.NewView("bootstrap", "views/contact.html")
	faqView = views.NewView("bootstrap", "views/faq.html")
	meView = views.NewView("base", "views/me.html")
	signUpView = views.NewView("bootstrap", "views/signup.html")
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.HandleFunc("/me", me)
	r.HandleFunc("/signup", signup)
	http.ListenAndServe(":3000", r)

}
