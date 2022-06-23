// Code generated by ent, DO NOT EDIT.

package take

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the take type in the database.
	Label = "take"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldSeed holds the string denoting the seed field in the database.
	FieldSeed = "seed"
	// FieldProgress holds the string denoting the progress field in the database.
	FieldProgress = "progress"
	// FieldPage holds the string denoting the page field in the database.
	FieldPage = "page"
	// FieldStartTime holds the string denoting the start_time field in the database.
	FieldStartTime = "start_time"
	// FieldEndTime holds the string denoting the end_time field in the database.
	FieldEndTime = "end_time"
	// FieldSuspicious holds the string denoting the suspicious field in the database.
	FieldSuspicious = "suspicious"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldMeta holds the string denoting the meta field in the database.
	FieldMeta = "meta"
	// EdgeResponses holds the string denoting the responses edge name in mutations.
	EdgeResponses = "responses"
	// EdgeResults holds the string denoting the results edge name in mutations.
	EdgeResults = "results"
	// EdgeTest holds the string denoting the test edge name in mutations.
	EdgeTest = "test"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the take in the database.
	Table = "takes"
	// ResponsesTable is the table that holds the responses relation/edge.
	ResponsesTable = "responses"
	// ResponsesInverseTable is the table name for the Response entity.
	// It exists in this package in order to avoid circular dependency with the "response" package.
	ResponsesInverseTable = "responses"
	// ResponsesColumn is the table column denoting the responses relation/edge.
	ResponsesColumn = "take_responses"
	// ResultsTable is the table that holds the results relation/edge.
	ResultsTable = "results"
	// ResultsInverseTable is the table name for the Result entity.
	// It exists in this package in order to avoid circular dependency with the "result" package.
	ResultsInverseTable = "results"
	// ResultsColumn is the table column denoting the results relation/edge.
	ResultsColumn = "take_results"
	// TestTable is the table that holds the test relation/edge.
	TestTable = "takes"
	// TestInverseTable is the table name for the Test entity.
	// It exists in this package in order to avoid circular dependency with the "test" package.
	TestInverseTable = "tests"
	// TestColumn is the table column denoting the test relation/edge.
	TestColumn = "test_takes"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "takes"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_takes"
)

// Columns holds all SQL columns for take fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldSeed,
	FieldProgress,
	FieldPage,
	FieldStartTime,
	FieldEndTime,
	FieldSuspicious,
	FieldStatus,
	FieldMeta,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "takes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"test_takes",
	"user_takes",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
	// DefaultSeed holds the default value on creation for the "seed" field.
	DefaultSeed int64
	// DefaultProgress holds the default value on creation for the "progress" field.
	DefaultProgress int
	// DefaultPage holds the default value on creation for the "page" field.
	DefaultPage int
	// DefaultSuspicious holds the default value on creation for the "suspicious" field.
	DefaultSuspicious bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// StatusIntro is the default value of the Status enum.
const DefaultStatus = StatusIntro

// Status values.
const (
	StatusIntro     Status = "intro"
	StatusQuestions Status = "questions"
	StatusFinish    Status = "finish"
	StatusResult    Status = "result"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusIntro, StatusQuestions, StatusFinish, StatusResult:
		return nil
	default:
		return fmt.Errorf("take: invalid enum value for status field: %q", s)
	}
}
