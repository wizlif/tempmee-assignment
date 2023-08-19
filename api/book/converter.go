package book

import (
	db "github.com/wizlif/tempmee_assignment/db/sqlc"
	bpb "github.com/wizlif/tempmee_assignment/pb/book/v1"
)

func dbBookToPbBook(book db.Book) *bpb.Book {
	return &bpb.Book{
		Id: int32(book.ID),
		Title:     book.Title,
		Author:    book.Author,
		Price:     book.Price,
		PageCount: int32(book.PageCount),
	}
}

func dbBooksToPbBooks(books []db.Book) (pbBooks []*bpb.Book) {
	for _, b := range books {
		pbBooks = append(pbBooks, dbBookToPbBook(b))
	}

	return
}
