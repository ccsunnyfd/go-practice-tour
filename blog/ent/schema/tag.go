package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Bool("is_available").Default(true),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("articles", Article.Type).Ref("tags"),
	}
}

// Mixin of the Tag
func (Tag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
		CommonFieldMixin{},
	}
}
