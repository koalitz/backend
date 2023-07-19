package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/koalitz/backend/pkg/middleware/bind"
)

// User holds the schema definition for the User dto.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{

		field.String("email").Unique().NotEmpty().Match(bind.EmailRegexp).
			StructTag(`json:"email,omitempty" validate:"required,email"`),

		field.String("role").Default("member").StructTag(`json:"role,omitempty" validate:"omitempty,enum=member*organizer*admin"`),

		field.String("first_name").MinLen(3).MaxLen(32).Nillable().
			StructTag(`json:"firstName,omitempty" validate:"omitempty,gte=3,lte=32"`),

		field.String("last_name").MinLen(3).MaxLen(32).Nillable().
			StructTag(`json:"lastName,omitempty" validate:"omitempty,gte=3,lte=32"`),

		field.Strings("sessions").Optional().StructTag(`json:"-"`),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type).StructTag(`json:"-"`),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
