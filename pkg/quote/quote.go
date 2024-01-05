package quote

import "github.com/google/uuid"

type Quote struct {
	Id     string `json:"id" bson:"_id"`
	Author string `json:"author" bson:"author""`
	Text   string `json:"text" bson:"text""`
	BookId string `json:"book_id" bson:"book_id"`
}

type QuoteRequest struct {
	Author string `json:"author" bson:"author"`
	Text   string `json:"text" bson:"text"`
	BookId string `json:"book_id" bson:"book_id"`
}

func NewQuote(req QuoteRequest) *Quote {
	return &Quote{
		Id:     uuid.NewString(),
		Author: req.Author,
		Text:   req.Text,
		BookId: req.BookId,
	}
}
