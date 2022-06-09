// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/item"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/response"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/take"
	"github.com/google/uuid"
)

// Response is the model entity for the Response schema.
type Response struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Value holds the value of the "value" field.
	Value int `json:"value,omitempty"`
	// Meta holds the value of the "meta" field.
	Meta map[string]interface{} `json:"meta,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ResponseQuery when eager-loading is set.
	Edges          ResponseEdges `json:"edges"`
	item_responses *uuid.UUID
	take_responses *uuid.UUID
}

// ResponseEdges holds the relations/edges for other nodes in the graph.
type ResponseEdges struct {
	// Item holds the value of the item edge.
	Item *Item `json:"item,omitempty"`
	// Take holds the value of the take edge.
	Take *Take `json:"take,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ItemOrErr returns the Item value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResponseEdges) ItemOrErr() (*Item, error) {
	if e.loadedTypes[0] {
		if e.Item == nil {
			// The edge item was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: item.Label}
		}
		return e.Item, nil
	}
	return nil, &NotLoadedError{edge: "item"}
}

// TakeOrErr returns the Take value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResponseEdges) TakeOrErr() (*Take, error) {
	if e.loadedTypes[1] {
		if e.Take == nil {
			// The edge take was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: take.Label}
		}
		return e.Take, nil
	}
	return nil, &NotLoadedError{edge: "take"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Response) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case response.FieldMeta:
			values[i] = new([]byte)
		case response.FieldValue:
			values[i] = new(sql.NullInt64)
		case response.FieldCreateTime, response.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case response.FieldID:
			values[i] = new(uuid.UUID)
		case response.ForeignKeys[0]: // item_responses
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case response.ForeignKeys[1]: // take_responses
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Response", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Response fields.
func (r *Response) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case response.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				r.ID = *value
			}
		case response.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				r.CreateTime = value.Time
			}
		case response.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				r.UpdateTime = value.Time
			}
		case response.FieldValue:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				r.Value = int(value.Int64)
			}
		case response.FieldMeta:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field meta", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.Meta); err != nil {
					return fmt.Errorf("unmarshal field meta: %w", err)
				}
			}
		case response.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field item_responses", values[i])
			} else if value.Valid {
				r.item_responses = new(uuid.UUID)
				*r.item_responses = *value.S.(*uuid.UUID)
			}
		case response.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field take_responses", values[i])
			} else if value.Valid {
				r.take_responses = new(uuid.UUID)
				*r.take_responses = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryItem queries the "item" edge of the Response entity.
func (r *Response) QueryItem() *ItemQuery {
	return (&ResponseClient{config: r.config}).QueryItem(r)
}

// QueryTake queries the "take" edge of the Response entity.
func (r *Response) QueryTake() *TakeQuery {
	return (&ResponseClient{config: r.config}).QueryTake(r)
}

// Update returns a builder for updating this Response.
// Note that you need to call Response.Unwrap() before calling this method if this Response
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Response) Update() *ResponseUpdateOne {
	return (&ResponseClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Response entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Response) Unwrap() *Response {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Response is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Response) String() string {
	var builder strings.Builder
	builder.WriteString("Response(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("create_time=")
	builder.WriteString(r.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(r.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(fmt.Sprintf("%v", r.Value))
	builder.WriteString(", ")
	builder.WriteString("meta=")
	builder.WriteString(fmt.Sprintf("%v", r.Meta))
	builder.WriteByte(')')
	return builder.String()
}

// Responses is a parsable slice of Response.
type Responses []*Response

func (r Responses) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
