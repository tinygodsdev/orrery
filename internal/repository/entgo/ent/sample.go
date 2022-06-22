// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/DanielTitkov/lentils/internal/domain"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/sample"
	"github.com/google/uuid"
)

// Sample is the model entity for the Sample schema.
type Sample struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Criteria holds the value of the "criteria" field.
	Criteria domain.SampleCriteria `json:"criteria,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SampleQuery when eager-loading is set.
	Edges SampleEdges `json:"edges"`
}

// SampleEdges holds the relations/edges for other nodes in the graph.
type SampleEdges struct {
	// Norms holds the value of the norms edge.
	Norms []*Norm `json:"norms,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// NormsOrErr returns the Norms value or an error if the edge
// was not loaded in eager-loading.
func (e SampleEdges) NormsOrErr() ([]*Norm, error) {
	if e.loadedTypes[0] {
		return e.Norms, nil
	}
	return nil, &NotLoadedError{edge: "norms"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Sample) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case sample.FieldCriteria:
			values[i] = new([]byte)
		case sample.FieldCode:
			values[i] = new(sql.NullString)
		case sample.FieldCreateTime, sample.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case sample.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Sample", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Sample fields.
func (s *Sample) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sample.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case sample.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				s.CreateTime = value.Time
			}
		case sample.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				s.UpdateTime = value.Time
			}
		case sample.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				s.Code = value.String
			}
		case sample.FieldCriteria:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field criteria", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.Criteria); err != nil {
					return fmt.Errorf("unmarshal field criteria: %w", err)
				}
			}
		}
	}
	return nil
}

// QueryNorms queries the "norms" edge of the Sample entity.
func (s *Sample) QueryNorms() *NormQuery {
	return (&SampleClient{config: s.config}).QueryNorms(s)
}

// Update returns a builder for updating this Sample.
// Note that you need to call Sample.Unwrap() before calling this method if this Sample
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Sample) Update() *SampleUpdateOne {
	return (&SampleClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Sample entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Sample) Unwrap() *Sample {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Sample is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Sample) String() string {
	var builder strings.Builder
	builder.WriteString("Sample(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("create_time=")
	builder.WriteString(s.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(s.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(s.Code)
	builder.WriteString(", ")
	builder.WriteString("criteria=")
	builder.WriteString(fmt.Sprintf("%v", s.Criteria))
	builder.WriteByte(')')
	return builder.String()
}

// Samples is a parsable slice of Sample.
type Samples []*Sample

func (s Samples) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
