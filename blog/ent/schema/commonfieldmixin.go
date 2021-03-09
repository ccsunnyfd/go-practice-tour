package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
)

// CommonFieldMixin holds the schema definition for the CommonFieldMixin entity.
type CommonFieldMixin struct {
	mixin.Schema
}

// Fields of the CommonFieldMixin.
func (CommonFieldMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("created_by").
			Immutable(),
		field.String("modified_by").
			Optional(),
		field.Time("deleted_at").
			Immutable().Optional(),
		field.Bool("is_del").Default(false),
	}
}
