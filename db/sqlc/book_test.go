package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wizlif/tempmee_assignment/util"
)

func TestListBooks(t *testing.T) {
	limit := 5
	var createdBooks []Book
	for i := 0; i < limit*2; i++ {
		createdBooks = append(createdBooks, createRandomBook(t))
	}

	arg := ListBooksParams{
		Limit:  int64(limit),
		Offset: 0,
	}

	books, err := testStore.ListBooks(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, books)
	require.Equal(t, limit, len(books))

	for index, book := range books {
		require.NotEmpty(t, book)
		require.Equal(t, createdBooks[index], book)
	}
}

func createRandomBook(t *testing.T) Book {
	arg := CreateBookParams{
		Title:     util.RandomTitle(),
		Author:    util.RandomTitle(),
		Price:     util.RandomFloat(30.0, 80.0),
		PageCount: util.RandomInt(30, 80),
	}

	bookId, err := testStore.CreateBook(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, bookId)

	book, err := testStore.GetBookById(context.Background(), bookId)
	require.NoError(t, err)
	require.NotEmpty(t, bookId)

	require.Equal(t, arg.Title, book.Title)
	require.Equal(t, arg.Author, book.Author)
	require.Equal(t, arg.Price, book.Price)
	require.Equal(t, arg.PageCount, book.PageCount)
	require.NotZero(t, book.CreatedAt)

	return book
}
