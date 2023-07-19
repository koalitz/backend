package service

import (
	"context"
	"github.com/koalitz/backend/ent"
	"github.com/koalitz/backend/internal/controller/dto"
)

type PostPostgres interface {
	FindPostByID(ctx context.Context, id int) (*ent.Post, error)
	FindPostByTitle(ctx context.Context, title string) ([]*ent.Post, error)
	CreatePost(ctx context.Context, imageName *string, info dto.PostInfo, id int) (*ent.Post, error)
}

type PostService struct {
	postgres PostPostgres
}

func NewPostService(postgres PostPostgres) *PostService {
	return &PostService{postgres: postgres}
}

func (p *PostService) FindPostByID(id int) (*ent.Post, error) {
	return p.postgres.FindPostByID(context.Background(), id)
}

func (p *PostService) FindPostByTitle(title string) ([]*ent.Post, error) {
	return p.postgres.FindPostByTitle(context.Background(), title)
}

func (p *PostService) CreatePost(imageName *string, info dto.PostInfo, id int) (*ent.Post, error) {
	return p.postgres.CreatePost(context.Background(), imageName, info, id)
}
