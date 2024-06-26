// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"github-bot/ent/pulls"
	"github-bot/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Pulls is the model entity for the Pulls schema.
type Pulls struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Owner holds the value of the "owner" field.
	Owner string `json:"owner,omitempty"`
	// PrID holds the value of the "prID" field.
	PrID int64 `json:"prID,omitempty"`
	// Repo holds the value of the "repo" field.
	Repo string `json:"repo,omitempty"`
	// RepoID holds the value of the "repoID" field.
	RepoID int64 `json:"repoID,omitempty"`
	// pr number
	Number int `json:"number,omitempty"`
	// bot comment id, nil present un-comment Pulls
	Comment int64 `json:"comment,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// MergedAt holds the value of the "mergedAt" field.
	MergedAt time.Time `json:"mergedAt,omitempty"`
	// CheckRunID holds the value of the "checkRunID" field.
	CheckRunID int64 `json:"checkRunID,omitempty"`
	// CheckRunResult holds the value of the "checkRunResult" field.
	CheckRunResult string `json:"checkRunResult,omitempty"`
	// HeadSha holds the value of the "headSha" field.
	HeadSha string `json:"headSha,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PullsQuery when eager-loading is set.
	Edges              PullsEdges `json:"edges"`
	user_pull_requests *int
	selectValues       sql.SelectValues
}

// PullsEdges holds the relations/edges for other nodes in the graph.
type PullsEdges struct {
	// Creator holds the value of the Creator edge.
	Creator *User `json:"Creator,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CreatorOrErr returns the Creator value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PullsEdges) CreatorOrErr() (*User, error) {
	if e.Creator != nil {
		return e.Creator, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "Creator"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Pulls) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case pulls.FieldID, pulls.FieldPrID, pulls.FieldRepoID, pulls.FieldNumber, pulls.FieldComment, pulls.FieldCheckRunID:
			values[i] = new(sql.NullInt64)
		case pulls.FieldOwner, pulls.FieldRepo, pulls.FieldCheckRunResult, pulls.FieldHeadSha:
			values[i] = new(sql.NullString)
		case pulls.FieldCreatedAt, pulls.FieldMergedAt:
			values[i] = new(sql.NullTime)
		case pulls.ForeignKeys[0]: // user_pull_requests
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pulls fields.
func (pu *Pulls) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pulls.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pu.ID = int(value.Int64)
		case pulls.FieldOwner:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner", values[i])
			} else if value.Valid {
				pu.Owner = value.String
			}
		case pulls.FieldPrID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field prID", values[i])
			} else if value.Valid {
				pu.PrID = value.Int64
			}
		case pulls.FieldRepo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field repo", values[i])
			} else if value.Valid {
				pu.Repo = value.String
			}
		case pulls.FieldRepoID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field repoID", values[i])
			} else if value.Valid {
				pu.RepoID = value.Int64
			}
		case pulls.FieldNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field number", values[i])
			} else if value.Valid {
				pu.Number = int(value.Int64)
			}
		case pulls.FieldComment:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field comment", values[i])
			} else if value.Valid {
				pu.Comment = value.Int64
			}
		case pulls.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				pu.CreatedAt = value.Time
			}
		case pulls.FieldMergedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field mergedAt", values[i])
			} else if value.Valid {
				pu.MergedAt = value.Time
			}
		case pulls.FieldCheckRunID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field checkRunID", values[i])
			} else if value.Valid {
				pu.CheckRunID = value.Int64
			}
		case pulls.FieldCheckRunResult:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field checkRunResult", values[i])
			} else if value.Valid {
				pu.CheckRunResult = value.String
			}
		case pulls.FieldHeadSha:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field headSha", values[i])
			} else if value.Valid {
				pu.HeadSha = value.String
			}
		case pulls.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_pull_requests", value)
			} else if value.Valid {
				pu.user_pull_requests = new(int)
				*pu.user_pull_requests = int(value.Int64)
			}
		default:
			pu.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Pulls.
// This includes values selected through modifiers, order, etc.
func (pu *Pulls) Value(name string) (ent.Value, error) {
	return pu.selectValues.Get(name)
}

// QueryCreator queries the "Creator" edge of the Pulls entity.
func (pu *Pulls) QueryCreator() *UserQuery {
	return NewPullsClient(pu.config).QueryCreator(pu)
}

// Update returns a builder for updating this Pulls.
// Note that you need to call Pulls.Unwrap() before calling this method if this Pulls
// was returned from a transaction, and the transaction was committed or rolled back.
func (pu *Pulls) Update() *PullsUpdateOne {
	return NewPullsClient(pu.config).UpdateOne(pu)
}

// Unwrap unwraps the Pulls entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pu *Pulls) Unwrap() *Pulls {
	_tx, ok := pu.config.driver.(*txDriver)
	if !ok {
		panic("ent: Pulls is not a transactional entity")
	}
	pu.config.driver = _tx.drv
	return pu
}

// String implements the fmt.Stringer.
func (pu *Pulls) String() string {
	var builder strings.Builder
	builder.WriteString("Pulls(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pu.ID))
	builder.WriteString("owner=")
	builder.WriteString(pu.Owner)
	builder.WriteString(", ")
	builder.WriteString("prID=")
	builder.WriteString(fmt.Sprintf("%v", pu.PrID))
	builder.WriteString(", ")
	builder.WriteString("repo=")
	builder.WriteString(pu.Repo)
	builder.WriteString(", ")
	builder.WriteString("repoID=")
	builder.WriteString(fmt.Sprintf("%v", pu.RepoID))
	builder.WriteString(", ")
	builder.WriteString("number=")
	builder.WriteString(fmt.Sprintf("%v", pu.Number))
	builder.WriteString(", ")
	builder.WriteString("comment=")
	builder.WriteString(fmt.Sprintf("%v", pu.Comment))
	builder.WriteString(", ")
	builder.WriteString("createdAt=")
	builder.WriteString(pu.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("mergedAt=")
	builder.WriteString(pu.MergedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("checkRunID=")
	builder.WriteString(fmt.Sprintf("%v", pu.CheckRunID))
	builder.WriteString(", ")
	builder.WriteString("checkRunResult=")
	builder.WriteString(pu.CheckRunResult)
	builder.WriteString(", ")
	builder.WriteString("headSha=")
	builder.WriteString(pu.HeadSha)
	builder.WriteByte(')')
	return builder.String()
}

// PullsSlice is a parsable slice of Pulls.
type PullsSlice []*Pulls
