package postgres

import (
	"context"
	"github.com/koalitz/backend/ent"
	"github.com/koalitz/backend/ent/post"
	"github.com/koalitz/backend/internal/controller/dto"
)

type PostStorage struct {
	postClient *ent.PostClient
}

func NewPostStorage(postClient *ent.PostClient) *PostStorage {
	return &PostStorage{postClient: postClient}
}

func (p *PostStorage) FindPostByID(ctx context.Context, id int) (*ent.Post, error) {
	return p.postClient.Get(ctx, id)
}

func (p *PostStorage) FindPostByTitle(ctx context.Context, title string) ([]*ent.Post, error) {
	return p.postClient.Query().Where(post.TitleContains(title)).All(ctx)
}

func (p *PostStorage) CreatePost(ctx context.Context, imageName *string, info dto.PostInfo, id int) (*ent.Post, error) {
	return p.postClient.Create().SetTitle(info.Title).
		SetNillableImage(imageName).SetOwnerID(id).
		SetPlace(info.Place).SetSummary(info.Summary).Save(ctx)
}
