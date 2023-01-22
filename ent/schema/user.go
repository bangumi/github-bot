package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("github_id").Positive().Unique(),
		field.Int64("bangumi_id").Positive().Unique().Optional(),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		// 唯一约束索引
		index.Fields("github_id").Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pull_requests", Pulls.Type),
	}
}
