package book

import (
	"context"

	db "github.com/wizlif/tempmee_assignment/db/sqlc"
	bpb "github.com/wizlif/tempmee_assignment/pb/book/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *BookServer) GetBooks(ctx context.Context, req *bpb.GetBooksRequest) (*bpb.GetBooksResponse, error) {
	page := func() int64 {
		if req.Page == nil || req.GetPage() <=0 {
			return 1
		}
		return req.GetPage()
	}()

	limit := func() int64 {
		if req.Limit == nil || req.GetLimit() <= 0 {
			return 10
		}
		return req.GetLimit()
	}()

	books, err := server.db.ListBooks(ctx, db.ListBooksParams{
		Offset: (page - 1) * limit,
		Limit: limit,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get books")
	}

	return &bpb.GetBooksResponse{
		Books: dbBooksToPbBooks(books),
	}, nil
}
