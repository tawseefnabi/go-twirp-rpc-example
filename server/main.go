package main

import "github.com/tawseefnabi/go-twirp-rpc-example.git/rpc/notes"

type noteService struct {
	Notes     []notes.Note
	CurrentId int32
}

func main() {

}
