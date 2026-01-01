package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"strconv"
)

type Note struct {
	ID      int
	Content string
}

var notes = []Note{} //slice of notes
var nextID int = 1

func getNotesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func postNotesHandler(w http.ResponseWriter, r *http.Request) {
	var theNote Note

	json.NewDecoder(r.Body).Decode(&theNote)

	theNote.ID = nextID
	nextID++

	notes = append(notes, theNote)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(theNote)
	fmt.Println(len(notes))
}

func deleteNotesHandler(w http.ResponseWriter, r *http.Request) {
	idString := strings.TrimPrefix(r.URL.Path, "/notes/")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Podane ID nie jest liczbą", http.StatusBadRequest)
		return
	}	

	targetIndex := -1
    for i, note := range notes {
        if note.ID == id {
            targetIndex = i
            break
        }
    }

	if targetIndex == -1 {
        http.Error(w, "Nie znaleziono notatki", http.StatusNotFound)
        return
    }

	notes = append(notes[:targetIndex], notes[targetIndex+1:]...)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getNotesHandler(w, r)
	case "POST":
		postNotesHandler(w, r)
	case "DELETE":
		deleteNotesHandler(w, r)
	default:
		fmt.Println("not a case")
	}
}

func main() {

	http.HandleFunc("/notes/", Handler)
	http.ListenAndServe(":8080", nil)

}
