package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pb/request"

	"github.com/infraboard/mcenter/apps/book"
)

func (s *service) CreateBook(ctx context.Context, req *book.CreateBookRequest) (
	*book.Book, error) {
	ins, err := book.NewBook(req)
	if err != nil {
		return nil, exception.NewBadRequest("validate create book error, %s", err)
	}

	if err := s.save(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *service) DescribeBook(ctx context.Context, req *book.DescribeBookRequest) (
	*book.Book, error) {
	return s.get(ctx, req.Id)
}

func (s *service) QueryBook(ctx context.Context, req *book.QueryBookRequest) (
	*book.BookSet, error) {
	query := newQueryBookRequest(req)
	return s.query(ctx, query)
}

func (s *service) UpdateBook(ctx context.Context, req *book.UpdateBookRequest) (
	*book.Book, error) {
	ins, err := s.DescribeBook(ctx, book.NewDescribeBookRequest(req.Id))
	if err != nil {
		return nil, err
	}

	switch req.UpdateMode {
	case request.UpdateMode_PUT:
		ins.Update(req)
	case request.UpdateMode_PATCH:
		err := ins.Patch(req)
		if err != nil {
			return nil, err
		}
	}

	// 校验更新后数据合法性
	if err := ins.Data.Validate(); err != nil {
		return nil, err
	}

	if err := s.update(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *service) DeleteBook(ctx context.Context, req *book.DeleteBookRequest) (
	*book.Book, error) {
	ins, err := s.DescribeBook(ctx, book.NewDescribeBookRequest(req.Id))
	if err != nil {
		return nil, err
	}

	if err := s.deleteBook(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}
