// Code generated by ent, DO NOT EDIT.

package test

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the test type in the database.
	Label = "test"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldPublished holds the string denoting the published field in the database.
	FieldPublished = "published"
	// FieldAvailableLocales holds the string denoting the available_locales field in the database.
	FieldAvailableLocales = "available_locales"
	// EdgeTakes holds the string denoting the takes edge name in mutations.
	EdgeTakes = "takes"
	// EdgeQuestions holds the string denoting the questions edge name in mutations.
	EdgeQuestions = "questions"
	// EdgeTranslations holds the string denoting the translations edge name in mutations.
	EdgeTranslations = "translations"
	// EdgeScales holds the string denoting the scales edge name in mutations.
	EdgeScales = "scales"
	// EdgeDisplay holds the string denoting the display edge name in mutations.
	EdgeDisplay = "display"
	// Table holds the table name of the test in the database.
	Table = "tests"
	// TakesTable is the table that holds the takes relation/edge.
	TakesTable = "takes"
	// TakesInverseTable is the table name for the Take entity.
	// It exists in this package in order to avoid circular dependency with the "take" package.
	TakesInverseTable = "takes"
	// TakesColumn is the table column denoting the takes relation/edge.
	TakesColumn = "test_takes"
	// QuestionsTable is the table that holds the questions relation/edge. The primary key declared below.
	QuestionsTable = "test_questions"
	// QuestionsInverseTable is the table name for the Question entity.
	// It exists in this package in order to avoid circular dependency with the "question" package.
	QuestionsInverseTable = "questions"
	// TranslationsTable is the table that holds the translations relation/edge.
	TranslationsTable = "test_translations"
	// TranslationsInverseTable is the table name for the TestTranslation entity.
	// It exists in this package in order to avoid circular dependency with the "testtranslation" package.
	TranslationsInverseTable = "test_translations"
	// TranslationsColumn is the table column denoting the translations relation/edge.
	TranslationsColumn = "test_translations"
	// ScalesTable is the table that holds the scales relation/edge. The primary key declared below.
	ScalesTable = "test_scales"
	// ScalesInverseTable is the table name for the Scale entity.
	// It exists in this package in order to avoid circular dependency with the "scale" package.
	ScalesInverseTable = "scales"
	// DisplayTable is the table that holds the display relation/edge.
	DisplayTable = "test_displays"
	// DisplayInverseTable is the table name for the TestDisplay entity.
	// It exists in this package in order to avoid circular dependency with the "testdisplay" package.
	DisplayInverseTable = "test_displays"
	// DisplayColumn is the table column denoting the display relation/edge.
	DisplayColumn = "test_display"
)

// Columns holds all SQL columns for test fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldCode,
	FieldPublished,
	FieldAvailableLocales,
}

var (
	// QuestionsPrimaryKey and QuestionsColumn2 are the table columns denoting the
	// primary key for the questions relation (M2M).
	QuestionsPrimaryKey = []string{"test_id", "question_id"}
	// ScalesPrimaryKey and ScalesColumn2 are the table columns denoting the
	// primary key for the scales relation (M2M).
	ScalesPrimaryKey = []string{"test_id", "scale_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// DefaultPublished holds the default value on creation for the "published" field.
	DefaultPublished bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
