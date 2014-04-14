package routes

import(
	"net/http"
	"html/template"
)

//Global Var for Templates
var templates = template.Must(template.ParseFiles("views/index.html"))


func RootHandler(w http.ResponseWriter, r *http.Request) {
	//Make All Other Routes Invalid Not "/"
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	//Renders Main Template
	err := templates.ExecuteTemplate(w, "index.html" ,nil)
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

}
