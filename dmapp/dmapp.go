package dmapp

import (
	"html/template"
	"net/http"
	"strconv"
)

func init() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/viewMonster/", viewMonsterHandler)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/base.gohtml", "templates/default.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func viewMonsterHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	key := path[len("/viewMonster/"):]
	monster, err := getMonster(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t := template.New("base.gohtml")
	t.Funcs(template.FuncMap{
		"modInt":    modInt,
		"commaList": CommaList,
	})
	t, err = t.ParseFiles("templates/base.gohtml", "templates/monsterInfo.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, monster)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getMonster(key string) (Monster, error) {
	i, err := strconv.ParseInt(key[len("Monster"):], 10, 64)
	return m[i], err
}
