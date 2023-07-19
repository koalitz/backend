// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/koalitz/backend/ent/post"
	"github.com/koalitz/backend/ent/schema"
	"github.com/koalitz/backend/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescTitle is the schema descriptor for title field.
	postDescTitle := postFields[0].Descriptor()
	// post.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	post.TitleValidator = postDescTitle.Validators[0].(func(string) error)
	// postDescImage is the schema descriptor for image field.
	postDescImage := postFields[1].Descriptor()
	// post.ImageValidator is a validator for the "image" field. It is called by the builders before save.
	post.ImageValidator = postDescImage.Validators[0].(func(string) error)
	// postDescSummary is the schema descriptor for summary field.
	postDescSummary := postFields[2].Descriptor()
	// post.SummaryValidator is a validator for the "summary" field. It is called by the builders before save.
	post.SummaryValidator = func() func(string) error {
		validators := postDescSummary.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(summary string) error {
			for _, fn := range fns {
				if err := fn(summary); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// postDescPlace is the schema descriptor for place field.
	postDescPlace := postFields[3].Descriptor()
	// post.PlaceValidator is a validator for the "place" field. It is called by the builders before save.
	post.PlaceValidator = postDescPlace.Validators[0].(func(string) error)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[0].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = func() func(string) error {
		validators := userDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(email string) error {
			for _, fn := range fns {
				if err := fn(email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescRole is the schema descriptor for role field.
	userDescRole := userFields[1].Descriptor()
	// user.DefaultRole holds the default value on creation for the role field.
	user.DefaultRole = userDescRole.Default.(string)
	// userDescFirstName is the schema descriptor for first_name field.
	userDescFirstName := userFields[2].Descriptor()
	// user.FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	user.FirstNameValidator = func() func(string) error {
		validators := userDescFirstName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(first_name string) error {
			for _, fn := range fns {
				if err := fn(first_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescLastName is the schema descriptor for last_name field.
	userDescLastName := userFields[3].Descriptor()
	// user.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	user.LastNameValidator = func() func(string) error {
		validators := userDescLastName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(last_name string) error {
			for _, fn := range fns {
				if err := fn(last_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
}