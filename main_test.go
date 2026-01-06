package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetNotesHandler(t *testing.T){
	//1.Arrange
	notes = []Note{
		{ID: 1, Content: "Testowa notatka"},
		{ID: 2, Content: "Druga notatka"},
	}
	//fake rquest
	req, err := http.NewRequest("GET", "/notes/", nil)
	if err != nil {
		t.Fatal(err)
	}
	//fake response recorder
	rr := httptest.NewRecorder()

	//2.Act
	getNotesHandler(rr, req)

	//3.Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v",
		status, http.StatusOK)
	}

	var got []Note
	err = json.Unmarshal(rr.Body.Bytes(), &got)
	if err != nil {
		t.Fatal("Failed to parse JSON response:", err)
	}
	expected := notes
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("handler returned unexpected body: got %v, want %v", got, expected)
	}
}