// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/scale"
	"github.com/google/uuid"
)

// Scale is the model entity for the Scale schema.
type Scale struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Global holds the value of the "global" field.
	Global bool `json:"global,omitempty"`
	// Type holds the value of the "type" field.
	Type scale.Type `json:"type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ScaleQuery when eager-loading is set.
	Edges ScaleEdges `json:"edges"`
}

// ScaleEdges holds the relations/edges for other nodes in the graph.
type ScaleEdges struct {
	// Items holds the value of the items edge.
	Items []*Item `json:"items,omitempty"`
	// Interpretations holds the value of the interpretations edge.
	Interpretations []*Interpretation `json:"interpretations,omitempty"`
	// Translations holds the value of the translations edge.
	Translations []*ScaleTranslation `json:"translations,omitempty"`
	// Norms holds the value of the norms edge.
	Norms []*Norm `json:"norms,omitempty"`
	// Results holds the value of the results edge.
	Results []*Result `json:"results,omitempty"`
	// Test holds the value of the test edge.
	Test []*Test `json:"test,omitempty"`
	// ScaleItem holds the value of the scale_item edge.
	ScaleItem []*ScaleItem `json:"scale_item,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [7]bool
}

// ItemsOrErr returns the Items value or an error if the edge
// was not loaded in eager-loading.
func (e ScaleEdges) ItemsOrErr() ([]*Item, error) {
	if e.loadedTypes[0] {
		return e.Items, nil
	}
	return nil, &NotLoadedError{edge: "items"}
}

// InterpretationsOrErr returns the Interpretations value or an error if the edge
// was not loaded in eager-loading.
func (e ScaleEdges) InterpretationsOrErr() ([]*Interpretation, error) {
	if e.loadedTypes[1] {
		return e.Interpretations, nil
	}
	return nil, &NotLoadedError{edge: "interpretations"}
}

// TranslationsOrErr returns the Translations value or an error if the edge
// was not loaded in eager-loading.
func (e ScaleEdges) TranslationsOrErr() ([]*ScaleTranslation, error) {
	if e.loadedTypes[2] {
		return e.Translations, nil
	}
	return nil, &NotLoadedError{edge: "translations"}
}

// NormsOrErr returns the Norms value or an error if the edge
// was not loaded in eager-loading.
func (e ScaleEdges) NormsOrErr() ([]*Norm, error) {
	if e.loadedTypes[3] {
		return e.Norms, nil
	}
	return nil, &NotLoadedError{edge: "norms"}
}

// ResultsOrErr returns the Results value or an error if the edge
// was not loaded in eager-loading.
func (e ScaleEdges) ResultsOrErr() ([]*Result, error) {
	if e.loadedTypes[4] {
		return e.Results, nil
	}
	return nil, &NotLoadedError{edge: "results"}
}

// TestOrErr returns the Test value or an error if the edge
// was not loaded in eager-loading.
func (e ScaleEdges) TestOrErr() ([]*Test, error) {
	if e.loadedTypes[5] {
		return e.Test, nil
	}
	return nil, &NotLoadedError{edge: "test"}
}

// ScaleItemOrErr returns the ScaleItem value or an error if the edge
// was not loaded in eager-loading.
func (e ScaleEdges) ScaleItemOrErr() ([]*ScaleItem, error) {
	if e.loadedTypes[6] {
		return e.ScaleItem, nil
	}
	return nil, &NotLoadedError{edge: "scale_item"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Scale) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case scale.FieldGlobal:
			values[i] = new(sql.NullBool)
		case scale.FieldCode, scale.FieldType:
			values[i] = new(sql.NullString)
		case scale.FieldCreateTime, scale.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case scale.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Scale", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Scale fields.
func (s *Scale) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case scale.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case scale.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				s.CreateTime = value.Time
			}
		case scale.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				s.UpdateTime = value.Time
			}
		case scale.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				s.Code = value.String
			}
		case scale.FieldGlobal:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field global", values[i])
			} else if value.Valid {
				s.Global = value.Bool
			}
		case scale.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				s.Type = scale.Type(value.String)
			}
		}
	}
	return nil
}

// QueryItems queries the "items" edge of the Scale entity.
func (s *Scale) QueryItems() *ItemQuery {
	return (&ScaleClient{config: s.config}).QueryItems(s)
}

// QueryInterpretations queries the "interpretations" edge of the Scale entity.
func (s *Scale) QueryInterpretations() *InterpretationQuery {
	return (&ScaleClient{config: s.config}).QueryInterpretations(s)
}

// QueryTranslations queries the "translations" edge of the Scale entity.
func (s *Scale) QueryTranslations() *ScaleTranslationQuery {
	return (&ScaleClient{config: s.config}).QueryTranslations(s)
}

// QueryNorms queries the "norms" edge of the Scale entity.
func (s *Scale) QueryNorms() *NormQuery {
	return (&ScaleClient{config: s.config}).QueryNorms(s)
}

// QueryResults queries the "results" edge of the Scale entity.
func (s *Scale) QueryResults() *ResultQuery {
	return (&ScaleClient{config: s.config}).QueryResults(s)
}

// QueryTest queries the "test" edge of the Scale entity.
func (s *Scale) QueryTest() *TestQuery {
	return (&ScaleClient{config: s.config}).QueryTest(s)
}

// QueryScaleItem queries the "scale_item" edge of the Scale entity.
func (s *Scale) QueryScaleItem() *ScaleItemQuery {
	return (&ScaleClient{config: s.config}).QueryScaleItem(s)
}

// Update returns a builder for updating this Scale.
// Note that you need to call Scale.Unwrap() before calling this method if this Scale
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Scale) Update() *ScaleUpdateOne {
	return (&ScaleClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Scale entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Scale) Unwrap() *Scale {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Scale is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Scale) String() string {
	var builder strings.Builder
	builder.WriteString("Scale(")
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
	builder.WriteString("global=")
	builder.WriteString(fmt.Sprintf("%v", s.Global))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", s.Type))
	builder.WriteByte(')')
	return builder.String()
}

// Scales is a parsable slice of Scale.
type Scales []*Scale

func (s Scales) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
