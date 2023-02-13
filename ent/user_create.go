// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"github-bot/ent/pulls"
	"github-bot/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetGithubID sets the "github_id" field.
func (uc *UserCreate) SetGithubID(i int64) *UserCreate {
	uc.mutation.SetGithubID(i)
	return uc
}

// SetBangumiID sets the "bangumi_id" field.
func (uc *UserCreate) SetBangumiID(i int64) *UserCreate {
	uc.mutation.SetBangumiID(i)
	return uc
}

// SetNillableBangumiID sets the "bangumi_id" field if the given value is not nil.
func (uc *UserCreate) SetNillableBangumiID(i *int64) *UserCreate {
	if i != nil {
		uc.SetBangumiID(*i)
	}
	return uc
}

// AddPullRequestIDs adds the "pull_requests" edge to the Pulls entity by IDs.
func (uc *UserCreate) AddPullRequestIDs(ids ...int) *UserCreate {
	uc.mutation.AddPullRequestIDs(ids...)
	return uc
}

// AddPullRequests adds the "pull_requests" edges to the Pulls entity.
func (uc *UserCreate) AddPullRequests(p ...*Pulls) *UserCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uc.AddPullRequestIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	return withHooks[*User, UserMutation](ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.GithubID(); !ok {
		return &ValidationError{Name: "github_id", err: errors.New(`ent: missing required field "User.github_id"`)}
	}
	if v, ok := uc.mutation.GithubID(); ok {
		if err := user.GithubIDValidator(v); err != nil {
			return &ValidationError{Name: "github_id", err: fmt.Errorf(`ent: validator failed for field "User.github_id": %w`, err)}
		}
	}
	if v, ok := uc.mutation.BangumiID(); ok {
		if err := user.BangumiIDValidator(v); err != nil {
			return &ValidationError{Name: "bangumi_id", err: fmt.Errorf(`ent: validator failed for field "User.bangumi_id": %w`, err)}
		}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	)
	_spec.OnConflict = uc.conflict
	if value, ok := uc.mutation.GithubID(); ok {
		_spec.SetField(user.FieldGithubID, field.TypeInt64, value)
		_node.GithubID = value
	}
	if value, ok := uc.mutation.BangumiID(); ok {
		_spec.SetField(user.FieldBangumiID, field.TypeInt64, value)
		_node.BangumiID = value
	}
	if nodes := uc.mutation.PullRequestsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PullRequestsTable,
			Columns: []string{user.PullRequestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pulls.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.Create().
//		SetGithubID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetGithubID(v+v).
//		}).
//		Exec(ctx)
func (uc *UserCreate) OnConflict(opts ...sql.ConflictOption) *UserUpsertOne {
	uc.conflict = opts
	return &UserUpsertOne{
		create: uc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (uc *UserCreate) OnConflictColumns(columns ...string) *UserUpsertOne {
	uc.conflict = append(uc.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertOne{
		create: uc,
	}
}

type (
	// UserUpsertOne is the builder for "upsert"-ing
	//  one User node.
	UserUpsertOne struct {
		create *UserCreate
	}

	// UserUpsert is the "OnConflict" setter.
	UserUpsert struct {
		*sql.UpdateSet
	}
)

// SetGithubID sets the "github_id" field.
func (u *UserUpsert) SetGithubID(v int64) *UserUpsert {
	u.Set(user.FieldGithubID, v)
	return u
}

// UpdateGithubID sets the "github_id" field to the value that was provided on create.
func (u *UserUpsert) UpdateGithubID() *UserUpsert {
	u.SetExcluded(user.FieldGithubID)
	return u
}

// AddGithubID adds v to the "github_id" field.
func (u *UserUpsert) AddGithubID(v int64) *UserUpsert {
	u.Add(user.FieldGithubID, v)
	return u
}

// SetBangumiID sets the "bangumi_id" field.
func (u *UserUpsert) SetBangumiID(v int64) *UserUpsert {
	u.Set(user.FieldBangumiID, v)
	return u
}

// UpdateBangumiID sets the "bangumi_id" field to the value that was provided on create.
func (u *UserUpsert) UpdateBangumiID() *UserUpsert {
	u.SetExcluded(user.FieldBangumiID)
	return u
}

// AddBangumiID adds v to the "bangumi_id" field.
func (u *UserUpsert) AddBangumiID(v int64) *UserUpsert {
	u.Add(user.FieldBangumiID, v)
	return u
}

// ClearBangumiID clears the value of the "bangumi_id" field.
func (u *UserUpsert) ClearBangumiID() *UserUpsert {
	u.SetNull(user.FieldBangumiID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserUpsertOne) UpdateNewValues() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UserUpsertOne) Ignore() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertOne) DoNothing() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreate.OnConflict
// documentation for more info.
func (u *UserUpsertOne) Update(set func(*UserUpsert)) *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetGithubID sets the "github_id" field.
func (u *UserUpsertOne) SetGithubID(v int64) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetGithubID(v)
	})
}

// AddGithubID adds v to the "github_id" field.
func (u *UserUpsertOne) AddGithubID(v int64) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.AddGithubID(v)
	})
}

// UpdateGithubID sets the "github_id" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateGithubID() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateGithubID()
	})
}

// SetBangumiID sets the "bangumi_id" field.
func (u *UserUpsertOne) SetBangumiID(v int64) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetBangumiID(v)
	})
}

// AddBangumiID adds v to the "bangumi_id" field.
func (u *UserUpsertOne) AddBangumiID(v int64) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.AddBangumiID(v)
	})
}

// UpdateBangumiID sets the "bangumi_id" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateBangumiID() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateBangumiID()
	})
}

// ClearBangumiID clears the value of the "bangumi_id" field.
func (u *UserUpsertOne) ClearBangumiID() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearBangumiID()
	})
}

// Exec executes the query.
func (u *UserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
	conflict []sql.ConflictOption
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetGithubID(v+v).
//		}).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserUpsertBulk {
	ucb.conflict = opts
	return &UserUpsertBulk{
		create: ucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflictColumns(columns ...string) *UserUpsertBulk {
	ucb.conflict = append(ucb.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertBulk{
		create: ucb,
	}
}

// UserUpsertBulk is the builder for "upsert"-ing
// a bulk of User nodes.
type UserUpsertBulk struct {
	create *UserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserUpsertBulk) UpdateNewValues() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UserUpsertBulk) Ignore() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertBulk) DoNothing() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreateBulk.OnConflict
// documentation for more info.
func (u *UserUpsertBulk) Update(set func(*UserUpsert)) *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetGithubID sets the "github_id" field.
func (u *UserUpsertBulk) SetGithubID(v int64) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetGithubID(v)
	})
}

// AddGithubID adds v to the "github_id" field.
func (u *UserUpsertBulk) AddGithubID(v int64) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.AddGithubID(v)
	})
}

// UpdateGithubID sets the "github_id" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateGithubID() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateGithubID()
	})
}

// SetBangumiID sets the "bangumi_id" field.
func (u *UserUpsertBulk) SetBangumiID(v int64) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetBangumiID(v)
	})
}

// AddBangumiID adds v to the "bangumi_id" field.
func (u *UserUpsertBulk) AddBangumiID(v int64) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.AddBangumiID(v)
	})
}

// UpdateBangumiID sets the "bangumi_id" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateBangumiID() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateBangumiID()
	})
}

// ClearBangumiID clears the value of the "bangumi_id" field.
func (u *UserUpsertBulk) ClearBangumiID() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearBangumiID()
	})
}

// Exec executes the query.
func (u *UserUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
