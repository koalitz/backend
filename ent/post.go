// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/koalitz/backend/ent/post"
	"github.com/koalitz/backend/ent/user"
)

// Post is the model entity for the Post schema.
type Post struct {
	config `json:"-" validate:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"name,omitempty" validate:"omitempty,gte=5,lte=70,name"`
	// Image holds the value of the "image" field.
	Image string `json:"image,omitempty"`
	// Summary holds the value of the "summary" field.
	Summary string `json:"summary,omitempty"`
	// Place holds the value of the "place" field.
	Place string `json:"place,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PostQuery when eager-loading is set.
	Edges      PostEdges `json:"edges"`
	user_posts *int
}

// PostEdges holds the relations/edges for other nodes in the graph.
type PostEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"-"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PostEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Post) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case post.FieldID:
			values[i] = new(sql.NullInt64)
		case post.FieldTitle, post.FieldImage, post.FieldSummary, post.FieldPlace:
			values[i] = new(sql.NullString)
		case post.ForeignKeys[0]: // user_posts
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Post", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Post fields.
func (po *Post) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case post.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			po.ID = int(value.Int64)
		case post.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				po.Title = value.String
			}
		case post.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				po.Image = value.String
			}
		case post.FieldSummary:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field summary", values[i])
			} else if value.Valid {
				po.Summary = value.String
			}
		case post.FieldPlace:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field place", values[i])
			} else if value.Valid {
				po.Place = value.String
			}
		case post.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_posts", value)
			} else if value.Valid {
				po.user_posts = new(int)
				*po.user_posts = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Post entity.
func (po *Post) QueryOwner() *UserQuery {
	return NewPostClient(po.config).QueryOwner(po)
}

// Update returns a builder for updating this Post.
// Note that you need to call Post.Unwrap() before calling this method if this Post
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Post) Update() *PostUpdateOne {
	return NewPostClient(po.config).UpdateOne(po)
}

// Unwrap unwraps the Post entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Post) Unwrap() *Post {
	_tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: Post is not a transactional entity")
	}
	po.config.driver = _tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Post) String() string {
	var builder strings.Builder
	builder.WriteString("Post(")
	builder.WriteString(fmt.Sprintf("id=%v, ", po.ID))
	builder.WriteString("title=")
	builder.WriteString(po.Title)
	builder.WriteString(", ")
	builder.WriteString("image=")
	builder.WriteString(po.Image)
	builder.WriteString(", ")
	builder.WriteString("summary=")
	builder.WriteString(po.Summary)
	builder.WriteString(", ")
	builder.WriteString("place=")
	builder.WriteString(po.Place)
	builder.WriteByte(')')
	return builder.String()
}

// Posts is a parsable slice of Post.
type Posts []*Post
