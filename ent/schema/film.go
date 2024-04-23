package schema

import "C"
import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type Film struct {
	ent.Schema
}

// Fields of the Club.
func (Film) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name").Optional(),
		field.String("description").Optional(),
		field.Int64("author_id").Optional(),
		field.Int64("manager_id").Optional(),
		field.Int64("stadium_id").Nillable().Optional(),
		field.Float("star").Nillable().Optional(),
		field.Float("attack_rating").Nillable().Optional(),
		field.Float("midfield_rating").Nillable().Optional(),
		field.Float("defence_rating").Nillable().Optional(),
		field.Float("goalkeeper_rating").Nillable().Optional(),
		field.Float("overall_rating").Nillable().Optional(),
		field.String("logo_url").Optional(),
		field.Time("created_at").Default(time.Now()),
	}
}

// Edges of the Club.
func (Film) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("author", User.Type).Field("author_id").Unique(),
		edge.To("manager", User.Type).Field("manager_id").Unique(),
	}
}
