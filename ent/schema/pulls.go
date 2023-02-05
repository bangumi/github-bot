package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Pulls holds the schema definition for the Pulls entity.
type Pulls struct {
	ent.Schema
}

// Fields of the Pulls.
func (Pulls) Fields() []ent.Field {
	return []ent.Field{
		field.String("owner"),
		field.String("repo"),
		field.Int("number").Unique().Comment("pr number"),
		field.Int64("comment").Optional().Nillable().Comment("bot comment id, nil present un-comment Pulls"),
		field.Time("createdAt"),
		field.Time("mergedAt").Optional(),
		field.Int64("checkRunID").Default(0),
		field.String("checkRunResult").Default(""),
		field.String("headSha").Default(""),
	}
}

// Edges of the Pulls.
func (Pulls) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Creator", User.Type).Ref("pull_requests").Required().Unique(),
	}
}
