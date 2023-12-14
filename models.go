package main

import "github.com/google/uuid"

type Quote struct {
	Id     string `json:"id" bson:"_id"`
	Author string `json:"author" bson:"author"`
	Text   string `json:"text" bson:"text"`
}

type CreateQuoteRequest struct {
	Author string `json:"author" bson:"author"`
	Text   string `json:"text" bson:"text"`
}

func NewQuote(author, text string) *Quote {
	return &Quote{
		Id:     uuid.NewString(),
		Author: author,
		Text:   text,
	}
}
