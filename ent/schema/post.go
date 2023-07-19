package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/koalitz/backend/pkg/middleware/bind"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Unique().Match(bind.TitleRegexp).
			StructTag(`json:"name,omitempty" validate:"omitempty,gte=5,lte=70,name"`),

		field.String("image").Optional().MinLen(20).
			StructTag(`json:"image,omitempty"`),

		field.Text("summary").NotEmpty().MaxLen(1024),

		field.String("place").Optional().MaxLen(100),
	}
}

func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("posts").
			Unique().StructTag(`json:"-"`),
	}
}
