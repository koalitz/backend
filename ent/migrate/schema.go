// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString, Unique: true},
		{Name: "image", Type: field.TypeString, Nullable: true},
		{Name: "summary", Type: field.TypeString, Size: 1024},
		{Name: "place", Type: field.TypeString, Nullable: true, Size: 100},
		{Name: "user_posts", Type: field.TypeInt, Nullable: true},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "posts_users_posts",
				Columns:    []*schema.Column{PostsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "role", Type: field.TypeString, Default: "member"},
		{Name: "first_name", Type: field.TypeString, Size: 32},
		{Name: "last_name", Type: field.TypeString, Size: 32},
		{Name: "sessions", Type: field.TypeJSON, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PostsTable,
		UsersTable,
	}
)

func init() {
	PostsTable.ForeignKeys[0].RefTable = UsersTable
}
