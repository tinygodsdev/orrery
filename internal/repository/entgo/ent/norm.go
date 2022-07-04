// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/norm"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/sample"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/scale"
	"github.com/google/uuid"
)

// Norm is the model entity for the Norm schema.
type Norm struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Base holds the value of the "base" field.
	Base int `json:"base,omitempty"`
	// Mean holds the value of the "mean" field.
	Mean float64 `json:"mean,omitempty"`
	// Sigma holds the value of the "sigma" field.
	Sigma float64 `json:"sigma,omitempty"`
	// Rank holds the value of the "rank" field.
	Rank int `json:"rank,omitempty"`
	// Meta holds the value of the "meta" field.
	Meta map[string]interface{} `json:"meta,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NormQuery when eager-loading is set.
	Edges        NormEdges `json:"edges"`
	sample_norms *uuid.UUID
	scale_norms  *uuid.UUID
}

// NormEdges holds the relations/edges for other nodes in the graph.
type NormEdges struct {
	// Sample holds the value of the sample edge.
	Sample *Sample `json:"sample,omitempty"`
	// Scale holds the value of the scale edge.
	Scale *Scale `json:"scale,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// SampleOrErr returns the Sample value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NormEdges) SampleOrErr() (*Sample, error) {
	if e.loadedTypes[0] {
		if e.Sample == nil {
			// The edge sample was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: sample.Label}
		}
		return e.Sample, nil
	}
	return nil, &NotLoadedError{edge: "sample"}
}

// ScaleOrErr returns the Scale value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NormEdges) ScaleOrErr() (*Scale, error) {
	if e.loadedTypes[1] {
		if e.Scale == nil {
			// The edge scale was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: scale.Label}
		}
		return e.Scale, nil
	}
	return nil, &NotLoadedError{edge: "scale"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Norm) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case norm.FieldMeta:
			values[i] = new([]byte)
		case norm.FieldMean, norm.FieldSigma:
			values[i] = new(sql.NullFloat64)
		case norm.FieldBase, norm.FieldRank:
			values[i] = new(sql.NullInt64)
		case norm.FieldName:
			values[i] = new(sql.NullString)
		case norm.FieldCreateTime, norm.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case norm.FieldID:
			values[i] = new(uuid.UUID)
		case norm.ForeignKeys[0]: // sample_norms
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case norm.ForeignKeys[1]: // scale_norms
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Norm", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Norm fields.
func (n *Norm) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case norm.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				n.ID = *value
			}
		case norm.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				n.CreateTime = value.Time
			}
		case norm.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				n.UpdateTime = value.Time
			}
		case norm.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				n.Name = value.String
			}
		case norm.FieldBase:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field base", values[i])
			} else if value.Valid {
				n.Base = int(value.Int64)
			}
		case norm.FieldMean:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field mean", values[i])
			} else if value.Valid {
				n.Mean = value.Float64
			}
		case norm.FieldSigma:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field sigma", values[i])
			} else if value.Valid {
				n.Sigma = value.Float64
			}
		case norm.FieldRank:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rank", values[i])
			} else if value.Valid {
				n.Rank = int(value.Int64)
			}
		case norm.FieldMeta:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field meta", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &n.Meta); err != nil {
					return fmt.Errorf("unmarshal field meta: %w", err)
				}
			}
		case norm.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field sample_norms", values[i])
			} else if value.Valid {
				n.sample_norms = new(uuid.UUID)
				*n.sample_norms = *value.S.(*uuid.UUID)
			}
		case norm.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field scale_norms", values[i])
			} else if value.Valid {
				n.scale_norms = new(uuid.UUID)
				*n.scale_norms = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QuerySample queries the "sample" edge of the Norm entity.
func (n *Norm) QuerySample() *SampleQuery {
	return (&NormClient{config: n.config}).QuerySample(n)
}

// QueryScale queries the "scale" edge of the Norm entity.
func (n *Norm) QueryScale() *ScaleQuery {
	return (&NormClient{config: n.config}).QueryScale(n)
}

// Update returns a builder for updating this Norm.
// Note that you need to call Norm.Unwrap() before calling this method if this Norm
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Norm) Update() *NormUpdateOne {
	return (&NormClient{config: n.config}).UpdateOne(n)
}

// Unwrap unwraps the Norm entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Norm) Unwrap() *Norm {
	_tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Norm is not a transactional entity")
	}
	n.config.driver = _tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Norm) String() string {
	var builder strings.Builder
	builder.WriteString("Norm(")
	builder.WriteString(fmt.Sprintf("id=%v, ", n.ID))
	builder.WriteString("create_time=")
	builder.WriteString(n.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(n.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(n.Name)
	builder.WriteString(", ")
	builder.WriteString("base=")
	builder.WriteString(fmt.Sprintf("%v", n.Base))
	builder.WriteString(", ")
	builder.WriteString("mean=")
	builder.WriteString(fmt.Sprintf("%v", n.Mean))
	builder.WriteString(", ")
	builder.WriteString("sigma=")
	builder.WriteString(fmt.Sprintf("%v", n.Sigma))
	builder.WriteString(", ")
	builder.WriteString("rank=")
	builder.WriteString(fmt.Sprintf("%v", n.Rank))
	builder.WriteString(", ")
	builder.WriteString("meta=")
	builder.WriteString(fmt.Sprintf("%v", n.Meta))
	builder.WriteByte(')')
	return builder.String()
}

// Norms is a parsable slice of Norm.
type Norms []*Norm

func (n Norms) config(cfg config) {
	for _i := range n {
		n[_i].config = cfg
	}
}
