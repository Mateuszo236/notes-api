Notes API in Go

A small CRUD HTTP API written in Go. The project demonstrates request routing, JSON encoding and decoding, input validation, structured logging and basic HTTP handler testing.

Features

Creating notes
Retrieving all notes
Updating an existing note
Deleting a note
JSON request and response handling
Input validation
Structured logging with log/slog

API endpoints

GET /notes/ Retrieve all notes
POST /notes/ Create a new note
PUT /notes/{id} Update an existing note
DELETE /notes/{id} Delete a note

Run locally

go run .

The server listens on port 8080.

Example request

curl -X POST http://localhost:8080/notes/ -H "Content-Type: application/json" -d '{"content":"Test note"}'

Tests

go test ./...

Current scope

The project uses in-memory storage, so all notes are removed when the application is restarted. It was created as an educational project focused on Go HTTP programming and API design.
