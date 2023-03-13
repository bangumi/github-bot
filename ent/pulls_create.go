// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"github-bot/ent/pulls"
	"github-bot/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PullsCreate is the builder for creating a Pulls entity.
type PullsCreate struct {
	config
	mutation *PullsMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetOwner sets the "owner" field.
func (pc *PullsCreate) SetOwner(s string) *PullsCreate {
	pc.mutation.SetOwner(s)
	return pc
}

// SetPrID sets the "prID" field.
func (pc *PullsCreate) SetPrID(i int64) *PullsCreate {
	pc.mutation.SetPrID(i)
	return pc
}

// SetNillablePrID sets the "prID" field if the given value is not nil.
func (pc *PullsCreate) SetNillablePrID(i *int64) *PullsCreate {
	if i != nil {
		pc.SetPrID(*i)
	}
	return pc
}

// SetRepo sets the "repo" field.
func (pc *PullsCreate) SetRepo(s string) *PullsCreate {
	pc.mutation.SetRepo(s)
	return pc
}

// SetRepoID sets the "repoID" field.
func (pc *PullsCreate) SetRepoID(i int64) *PullsCreate {
	pc.mutation.SetRepoID(i)
	return pc
}

// SetNillableRepoID sets the "repoID" field if the given value is not nil.
func (pc *PullsCreate) SetNillableRepoID(i *int64) *PullsCreate {
	if i != nil {
		pc.SetRepoID(*i)
	}
	return pc
}

// SetNumber sets the "number" field.
func (pc *PullsCreate) SetNumber(i int) *PullsCreate {
	pc.mutation.SetNumber(i)
	return pc
}

// SetComment sets the "comment" field.
func (pc *PullsCreate) SetComment(i int64) *PullsCreate {
	pc.mutation.SetComment(i)
	return pc
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (pc *PullsCreate) SetNillableComment(i *int64) *PullsCreate {
	if i != nil {
		pc.SetComment(*i)
	}
	return pc
}

// SetCreatedAt sets the "createdAt" field.
func (pc *PullsCreate) SetCreatedAt(t time.Time) *PullsCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetMergedAt sets the "mergedAt" field.
func (pc *PullsCreate) SetMergedAt(t time.Time) *PullsCreate {
	pc.mutation.SetMergedAt(t)
	return pc
}

// SetNillableMergedAt sets the "mergedAt" field if the given value is not nil.
func (pc *PullsCreate) SetNillableMergedAt(t *time.Time) *PullsCreate {
	if t != nil {
		pc.SetMergedAt(*t)
	}
	return pc
}

// SetCheckRunID sets the "checkRunID" field.
func (pc *PullsCreate) SetCheckRunID(i int64) *PullsCreate {
	pc.mutation.SetCheckRunID(i)
	return pc
}

// SetNillableCheckRunID sets the "checkRunID" field if the given value is not nil.
func (pc *PullsCreate) SetNillableCheckRunID(i *int64) *PullsCreate {
	if i != nil {
		pc.SetCheckRunID(*i)
	}
	return pc
}

// SetCheckRunResult sets the "checkRunResult" field.
func (pc *PullsCreate) SetCheckRunResult(s string) *PullsCreate {
	pc.mutation.SetCheckRunResult(s)
	return pc
}

// SetNillableCheckRunResult sets the "checkRunResult" field if the given value is not nil.
func (pc *PullsCreate) SetNillableCheckRunResult(s *string) *PullsCreate {
	if s != nil {
		pc.SetCheckRunResult(*s)
	}
	return pc
}

// SetHeadSha sets the "headSha" field.
func (pc *PullsCreate) SetHeadSha(s string) *PullsCreate {
	pc.mutation.SetHeadSha(s)
	return pc
}

// SetNillableHeadSha sets the "headSha" field if the given value is not nil.
func (pc *PullsCreate) SetNillableHeadSha(s *string) *PullsCreate {
	if s != nil {
		pc.SetHeadSha(*s)
	}
	return pc
}

// SetCreatorID sets the "Creator" edge to the User entity by ID.
func (pc *PullsCreate) SetCreatorID(id int) *PullsCreate {
	pc.mutation.SetCreatorID(id)
	return pc
}

// SetCreator sets the "Creator" edge to the User entity.
func (pc *PullsCreate) SetCreator(u *User) *PullsCreate {
	return pc.SetCreatorID(u.ID)
}

// Mutation returns the PullsMutation object of the builder.
func (pc *PullsCreate) Mutation() *PullsMutation {
	return pc.mutation
}

// Save creates the Pulls in the database.
func (pc *PullsCreate) Save(ctx context.Context) (*Pulls, error) {
	pc.defaults()
	return withHooks[*Pulls, PullsMutation](ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PullsCreate) SaveX(ctx context.Context) *Pulls {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PullsCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PullsCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PullsCreate) defaults() {
	if _, ok := pc.mutation.PrID(); !ok {
		v := pulls.DefaultPrID
		pc.mutation.SetPrID(v)
	}
	if _, ok := pc.mutation.RepoID(); !ok {
		v := pulls.DefaultRepoID
		pc.mutation.SetRepoID(v)
	}
	if _, ok := pc.mutation.Comment(); !ok {
		v := pulls.DefaultComment
		pc.mutation.SetComment(v)
	}
	if _, ok := pc.mutation.CheckRunID(); !ok {
		v := pulls.DefaultCheckRunID
		pc.mutation.SetCheckRunID(v)
	}
	if _, ok := pc.mutation.CheckRunResult(); !ok {
		v := pulls.DefaultCheckRunResult
		pc.mutation.SetCheckRunResult(v)
	}
	if _, ok := pc.mutation.HeadSha(); !ok {
		v := pulls.DefaultHeadSha
		pc.mutation.SetHeadSha(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PullsCreate) check() error {
	if _, ok := pc.mutation.Owner(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required field "Pulls.owner"`)}
	}
	if _, ok := pc.mutation.PrID(); !ok {
		return &ValidationError{Name: "prID", err: errors.New(`ent: missing required field "Pulls.prID"`)}
	}
	if _, ok := pc.mutation.Repo(); !ok {
		return &ValidationError{Name: "repo", err: errors.New(`ent: missing required field "Pulls.repo"`)}
	}
	if _, ok := pc.mutation.RepoID(); !ok {
		return &ValidationError{Name: "repoID", err: errors.New(`ent: missing required field "Pulls.repoID"`)}
	}
	if _, ok := pc.mutation.Number(); !ok {
		return &ValidationError{Name: "number", err: errors.New(`ent: missing required field "Pulls.number"`)}
	}
	if _, ok := pc.mutation.Comment(); !ok {
		return &ValidationError{Name: "comment", err: errors.New(`ent: missing required field "Pulls.comment"`)}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "Pulls.createdAt"`)}
	}
	if _, ok := pc.mutation.CheckRunID(); !ok {
		return &ValidationError{Name: "checkRunID", err: errors.New(`ent: missing required field "Pulls.checkRunID"`)}
	}
	if _, ok := pc.mutation.CheckRunResult(); !ok {
		return &ValidationError{Name: "checkRunResult", err: errors.New(`ent: missing required field "Pulls.checkRunResult"`)}
	}
	if _, ok := pc.mutation.HeadSha(); !ok {
		return &ValidationError{Name: "headSha", err: errors.New(`ent: missing required field "Pulls.headSha"`)}
	}
	if _, ok := pc.mutation.CreatorID(); !ok {
		return &ValidationError{Name: "Creator", err: errors.New(`ent: missing required edge "Pulls.Creator"`)}
	}
	return nil
}

func (pc *PullsCreate) sqlSave(ctx context.Context) (*Pulls, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PullsCreate) createSpec() (*Pulls, *sqlgraph.CreateSpec) {
	var (
		_node = &Pulls{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(pulls.Table, sqlgraph.NewFieldSpec(pulls.FieldID, field.TypeInt))
	)
	_spec.OnConflict = pc.conflict
	if value, ok := pc.mutation.Owner(); ok {
		_spec.SetField(pulls.FieldOwner, field.TypeString, value)
		_node.Owner = value
	}
	if value, ok := pc.mutation.PrID(); ok {
		_spec.SetField(pulls.FieldPrID, field.TypeInt64, value)
		_node.PrID = value
	}
	if value, ok := pc.mutation.Repo(); ok {
		_spec.SetField(pulls.FieldRepo, field.TypeString, value)
		_node.Repo = value
	}
	if value, ok := pc.mutation.RepoID(); ok {
		_spec.SetField(pulls.FieldRepoID, field.TypeInt64, value)
		_node.RepoID = value
	}
	if value, ok := pc.mutation.Number(); ok {
		_spec.SetField(pulls.FieldNumber, field.TypeInt, value)
		_node.Number = value
	}
	if value, ok := pc.mutation.Comment(); ok {
		_spec.SetField(pulls.FieldComment, field.TypeInt64, value)
		_node.Comment = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(pulls.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.MergedAt(); ok {
		_spec.SetField(pulls.FieldMergedAt, field.TypeTime, value)
		_node.MergedAt = value
	}
	if value, ok := pc.mutation.CheckRunID(); ok {
		_spec.SetField(pulls.FieldCheckRunID, field.TypeInt64, value)
		_node.CheckRunID = value
	}
	if value, ok := pc.mutation.CheckRunResult(); ok {
		_spec.SetField(pulls.FieldCheckRunResult, field.TypeString, value)
		_node.CheckRunResult = value
	}
	if value, ok := pc.mutation.HeadSha(); ok {
		_spec.SetField(pulls.FieldHeadSha, field.TypeString, value)
		_node.HeadSha = value
	}
	if nodes := pc.mutation.CreatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pulls.CreatorTable,
			Columns: []string{pulls.CreatorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_pull_requests = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Pulls.Create().
//		SetOwner(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PullsUpsert) {
//			SetOwner(v+v).
//		}).
//		Exec(ctx)
func (pc *PullsCreate) OnConflict(opts ...sql.ConflictOption) *PullsUpsertOne {
	pc.conflict = opts
	return &PullsUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Pulls.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *PullsCreate) OnConflictColumns(columns ...string) *PullsUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &PullsUpsertOne{
		create: pc,
	}
}

type (
	// PullsUpsertOne is the builder for "upsert"-ing
	//  one Pulls node.
	PullsUpsertOne struct {
		create *PullsCreate
	}

	// PullsUpsert is the "OnConflict" setter.
	PullsUpsert struct {
		*sql.UpdateSet
	}
)

// SetOwner sets the "owner" field.
func (u *PullsUpsert) SetOwner(v string) *PullsUpsert {
	u.Set(pulls.FieldOwner, v)
	return u
}

// UpdateOwner sets the "owner" field to the value that was provided on create.
func (u *PullsUpsert) UpdateOwner() *PullsUpsert {
	u.SetExcluded(pulls.FieldOwner)
	return u
}

// SetPrID sets the "prID" field.
func (u *PullsUpsert) SetPrID(v int64) *PullsUpsert {
	u.Set(pulls.FieldPrID, v)
	return u
}

// UpdatePrID sets the "prID" field to the value that was provided on create.
func (u *PullsUpsert) UpdatePrID() *PullsUpsert {
	u.SetExcluded(pulls.FieldPrID)
	return u
}

// AddPrID adds v to the "prID" field.
func (u *PullsUpsert) AddPrID(v int64) *PullsUpsert {
	u.Add(pulls.FieldPrID, v)
	return u
}

// SetRepo sets the "repo" field.
func (u *PullsUpsert) SetRepo(v string) *PullsUpsert {
	u.Set(pulls.FieldRepo, v)
	return u
}

// UpdateRepo sets the "repo" field to the value that was provided on create.
func (u *PullsUpsert) UpdateRepo() *PullsUpsert {
	u.SetExcluded(pulls.FieldRepo)
	return u
}

// SetRepoID sets the "repoID" field.
func (u *PullsUpsert) SetRepoID(v int64) *PullsUpsert {
	u.Set(pulls.FieldRepoID, v)
	return u
}

// UpdateRepoID sets the "repoID" field to the value that was provided on create.
func (u *PullsUpsert) UpdateRepoID() *PullsUpsert {
	u.SetExcluded(pulls.FieldRepoID)
	return u
}

// AddRepoID adds v to the "repoID" field.
func (u *PullsUpsert) AddRepoID(v int64) *PullsUpsert {
	u.Add(pulls.FieldRepoID, v)
	return u
}

// SetNumber sets the "number" field.
func (u *PullsUpsert) SetNumber(v int) *PullsUpsert {
	u.Set(pulls.FieldNumber, v)
	return u
}

// UpdateNumber sets the "number" field to the value that was provided on create.
func (u *PullsUpsert) UpdateNumber() *PullsUpsert {
	u.SetExcluded(pulls.FieldNumber)
	return u
}

// AddNumber adds v to the "number" field.
func (u *PullsUpsert) AddNumber(v int) *PullsUpsert {
	u.Add(pulls.FieldNumber, v)
	return u
}

// SetComment sets the "comment" field.
func (u *PullsUpsert) SetComment(v int64) *PullsUpsert {
	u.Set(pulls.FieldComment, v)
	return u
}

// UpdateComment sets the "comment" field to the value that was provided on create.
func (u *PullsUpsert) UpdateComment() *PullsUpsert {
	u.SetExcluded(pulls.FieldComment)
	return u
}

// AddComment adds v to the "comment" field.
func (u *PullsUpsert) AddComment(v int64) *PullsUpsert {
	u.Add(pulls.FieldComment, v)
	return u
}

// SetCreatedAt sets the "createdAt" field.
func (u *PullsUpsert) SetCreatedAt(v time.Time) *PullsUpsert {
	u.Set(pulls.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "createdAt" field to the value that was provided on create.
func (u *PullsUpsert) UpdateCreatedAt() *PullsUpsert {
	u.SetExcluded(pulls.FieldCreatedAt)
	return u
}

// SetMergedAt sets the "mergedAt" field.
func (u *PullsUpsert) SetMergedAt(v time.Time) *PullsUpsert {
	u.Set(pulls.FieldMergedAt, v)
	return u
}

// UpdateMergedAt sets the "mergedAt" field to the value that was provided on create.
func (u *PullsUpsert) UpdateMergedAt() *PullsUpsert {
	u.SetExcluded(pulls.FieldMergedAt)
	return u
}

// ClearMergedAt clears the value of the "mergedAt" field.
func (u *PullsUpsert) ClearMergedAt() *PullsUpsert {
	u.SetNull(pulls.FieldMergedAt)
	return u
}

// SetCheckRunID sets the "checkRunID" field.
func (u *PullsUpsert) SetCheckRunID(v int64) *PullsUpsert {
	u.Set(pulls.FieldCheckRunID, v)
	return u
}

// UpdateCheckRunID sets the "checkRunID" field to the value that was provided on create.
func (u *PullsUpsert) UpdateCheckRunID() *PullsUpsert {
	u.SetExcluded(pulls.FieldCheckRunID)
	return u
}

// AddCheckRunID adds v to the "checkRunID" field.
func (u *PullsUpsert) AddCheckRunID(v int64) *PullsUpsert {
	u.Add(pulls.FieldCheckRunID, v)
	return u
}

// SetCheckRunResult sets the "checkRunResult" field.
func (u *PullsUpsert) SetCheckRunResult(v string) *PullsUpsert {
	u.Set(pulls.FieldCheckRunResult, v)
	return u
}

// UpdateCheckRunResult sets the "checkRunResult" field to the value that was provided on create.
func (u *PullsUpsert) UpdateCheckRunResult() *PullsUpsert {
	u.SetExcluded(pulls.FieldCheckRunResult)
	return u
}

// SetHeadSha sets the "headSha" field.
func (u *PullsUpsert) SetHeadSha(v string) *PullsUpsert {
	u.Set(pulls.FieldHeadSha, v)
	return u
}

// UpdateHeadSha sets the "headSha" field to the value that was provided on create.
func (u *PullsUpsert) UpdateHeadSha() *PullsUpsert {
	u.SetExcluded(pulls.FieldHeadSha)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Pulls.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *PullsUpsertOne) UpdateNewValues() *PullsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Pulls.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PullsUpsertOne) Ignore() *PullsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PullsUpsertOne) DoNothing() *PullsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PullsCreate.OnConflict
// documentation for more info.
func (u *PullsUpsertOne) Update(set func(*PullsUpsert)) *PullsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PullsUpsert{UpdateSet: update})
	}))
	return u
}

// SetOwner sets the "owner" field.
func (u *PullsUpsertOne) SetOwner(v string) *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.SetOwner(v)
	})
}

// UpdateOwner sets the "owner" field to the value that was provided on create.
func (u *PullsUpsertOne) UpdateOwner() *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.UpdateOwner()
	})
}

// SetPrID sets the "prID" field.
func (u *PullsUpsertOne) SetPrID(v int64) *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.SetPrID(v)
	})
}

// AddPrID adds v to the "prID" field.
func (u *PullsUpsertOne) AddPrID(v int64) *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.AddPrID(v)
	})
}

// UpdatePrID sets the "prID" field to the value that was provided on create.
func (u *PullsUpsertOne) UpdatePrID() *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.UpdatePrID()
	})
}

// SetRepo sets the "repo" field.
func (u *PullsUpsertOne) SetRepo(v string) *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.SetRepo(v)
	})
}

// UpdateRepo sets the "repo" field to the value that was provided on create.
func (u *PullsUpsertOne) UpdateRepo() *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.UpdateRepo()
	})
}

// SetRepoID sets the "repoID" field.
func (u *PullsUpsertOne) SetRepoID(v int64) *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.SetRepoID(v)
	})
}

// AddRepoID adds v to the "repoID" field.
func (u *PullsUpsertOne) AddRepoID(v int64) *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.AddRepoID(v)
	})
}

// UpdateRepoID sets the "repoID" field to the value that was provided on create.
func (u *PullsUpsertOne) UpdateRepoID() *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.UpdateRepoID()
	})
}

// SetNumber sets the "number" field.
func (u *PullsUpsertOne) SetNumber(v int) *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.SetNumber(v)
	})
}

// AddNumber adds v to the "number" field.
func (u *PullsUpsertOne) AddNumber(v int) *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.AddNumber(v)
	})
}

// UpdateNumber sets the "number" field to the value that was provided on create.
func (u *PullsUpsertOne) UpdateNumber() *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.UpdateNumber()
	})
}

// SetComment sets the "comment" field.
func (u *PullsUpsertOne) SetComment(v int64) *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.SetComment(v)
	})
}

// AddComment adds v to the "comment" field.
func (u *PullsUpsertOne) AddComment(v int64) *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.AddComment(v)
	})
}

// UpdateComment sets the "comment" field to the value that was provided on create.
func (u *PullsUpsertOne) UpdateComment() *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.UpdateComment()
	})
}

// SetCreatedAt sets the "createdAt" field.
func (u *PullsUpsertOne) SetCreatedAt(v time.Time) *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "createdAt" field to the value that was provided on create.
func (u *PullsUpsertOne) UpdateCreatedAt() *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetMergedAt sets the "mergedAt" field.
func (u *PullsUpsertOne) SetMergedAt(v time.Time) *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.SetMergedAt(v)
	})
}

// UpdateMergedAt sets the "mergedAt" field to the value that was provided on create.
func (u *PullsUpsertOne) UpdateMergedAt() *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.UpdateMergedAt()
	})
}

// ClearMergedAt clears the value of the "mergedAt" field.
func (u *PullsUpsertOne) ClearMergedAt() *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.ClearMergedAt()
	})
}

// SetCheckRunID sets the "checkRunID" field.
func (u *PullsUpsertOne) SetCheckRunID(v int64) *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.SetCheckRunID(v)
	})
}

// AddCheckRunID adds v to the "checkRunID" field.
func (u *PullsUpsertOne) AddCheckRunID(v int64) *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.AddCheckRunID(v)
	})
}

// UpdateCheckRunID sets the "checkRunID" field to the value that was provided on create.
func (u *PullsUpsertOne) UpdateCheckRunID() *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.UpdateCheckRunID()
	})
}

// SetCheckRunResult sets the "checkRunResult" field.
func (u *PullsUpsertOne) SetCheckRunResult(v string) *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.SetCheckRunResult(v)
	})
}

// UpdateCheckRunResult sets the "checkRunResult" field to the value that was provided on create.
func (u *PullsUpsertOne) UpdateCheckRunResult() *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.UpdateCheckRunResult()
	})
}

// SetHeadSha sets the "headSha" field.
func (u *PullsUpsertOne) SetHeadSha(v string) *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.SetHeadSha(v)
	})
}

// UpdateHeadSha sets the "headSha" field to the value that was provided on create.
func (u *PullsUpsertOne) UpdateHeadSha() *PullsUpsertOne {
	return u.Update(func(s *PullsUpsert) {
		s.UpdateHeadSha()
	})
}

// Exec executes the query.
func (u *PullsUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PullsCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PullsUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PullsUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PullsUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PullsCreateBulk is the builder for creating many Pulls entities in bulk.
type PullsCreateBulk struct {
	config
	builders []*PullsCreate
	conflict []sql.ConflictOption
}

// Save creates the Pulls entities in the database.
func (pcb *PullsCreateBulk) Save(ctx context.Context) ([]*Pulls, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Pulls, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PullsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PullsCreateBulk) SaveX(ctx context.Context) []*Pulls {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PullsCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PullsCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Pulls.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PullsUpsert) {
//			SetOwner(v+v).
//		}).
//		Exec(ctx)
func (pcb *PullsCreateBulk) OnConflict(opts ...sql.ConflictOption) *PullsUpsertBulk {
	pcb.conflict = opts
	return &PullsUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Pulls.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *PullsCreateBulk) OnConflictColumns(columns ...string) *PullsUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &PullsUpsertBulk{
		create: pcb,
	}
}

// PullsUpsertBulk is the builder for "upsert"-ing
// a bulk of Pulls nodes.
type PullsUpsertBulk struct {
	create *PullsCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Pulls.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *PullsUpsertBulk) UpdateNewValues() *PullsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Pulls.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PullsUpsertBulk) Ignore() *PullsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PullsUpsertBulk) DoNothing() *PullsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PullsCreateBulk.OnConflict
// documentation for more info.
func (u *PullsUpsertBulk) Update(set func(*PullsUpsert)) *PullsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PullsUpsert{UpdateSet: update})
	}))
	return u
}

// SetOwner sets the "owner" field.
func (u *PullsUpsertBulk) SetOwner(v string) *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.SetOwner(v)
	})
}

// UpdateOwner sets the "owner" field to the value that was provided on create.
func (u *PullsUpsertBulk) UpdateOwner() *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.UpdateOwner()
	})
}

// SetPrID sets the "prID" field.
func (u *PullsUpsertBulk) SetPrID(v int64) *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.SetPrID(v)
	})
}

// AddPrID adds v to the "prID" field.
func (u *PullsUpsertBulk) AddPrID(v int64) *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.AddPrID(v)
	})
}

// UpdatePrID sets the "prID" field to the value that was provided on create.
func (u *PullsUpsertBulk) UpdatePrID() *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.UpdatePrID()
	})
}

// SetRepo sets the "repo" field.
func (u *PullsUpsertBulk) SetRepo(v string) *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.SetRepo(v)
	})
}

// UpdateRepo sets the "repo" field to the value that was provided on create.
func (u *PullsUpsertBulk) UpdateRepo() *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.UpdateRepo()
	})
}

// SetRepoID sets the "repoID" field.
func (u *PullsUpsertBulk) SetRepoID(v int64) *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.SetRepoID(v)
	})
}

// AddRepoID adds v to the "repoID" field.
func (u *PullsUpsertBulk) AddRepoID(v int64) *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.AddRepoID(v)
	})
}

// UpdateRepoID sets the "repoID" field to the value that was provided on create.
func (u *PullsUpsertBulk) UpdateRepoID() *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.UpdateRepoID()
	})
}

// SetNumber sets the "number" field.
func (u *PullsUpsertBulk) SetNumber(v int) *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.SetNumber(v)
	})
}

// AddNumber adds v to the "number" field.
func (u *PullsUpsertBulk) AddNumber(v int) *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.AddNumber(v)
	})
}

// UpdateNumber sets the "number" field to the value that was provided on create.
func (u *PullsUpsertBulk) UpdateNumber() *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.UpdateNumber()
	})
}

// SetComment sets the "comment" field.
func (u *PullsUpsertBulk) SetComment(v int64) *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.SetComment(v)
	})
}

// AddComment adds v to the "comment" field.
func (u *PullsUpsertBulk) AddComment(v int64) *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.AddComment(v)
	})
}

// UpdateComment sets the "comment" field to the value that was provided on create.
func (u *PullsUpsertBulk) UpdateComment() *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.UpdateComment()
	})
}

// SetCreatedAt sets the "createdAt" field.
func (u *PullsUpsertBulk) SetCreatedAt(v time.Time) *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "createdAt" field to the value that was provided on create.
func (u *PullsUpsertBulk) UpdateCreatedAt() *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetMergedAt sets the "mergedAt" field.
func (u *PullsUpsertBulk) SetMergedAt(v time.Time) *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.SetMergedAt(v)
	})
}

// UpdateMergedAt sets the "mergedAt" field to the value that was provided on create.
func (u *PullsUpsertBulk) UpdateMergedAt() *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.UpdateMergedAt()
	})
}

// ClearMergedAt clears the value of the "mergedAt" field.
func (u *PullsUpsertBulk) ClearMergedAt() *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.ClearMergedAt()
	})
}

// SetCheckRunID sets the "checkRunID" field.
func (u *PullsUpsertBulk) SetCheckRunID(v int64) *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.SetCheckRunID(v)
	})
}

// AddCheckRunID adds v to the "checkRunID" field.
func (u *PullsUpsertBulk) AddCheckRunID(v int64) *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.AddCheckRunID(v)
	})
}

// UpdateCheckRunID sets the "checkRunID" field to the value that was provided on create.
func (u *PullsUpsertBulk) UpdateCheckRunID() *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.UpdateCheckRunID()
	})
}

// SetCheckRunResult sets the "checkRunResult" field.
func (u *PullsUpsertBulk) SetCheckRunResult(v string) *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.SetCheckRunResult(v)
	})
}

// UpdateCheckRunResult sets the "checkRunResult" field to the value that was provided on create.
func (u *PullsUpsertBulk) UpdateCheckRunResult() *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.UpdateCheckRunResult()
	})
}

// SetHeadSha sets the "headSha" field.
func (u *PullsUpsertBulk) SetHeadSha(v string) *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.SetHeadSha(v)
	})
}

// UpdateHeadSha sets the "headSha" field to the value that was provided on create.
func (u *PullsUpsertBulk) UpdateHeadSha() *PullsUpsertBulk {
	return u.Update(func(s *PullsUpsert) {
		s.UpdateHeadSha()
	})
}

// Exec executes the query.
func (u *PullsUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PullsCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PullsCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PullsUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
