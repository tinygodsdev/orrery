// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/test"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent/testdisplay"
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
	// AvailableLocales holds the value of the "available_locales" field.
	AvailableLocales []string `json:"available_locales,omitempty"`
	// Mark holds the value of the "mark" field.
	Mark float64 `json:"mark,omitempty"`
	// Duration holds the value of the "duration" field.
	Duration *time.Duration `json:"duration,omitempty"`
	// QuestionCount holds the value of the "question_count" field.
	QuestionCount int `json:"question_count,omitempty"`
	// Image holds the value of the "image" field.
	Image string `json:"image,omitempty"`
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
	// Display holds the value of the display edge.
	Display *TestDisplay `json:"display,omitempty"`
	// Tags holds the value of the tags edge.
	Tags []*Tag `json:"tags,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
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

// DisplayOrErr returns the Display value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TestEdges) DisplayOrErr() (*TestDisplay, error) {
	if e.loadedTypes[4] {
		if e.Display == nil {
			// The edge display was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: testdisplay.Label}
		}
		return e.Display, nil
	}
	return nil, &NotLoadedError{edge: "display"}
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e TestEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[5] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Test) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case test.FieldAvailableLocales:
			values[i] = new([]byte)
		case test.FieldPublished:
			values[i] = new(sql.NullBool)
		case test.FieldMark:
			values[i] = new(sql.NullFloat64)
		case test.FieldDuration, test.FieldQuestionCount:
			values[i] = new(sql.NullInt64)
		case test.FieldCode, test.FieldImage:
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
		case test.FieldAvailableLocales:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field available_locales", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.AvailableLocales); err != nil {
					return fmt.Errorf("unmarshal field available_locales: %w", err)
				}
			}
		case test.FieldMark:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field mark", values[i])
			} else if value.Valid {
				t.Mark = value.Float64
			}
		case test.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				t.Duration = new(time.Duration)
				*t.Duration = time.Duration(value.Int64)
			}
		case test.FieldQuestionCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field question_count", values[i])
			} else if value.Valid {
				t.QuestionCount = int(value.Int64)
			}
		case test.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				t.Image = value.String
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

// QueryDisplay queries the "display" edge of the Test entity.
func (t *Test) QueryDisplay() *TestDisplayQuery {
	return (&TestClient{config: t.config}).QueryDisplay(t)
}

// QueryTags queries the "tags" edge of the Test entity.
func (t *Test) QueryTags() *TagQuery {
	return (&TestClient{config: t.config}).QueryTags(t)
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
	builder.WriteString(", ")
	builder.WriteString("available_locales=")
	builder.WriteString(fmt.Sprintf("%v", t.AvailableLocales))
	builder.WriteString(", ")
	builder.WriteString("mark=")
	builder.WriteString(fmt.Sprintf("%v", t.Mark))
	builder.WriteString(", ")
	if v := t.Duration; v != nil {
		builder.WriteString("duration=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("question_count=")
	builder.WriteString(fmt.Sprintf("%v", t.QuestionCount))
	builder.WriteString(", ")
	builder.WriteString("image=")
	builder.WriteString(t.Image)
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
