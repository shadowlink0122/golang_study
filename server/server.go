package main

import(
	"net/http"
	"text/template"
)

type Page struct{
	Title string
	Count int
}

func viewHandler(w http.ResponseWriter, r *http.Request){
	page := Page{"Hello World", 0}
	temp, err := template.ParseFiles("mainPage.html")

	if err != nil{
		panic(err)
	}

	err = temp.Execute(w, page)

	if err != nil{
		panic(err)
	}
}

func main(){
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8000", nil)
}

