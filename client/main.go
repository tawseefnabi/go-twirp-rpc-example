package main

import (
	"context"
	"log"
	"net/http"

	"github.com/tawseefnabi/go-twirp-rpc-example.git/rpc/notes"
)

func main() {
	client := notes.NewNotesServiceProtobufClient("http://localhost:9000", &http.Client{})

	ctx := context.Background()

	_, err := client.CreateNote(ctx, &notes.CreateNoteParams{Text: "Hello World2"})
	if err != nil {
		log.Fatal(err)
	}

	allNotes, err := client.GetAllNotes(ctx, &notes.GetAllNotesParams{})
	if err != nil {
		log.Fatal(err)
	}

	for _, note := range allNotes.Notes {
		log.Println(note)
	}
}
