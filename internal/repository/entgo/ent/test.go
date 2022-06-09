// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/test"
	"github.com/google/uuid"
)

// Test is the model entity for the Test schema.
type Test struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Published holds the value of the "published" field.
	Published bool `json:"published,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TestQuery when eager-loading is set.
	Edges TestEdges `json:"edges"`
}

// TestEdges holds the relations/edges for other nodes in the graph.
type TestEdges struct {
	// Takes holds the value of the takes edge.
	Takes []*Take `json:"takes,omitempty"`
	// Questions holds the value of the questions edge.
	Questions []*Question `json:"questions,omitempty"`
	// Translations holds the value of the translations edge.
	Translations []*TestTranslation `json:"translations,omitempty"`
	// Scales holds the value of the scales edge.
	Scales []*Scale `json:"scales,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// TakesOrErr returns the Takes value or an error if the edge
// was not loaded in eager-loading.
func (e TestEdges) TakesOrErr() ([]*Take, error) {
	if e.loadedTypes[0] {
		return e.Takes, nil
	}
	return nil, &NotLoadedError{edge: "takes"}
}

// QuestionsOrErr returns the Questions value or an error if the edge
// was not loaded in eager-loading.
func (e TestEdges) QuestionsOrErr() ([]*Question, error) {
	if e.loadedTypes[1] {
		return e.Questions, nil
	}
	return nil, &NotLoadedError{edge: "questions"}
}

// TranslationsOrErr returns the Translations value or an error if the edge
// was not loaded in eager-loading.
func (e TestEdges) TranslationsOrErr() ([]*TestTranslation, error) {
	if e.loadedTypes[2] {
		return e.Translations, nil
	}
	return nil, &NotLoadedError{edge: "translations"}
}

// ScalesOrErr returns the Scales value or an error if the edge
// was not loaded in eager-loading.
func (e TestEdges) ScalesOrErr() ([]*Scale, error) {
	if e.loadedTypes[3] {
		return e.Scales, nil
	}
	return nil, &NotLoadedError{edge: "scales"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Test) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case test.FieldPublished:
			values[i] = new(sql.NullBool)
		case test.FieldCode:
			values[i] = new(sql.NullString)
		case test.FieldCreateTime, test.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case test.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Test", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Test fields.
func (t *Test) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case test.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case test.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				t.CreateTime = value.Time
			}
		case test.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				t.UpdateTime = value.Time
			}
		case test.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				t.Code = value.String
			}
		case test.FieldPublished:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field published", values[i])
			} else if value.Valid {
				t.Published = value.Bool
			}
		}
	}
	return nil
}

// QueryTakes queries the "takes" edge of the Test entity.
func (t *Test) QueryTakes() *TakeQuery {
	return (&TestClient{config: t.config}).QueryTakes(t)
}

// QueryQuestions queries the "questions" edge of the Test entity.
func (t *Test) QueryQuestions() *QuestionQuery {
	return (&TestClient{config: t.config}).QueryQuestions(t)
}

// QueryTranslations queries the "translations" edge of the Test entity.
func (t *Test) QueryTranslations() *TestTranslationQuery {
	return (&TestClient{config: t.config}).QueryTranslations(t)
}

// QueryScales queries the "scales" edge of the Test entity.
func (t *Test) QueryScales() *ScaleQuery {
	return (&TestClient{config: t.config}).QueryScales(t)
}

// Update returns a builder for updating this Test.
// Note that you need to call Test.Unwrap() before calling this method if this Test
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Test) Update() *TestUpdateOne {
	return (&TestClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Test entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Test) Unwrap() *Test {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Test is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Test) String() string {
	var builder strings.Builder
	builder.WriteString("Test(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("create_time=")
	builder.WriteString(t.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(t.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(t.Code)
	builder.WriteString(", ")
	builder.WriteString("published=")
	builder.WriteString(fmt.Sprintf("%v", t.Published))
	builder.WriteByte(')')
	return builder.String()
}

// Tests is a parsable slice of Test.
type Tests []*Test

func (t Tests) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
