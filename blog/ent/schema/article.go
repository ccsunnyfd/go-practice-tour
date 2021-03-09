package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("desc").Default(""),
		field.String("content").Default(""),
		field.Bool("is_available").Default(true),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", Tag.Type),
	}
}

// Mixin of the Article
func (Article) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
		CommonFieldMixin{},
	}
}
