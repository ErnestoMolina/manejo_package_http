package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	rps "rpsweb/RPS"
	"strconv"
)

const (
	templateDir  = "templates/"
	templateBase = templateDir + "base.html"
)

type Player struct {
	Name string
}

var player Player

func Index(w http.ResponseWriter, r *http.Request) {
	restarValue()
	RenderTemplate(w, "index.html", nil)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	restarValue()
	RenderTemplate(w, "newGame.html", nil)
}

func Game(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}
		player.Name = r.Form.Get("name")
	}
	// Redireccionamiento a otra pagina
	if player.Name == "" {
		http.Redirect(w, r, "/New", http.StatusFound)
	}
	// fmt.Println(player.Name)
	RenderTemplate(w, "game.html", player)
}

func Play(w http.ResponseWriter, r *http.Request) {
	playerChoise, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoise)
	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func About(w http.ResponseWriter, r *http.Request) {
	restarValue()
	RenderTemplate(w, "about.html", nil)
}

func RenderTemplate(w http.ResponseWriter, page string, data any) {
	tpl := template.Must(template.ParseFiles(templateBase, templateDir+page))

	err := tpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Error al renderizar plantilla", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

// Reiniciar valores
func restarValue() {
	player.Name = ""
	rps.ComputerScore = 0
	rps.PlayerScore = 0
}
