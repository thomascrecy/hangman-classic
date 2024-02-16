package main

import (
    "net/http"
	"html/template"
)

type PageVariables struct {
	Title string
}

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/tanguy", TanguyPage)

	http.ListenAndServe(":8080", nil)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	// Définition des données à afficher dans la page
	pageVariables := PageVariables{
		Title: "Page de base",
	}

	// Parsing du fichier HTML
	tmpl, err := template.New("index").Parse(`
		<html>
		<head>
			<title>{{.Title}}</title>
		</head>
		<body>
			<h1>{{.Title}}</h1>
			<form action="/tanguy" method="get">
				<button type="submit">Aller à Tanguy</button>
			</form>
		</body>
		</html>
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Exécution du template avec les données
	err = tmpl.Execute(w, pageVariables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func TanguyPage(w http.ResponseWriter, r *http.Request) {
	// Définition des données à afficher dans la page
	pageVariables := PageVariables{
		Title: "Page Tanguy",
	}

	// Parsing du fichier HTML
	tmpl, err := template.New("tanguy").Parse(`
		<html>
		<head>
			<title>{{.Title}}</title>
		</head>
		<body>
			<h1>{{.Title}}</h1>
			<img src="cat.png" alt="ceci est un chat" width="854" height="480">
			<form action="/" method="get">
				<button type="submit">Retour à la page de base</button>
			</form>
		</body>
		</html>
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Exécution du template avec les données
	err = tmpl.Execute(w, pageVariables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}