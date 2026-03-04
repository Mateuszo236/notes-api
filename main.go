package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"strconv"
	"log/slog"
)

type Note struct {
	ID      int
	Content string
}

var notes = []Note{} //slice of notes
var nextID int = 1

func getNotesHandler(w http.ResponseWriter, r *http.Request) {//Read
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
	slog.Info("Notes retrieved", "count", len(notes))
}

func postNotesHandler(w http.ResponseWriter, r *http.Request) {//Creat
	var theNote Note

	if err := json.NewDecoder(r.Body).Decode(&theNote);err != nil {
		slog.Error("Failed to decode JSON", "error", err)
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(theNote.Content) == "" {
		slog.Warn("Content is empty", "note", theNote)
        http.Error(w, "Content cannot be empty", http.StatusBadRequest)
        return
    }

	theNote.ID = nextID
	nextID++

	notes = append(notes, theNote)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(theNote)
	slog.Info("Note created", "id", theNote.ID, "content_length", len(theNote.Content))
}

func updateNotesHandler(w http.ResponseWriter, r *http.Request) {

	 
	

	idString := strings.TrimPrefix(r.URL.Path, "/notes/")
	id, err := strconv.Atoi(idString)
	if err != nil {
		slog.Warn("Invalid ID format", "id_string", idString, "error", err)
		http.Error(w, "Gived ID is not a numer", http.StatusBadRequest)
		return
	}

	var theNote Note

	if err := json.NewDecoder(r.Body).Decode(&theNote);err != nil {
		slog.Error("Failed to decode JSON", "error", err)
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(theNote.Content) == "" {
		slog.Warn("Content is empty", "note", theNote)
        http.Error(w, "Content cannot be empty", http.StatusBadRequest)
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
		slog.Warn("ID not found for update", "id", id)
		http.Error(w, "note with given ID not found", http.StatusNotFound)
		return
	}
	
	notes[targetIndex].Content = theNote.Content
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(theNote)

	slog.Info("Note updated successfully", "id", id, "new_content_length", len(theNote.Content))


}

func deleteNotesHandler(w http.ResponseWriter, r *http.Request) {//Delete
	idString := strings.TrimPrefix(r.URL.Path, "/notes/")
	id, err := strconv.Atoi(idString)
	if err != nil {
		slog.Warn("Invalid ID format", "id_string", idString, "error", err)
		http.Error(w, "Gived ID is not a number", http.StatusBadRequest)
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
		slog.Warn("ID not found for deletion", "id", id)
        http.Error(w, "note with given ID not found", http.StatusNotFound)
        return
    }

	notes = append(notes[:targetIndex], notes[targetIndex+1:]...)

	slog.Info("Note deleted successfully", "id", id)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getNotesHandler(w, r)
	case "POST":
		postNotesHandler(w, r)
	case "DELETE":
		deleteNotesHandler(w, r)
	case "PUT":
		updateNotesHandler(w, r)
	default:
		fmt.Println("not a case")
	}
}

func main() {

	http.HandleFunc("/notes/", Handler)
	http.ListenAndServe(":8080", nil)

}
