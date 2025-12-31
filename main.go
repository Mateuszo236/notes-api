package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type Note struct {
	ID int
	Content string 
}

var notes []Note//slice of notes
var numberOfNotes int = 0

func createNoteHandler(w http.ResponseWriter, r *http.Request) {
	var theNote Note
	json.NewDecoder(r.Body).Decode(&theNote)
	theNote.ID = numberOfNotes
	numberOfNotes++
	notes = append(notes, theNote)
	fmt.Printf("dodano notatke %+v. Liczba notatek %d\n", theNote, len(notes))
	fmt.Fprintf(w, "Numer notatki %d", theNote.ID)
}

	

//HealthHandler
func HealthHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "OK\n")
}

func main(){

	http.HandleFunc("/health", HealthHandler)
	http.HandleFunc("/notes", createNoteHandler)
	http.ListenAndServe(":8080" , nil)

	
}
