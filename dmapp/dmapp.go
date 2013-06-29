package dmapp

import (
	"appengine"
	"html/template"
	"net/http"
	"strconv"
)

func init() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/viewMonster/", viewMonsterHandler)
	http.HandleFunc("/newMonster/", newMonsterHandler)
	http.HandleFunc("/deleteMonster/", deleteMonsterHandler)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	t, err := template.ParseFiles("templates/base.gohtml", "templates/default.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	monsters, err := getAllMonsters(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, monsters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func viewMonsterHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	path := r.URL.Path
	encodedKey := path[len("/viewMonster/"):]
	monster, err := getMonster(c, encodedKey)
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

func newMonsterHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	t, err := template.ParseFiles("templates/base.gohtml", "templates/newMonster.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if r.Method == "GET" {
		err = t.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		var monster = new(Monster)
		var tmpInt int64

		monster.Name = r.FormValue("Name")
		tmpInt, err = strconv.ParseInt(r.FormValue("Level"), 10, 64)
		monster.Level = int(tmpInt)
		monster.Role = r.FormValue("Role")
		monster.Size = r.FormValue("Size")
		monster.Origin = r.FormValue("Origin")
		monster.Type = r.FormValue("Type")
		tmpInt, err = strconv.ParseInt(r.FormValue("XP"), 10, 64)
		monster.XP = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("Health"), 10, 64)
		monster.Health = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("Initiative"), 10, 64)
		monster.InitiativeBonus = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("ArmorClass"), 10, 64)
		monster.ArmorClass = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("Fortitude"), 10, 64)
		monster.Fortitude = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("Reflex"), 10, 64)
		monster.Reflex = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("Will"), 10, 64)
		monster.Will = int(tmpInt)
		monster.Speed = r.FormValue("Speed")
		tmpInt, err = strconv.ParseInt(r.FormValue("SavingThrows"), 10, 64)
		monster.SavingThrows = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("ActionPoints"), 10, 64)
		monster.ActionPoints = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("Acrobatics"), 10, 64)
		monster.Acrobatics = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("Arcana"), 10, 64)
		monster.Arcana = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("Athletics"), 10, 64)
		monster.Athletics = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("Bluff"), 10, 64)
		monster.Bluff = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("Diplomacy"), 10, 64)
		monster.Diplomacy = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("Dungeoneering"), 10, 64)
		monster.Dungeoneering = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("Endurance"), 10, 64)
		monster.Endurance = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("Heal"), 10, 64)
		monster.Heal = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("History"), 10, 64)
		monster.History = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("Insight"), 10, 64)
		monster.Insight = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("Intimidate"), 10, 64)
		monster.Intimidate = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("Nature"), 10, 64)
		monster.Nature = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("Perception"), 10, 64)
		monster.Perception = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("Religion"), 10, 64)
		monster.Religion = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("Stealth"), 10, 64)
		monster.Stealth = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("Streetwise"), 10, 64)
		monster.Streetwise = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("Thievery"), 10, 64)
		monster.Thievery = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("Strength"), 10, 64)
		monster.Strength = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("Constitution"), 10, 64)
		monster.Constitution = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("Dexterity"), 10, 64)
		monster.Dexterity = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("Intelligence"), 10, 64)
		monster.Intelligence = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("Wisdom"), 10, 64)
		monster.Wisdom = int(tmpInt)
		tmpInt, err = strconv.ParseInt(r.FormValue("Charisma"), 10, 64)
		monster.Charisma = int(tmpInt)
		monster.Alignment = r.FormValue("Alignment")

		encodedKey, err := saveMonster(c, monster)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/viewMonster/"+encodedKey, http.StatusSeeOther)
	}
}

func deleteMonsterHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	if r.Method == "GET" {
		path := r.URL.Path
		encodedKey := path[len("/deleteMonster/"):]
		monster, err := getMonster(c, encodedKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t, err := template.ParseFiles("templates/base.gohtml", "templates/deleteMonster.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.Execute(w, monster)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		encodedKey := r.FormValue("Key")
		err := deleteMonster(c, encodedKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
