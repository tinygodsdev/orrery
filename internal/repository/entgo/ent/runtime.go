// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/interpretation"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/interpretationtranslation"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/item"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/itemtranslation"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/norm"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/question"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/questiontranslation"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/response"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/sample"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/scale"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/scaleitem"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/scaletranslation"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/schema"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/take"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/test"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/testdisplay"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/testtranslation"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/user"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/usersession"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	interpretationMixin := schema.Interpretation{}.Mixin()
	interpretationMixinFields0 := interpretationMixin[0].Fields()
	_ = interpretationMixinFields0
	interpretationFields := schema.Interpretation{}.Fields()
	_ = interpretationFields
	// interpretationDescCreateTime is the schema descriptor for create_time field.
	interpretationDescCreateTime := interpretationMixinFields0[0].Descriptor()
	// interpretation.DefaultCreateTime holds the default value on creation for the create_time field.
	interpretation.DefaultCreateTime = interpretationDescCreateTime.Default.(func() time.Time)
	// interpretationDescUpdateTime is the schema descriptor for update_time field.
	interpretationDescUpdateTime := interpretationMixinFields0[1].Descriptor()
	// interpretation.DefaultUpdateTime holds the default value on creation for the update_time field.
	interpretation.DefaultUpdateTime = interpretationDescUpdateTime.Default.(func() time.Time)
	// interpretation.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	interpretation.UpdateDefaultUpdateTime = interpretationDescUpdateTime.UpdateDefault.(func() time.Time)
	// interpretationDescID is the schema descriptor for id field.
	interpretationDescID := interpretationFields[0].Descriptor()
	// interpretation.DefaultID holds the default value on creation for the id field.
	interpretation.DefaultID = interpretationDescID.Default.(func() uuid.UUID)
	interpretationtranslationFields := schema.InterpretationTranslation{}.Fields()
	_ = interpretationtranslationFields
	// interpretationtranslationDescContent is the schema descriptor for content field.
	interpretationtranslationDescContent := interpretationtranslationFields[1].Descriptor()
	// interpretationtranslation.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	interpretationtranslation.ContentValidator = interpretationtranslationDescContent.Validators[0].(func(string) error)
	// interpretationtranslationDescID is the schema descriptor for id field.
	interpretationtranslationDescID := interpretationtranslationFields[0].Descriptor()
	// interpretationtranslation.DefaultID holds the default value on creation for the id field.
	interpretationtranslation.DefaultID = interpretationtranslationDescID.Default.(func() uuid.UUID)
	itemMixin := schema.Item{}.Mixin()
	itemMixinFields0 := itemMixin[0].Fields()
	_ = itemMixinFields0
	itemFields := schema.Item{}.Fields()
	_ = itemFields
	// itemDescCreateTime is the schema descriptor for create_time field.
	itemDescCreateTime := itemMixinFields0[0].Descriptor()
	// item.DefaultCreateTime holds the default value on creation for the create_time field.
	item.DefaultCreateTime = itemDescCreateTime.Default.(func() time.Time)
	// itemDescUpdateTime is the schema descriptor for update_time field.
	itemDescUpdateTime := itemMixinFields0[1].Descriptor()
	// item.DefaultUpdateTime holds the default value on creation for the update_time field.
	item.DefaultUpdateTime = itemDescUpdateTime.Default.(func() time.Time)
	// item.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	item.UpdateDefaultUpdateTime = itemDescUpdateTime.UpdateDefault.(func() time.Time)
	// itemDescCode is the schema descriptor for code field.
	itemDescCode := itemFields[1].Descriptor()
	// item.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	item.CodeValidator = itemDescCode.Validators[0].(func(string) error)
	// itemDescSteps is the schema descriptor for steps field.
	itemDescSteps := itemFields[2].Descriptor()
	// item.DefaultSteps holds the default value on creation for the steps field.
	item.DefaultSteps = itemDescSteps.Default.(int)
	// itemDescID is the schema descriptor for id field.
	itemDescID := itemFields[0].Descriptor()
	// item.DefaultID holds the default value on creation for the id field.
	item.DefaultID = itemDescID.Default.(func() uuid.UUID)
	itemtranslationFields := schema.ItemTranslation{}.Fields()
	_ = itemtranslationFields
	// itemtranslationDescContent is the schema descriptor for content field.
	itemtranslationDescContent := itemtranslationFields[1].Descriptor()
	// itemtranslation.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	itemtranslation.ContentValidator = itemtranslationDescContent.Validators[0].(func(string) error)
	// itemtranslationDescID is the schema descriptor for id field.
	itemtranslationDescID := itemtranslationFields[0].Descriptor()
	// itemtranslation.DefaultID holds the default value on creation for the id field.
	itemtranslation.DefaultID = itemtranslationDescID.Default.(func() uuid.UUID)
	normMixin := schema.Norm{}.Mixin()
	normMixinFields0 := normMixin[0].Fields()
	_ = normMixinFields0
	normFields := schema.Norm{}.Fields()
	_ = normFields
	// normDescCreateTime is the schema descriptor for create_time field.
	normDescCreateTime := normMixinFields0[0].Descriptor()
	// norm.DefaultCreateTime holds the default value on creation for the create_time field.
	norm.DefaultCreateTime = normDescCreateTime.Default.(func() time.Time)
	// normDescUpdateTime is the schema descriptor for update_time field.
	normDescUpdateTime := normMixinFields0[1].Descriptor()
	// norm.DefaultUpdateTime holds the default value on creation for the update_time field.
	norm.DefaultUpdateTime = normDescUpdateTime.Default.(func() time.Time)
	// norm.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	norm.UpdateDefaultUpdateTime = normDescUpdateTime.UpdateDefault.(func() time.Time)
	// normDescName is the schema descriptor for name field.
	normDescName := normFields[1].Descriptor()
	// norm.NameValidator is a validator for the "name" field. It is called by the builders before save.
	norm.NameValidator = normDescName.Validators[0].(func(string) error)
	// normDescBase is the schema descriptor for base field.
	normDescBase := normFields[2].Descriptor()
	// norm.DefaultBase holds the default value on creation for the base field.
	norm.DefaultBase = normDescBase.Default.(int)
	// norm.BaseValidator is a validator for the "base" field. It is called by the builders before save.
	norm.BaseValidator = normDescBase.Validators[0].(func(int) error)
	// normDescID is the schema descriptor for id field.
	normDescID := normFields[0].Descriptor()
	// norm.DefaultID holds the default value on creation for the id field.
	norm.DefaultID = normDescID.Default.(func() uuid.UUID)
	questionMixin := schema.Question{}.Mixin()
	questionMixinFields0 := questionMixin[0].Fields()
	_ = questionMixinFields0
	questionFields := schema.Question{}.Fields()
	_ = questionFields
	// questionDescCreateTime is the schema descriptor for create_time field.
	questionDescCreateTime := questionMixinFields0[0].Descriptor()
	// question.DefaultCreateTime holds the default value on creation for the create_time field.
	question.DefaultCreateTime = questionDescCreateTime.Default.(func() time.Time)
	// questionDescUpdateTime is the schema descriptor for update_time field.
	questionDescUpdateTime := questionMixinFields0[1].Descriptor()
	// question.DefaultUpdateTime holds the default value on creation for the update_time field.
	question.DefaultUpdateTime = questionDescUpdateTime.Default.(func() time.Time)
	// question.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	question.UpdateDefaultUpdateTime = questionDescUpdateTime.UpdateDefault.(func() time.Time)
	// questionDescOrder is the schema descriptor for order field.
	questionDescOrder := questionFields[1].Descriptor()
	// question.DefaultOrder holds the default value on creation for the order field.
	question.DefaultOrder = questionDescOrder.Default.(int)
	// questionDescCode is the schema descriptor for code field.
	questionDescCode := questionFields[2].Descriptor()
	// question.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	question.CodeValidator = questionDescCode.Validators[0].(func(string) error)
	// questionDescID is the schema descriptor for id field.
	questionDescID := questionFields[0].Descriptor()
	// question.DefaultID holds the default value on creation for the id field.
	question.DefaultID = questionDescID.Default.(func() uuid.UUID)
	questiontranslationFields := schema.QuestionTranslation{}.Fields()
	_ = questiontranslationFields
	// questiontranslationDescID is the schema descriptor for id field.
	questiontranslationDescID := questiontranslationFields[0].Descriptor()
	// questiontranslation.DefaultID holds the default value on creation for the id field.
	questiontranslation.DefaultID = questiontranslationDescID.Default.(func() uuid.UUID)
	responseMixin := schema.Response{}.Mixin()
	responseMixinFields0 := responseMixin[0].Fields()
	_ = responseMixinFields0
	responseFields := schema.Response{}.Fields()
	_ = responseFields
	// responseDescCreateTime is the schema descriptor for create_time field.
	responseDescCreateTime := responseMixinFields0[0].Descriptor()
	// response.DefaultCreateTime holds the default value on creation for the create_time field.
	response.DefaultCreateTime = responseDescCreateTime.Default.(func() time.Time)
	// responseDescUpdateTime is the schema descriptor for update_time field.
	responseDescUpdateTime := responseMixinFields0[1].Descriptor()
	// response.DefaultUpdateTime holds the default value on creation for the update_time field.
	response.DefaultUpdateTime = responseDescUpdateTime.Default.(func() time.Time)
	// response.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	response.UpdateDefaultUpdateTime = responseDescUpdateTime.UpdateDefault.(func() time.Time)
	// responseDescValue is the schema descriptor for value field.
	responseDescValue := responseFields[1].Descriptor()
	// response.DefaultValue holds the default value on creation for the value field.
	response.DefaultValue = responseDescValue.Default.(int)
	// responseDescID is the schema descriptor for id field.
	responseDescID := responseFields[0].Descriptor()
	// response.DefaultID holds the default value on creation for the id field.
	response.DefaultID = responseDescID.Default.(func() uuid.UUID)
	sampleMixin := schema.Sample{}.Mixin()
	sampleMixinFields0 := sampleMixin[0].Fields()
	_ = sampleMixinFields0
	sampleFields := schema.Sample{}.Fields()
	_ = sampleFields
	// sampleDescCreateTime is the schema descriptor for create_time field.
	sampleDescCreateTime := sampleMixinFields0[0].Descriptor()
	// sample.DefaultCreateTime holds the default value on creation for the create_time field.
	sample.DefaultCreateTime = sampleDescCreateTime.Default.(func() time.Time)
	// sampleDescUpdateTime is the schema descriptor for update_time field.
	sampleDescUpdateTime := sampleMixinFields0[1].Descriptor()
	// sample.DefaultUpdateTime holds the default value on creation for the update_time field.
	sample.DefaultUpdateTime = sampleDescUpdateTime.Default.(func() time.Time)
	// sample.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	sample.UpdateDefaultUpdateTime = sampleDescUpdateTime.UpdateDefault.(func() time.Time)
	// sampleDescCode is the schema descriptor for code field.
	sampleDescCode := sampleFields[1].Descriptor()
	// sample.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	sample.CodeValidator = sampleDescCode.Validators[0].(func(string) error)
	// sampleDescID is the schema descriptor for id field.
	sampleDescID := sampleFields[0].Descriptor()
	// sample.DefaultID holds the default value on creation for the id field.
	sample.DefaultID = sampleDescID.Default.(func() uuid.UUID)
	scaleMixin := schema.Scale{}.Mixin()
	scaleMixinFields0 := scaleMixin[0].Fields()
	_ = scaleMixinFields0
	scaleFields := schema.Scale{}.Fields()
	_ = scaleFields
	// scaleDescCreateTime is the schema descriptor for create_time field.
	scaleDescCreateTime := scaleMixinFields0[0].Descriptor()
	// scale.DefaultCreateTime holds the default value on creation for the create_time field.
	scale.DefaultCreateTime = scaleDescCreateTime.Default.(func() time.Time)
	// scaleDescUpdateTime is the schema descriptor for update_time field.
	scaleDescUpdateTime := scaleMixinFields0[1].Descriptor()
	// scale.DefaultUpdateTime holds the default value on creation for the update_time field.
	scale.DefaultUpdateTime = scaleDescUpdateTime.Default.(func() time.Time)
	// scale.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	scale.UpdateDefaultUpdateTime = scaleDescUpdateTime.UpdateDefault.(func() time.Time)
	// scaleDescCode is the schema descriptor for code field.
	scaleDescCode := scaleFields[1].Descriptor()
	// scale.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	scale.CodeValidator = scaleDescCode.Validators[0].(func(string) error)
	// scaleDescGlobal is the schema descriptor for global field.
	scaleDescGlobal := scaleFields[2].Descriptor()
	// scale.DefaultGlobal holds the default value on creation for the global field.
	scale.DefaultGlobal = scaleDescGlobal.Default.(bool)
	// scaleDescID is the schema descriptor for id field.
	scaleDescID := scaleFields[0].Descriptor()
	// scale.DefaultID holds the default value on creation for the id field.
	scale.DefaultID = scaleDescID.Default.(func() uuid.UUID)
	scaleitemFields := schema.ScaleItem{}.Fields()
	_ = scaleitemFields
	// scaleitemDescReverse is the schema descriptor for reverse field.
	scaleitemDescReverse := scaleitemFields[0].Descriptor()
	// scaleitem.DefaultReverse holds the default value on creation for the reverse field.
	scaleitem.DefaultReverse = scaleitemDescReverse.Default.(bool)
	scaletranslationFields := schema.ScaleTranslation{}.Fields()
	_ = scaletranslationFields
	// scaletranslationDescTitle is the schema descriptor for title field.
	scaletranslationDescTitle := scaletranslationFields[1].Descriptor()
	// scaletranslation.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	scaletranslation.TitleValidator = scaletranslationDescTitle.Validators[0].(func(string) error)
	// scaletranslationDescID is the schema descriptor for id field.
	scaletranslationDescID := scaletranslationFields[0].Descriptor()
	// scaletranslation.DefaultID holds the default value on creation for the id field.
	scaletranslation.DefaultID = scaletranslationDescID.Default.(func() uuid.UUID)
	takeMixin := schema.Take{}.Mixin()
	takeMixinFields0 := takeMixin[0].Fields()
	_ = takeMixinFields0
	takeFields := schema.Take{}.Fields()
	_ = takeFields
	// takeDescCreateTime is the schema descriptor for create_time field.
	takeDescCreateTime := takeMixinFields0[0].Descriptor()
	// take.DefaultCreateTime holds the default value on creation for the create_time field.
	take.DefaultCreateTime = takeDescCreateTime.Default.(func() time.Time)
	// takeDescUpdateTime is the schema descriptor for update_time field.
	takeDescUpdateTime := takeMixinFields0[1].Descriptor()
	// take.DefaultUpdateTime holds the default value on creation for the update_time field.
	take.DefaultUpdateTime = takeDescUpdateTime.Default.(func() time.Time)
	// take.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	take.UpdateDefaultUpdateTime = takeDescUpdateTime.UpdateDefault.(func() time.Time)
	// takeDescSeed is the schema descriptor for seed field.
	takeDescSeed := takeFields[1].Descriptor()
	// take.DefaultSeed holds the default value on creation for the seed field.
	take.DefaultSeed = takeDescSeed.Default.(int64)
	// takeDescProgress is the schema descriptor for progress field.
	takeDescProgress := takeFields[2].Descriptor()
	// take.DefaultProgress holds the default value on creation for the progress field.
	take.DefaultProgress = takeDescProgress.Default.(int)
	// takeDescPage is the schema descriptor for page field.
	takeDescPage := takeFields[3].Descriptor()
	// take.DefaultPage holds the default value on creation for the page field.
	take.DefaultPage = takeDescPage.Default.(int)
	// takeDescID is the schema descriptor for id field.
	takeDescID := takeFields[0].Descriptor()
	// take.DefaultID holds the default value on creation for the id field.
	take.DefaultID = takeDescID.Default.(func() uuid.UUID)
	testMixin := schema.Test{}.Mixin()
	testMixinFields0 := testMixin[0].Fields()
	_ = testMixinFields0
	testFields := schema.Test{}.Fields()
	_ = testFields
	// testDescCreateTime is the schema descriptor for create_time field.
	testDescCreateTime := testMixinFields0[0].Descriptor()
	// test.DefaultCreateTime holds the default value on creation for the create_time field.
	test.DefaultCreateTime = testDescCreateTime.Default.(func() time.Time)
	// testDescUpdateTime is the schema descriptor for update_time field.
	testDescUpdateTime := testMixinFields0[1].Descriptor()
	// test.DefaultUpdateTime holds the default value on creation for the update_time field.
	test.DefaultUpdateTime = testDescUpdateTime.Default.(func() time.Time)
	// test.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	test.UpdateDefaultUpdateTime = testDescUpdateTime.UpdateDefault.(func() time.Time)
	// testDescCode is the schema descriptor for code field.
	testDescCode := testFields[1].Descriptor()
	// test.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	test.CodeValidator = func() func(string) error {
		validators := testDescCode.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(code string) error {
			for _, fn := range fns {
				if err := fn(code); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// testDescPublished is the schema descriptor for published field.
	testDescPublished := testFields[2].Descriptor()
	// test.DefaultPublished holds the default value on creation for the published field.
	test.DefaultPublished = testDescPublished.Default.(bool)
	// testDescID is the schema descriptor for id field.
	testDescID := testFields[0].Descriptor()
	// test.DefaultID holds the default value on creation for the id field.
	test.DefaultID = testDescID.Default.(func() uuid.UUID)
	testdisplayFields := schema.TestDisplay{}.Fields()
	_ = testdisplayFields
	// testdisplayDescRandomizeOrder is the schema descriptor for randomize_order field.
	testdisplayDescRandomizeOrder := testdisplayFields[1].Descriptor()
	// testdisplay.DefaultRandomizeOrder holds the default value on creation for the randomize_order field.
	testdisplay.DefaultRandomizeOrder = testdisplayDescRandomizeOrder.Default.(bool)
	// testdisplayDescQuestionsPerPage is the schema descriptor for questions_per_page field.
	testdisplayDescQuestionsPerPage := testdisplayFields[2].Descriptor()
	// testdisplay.DefaultQuestionsPerPage holds the default value on creation for the questions_per_page field.
	testdisplay.DefaultQuestionsPerPage = testdisplayDescQuestionsPerPage.Default.(int)
	// testdisplay.QuestionsPerPageValidator is a validator for the "questions_per_page" field. It is called by the builders before save.
	testdisplay.QuestionsPerPageValidator = testdisplayDescQuestionsPerPage.Validators[0].(func(int) error)
	// testdisplayDescID is the schema descriptor for id field.
	testdisplayDescID := testdisplayFields[0].Descriptor()
	// testdisplay.DefaultID holds the default value on creation for the id field.
	testdisplay.DefaultID = testdisplayDescID.Default.(func() uuid.UUID)
	testtranslationFields := schema.TestTranslation{}.Fields()
	_ = testtranslationFields
	// testtranslationDescTitle is the schema descriptor for title field.
	testtranslationDescTitle := testtranslationFields[1].Descriptor()
	// testtranslation.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	testtranslation.TitleValidator = func() func(string) error {
		validators := testtranslationDescTitle.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(title string) error {
			for _, fn := range fns {
				if err := fn(title); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// testtranslationDescID is the schema descriptor for id field.
	testtranslationDescID := testtranslationFields[0].Descriptor()
	// testtranslation.DefaultID holds the default value on creation for the id field.
	testtranslation.DefaultID = testtranslationDescID.Default.(func() uuid.UUID)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPicture is the schema descriptor for picture field.
	userDescPicture := userFields[3].Descriptor()
	// user.DefaultPicture holds the default value on creation for the picture field.
	user.DefaultPicture = userDescPicture.Default.(string)
	// userDescAdmin is the schema descriptor for admin field.
	userDescAdmin := userFields[5].Descriptor()
	// user.DefaultAdmin holds the default value on creation for the admin field.
	user.DefaultAdmin = userDescAdmin.Default.(bool)
	// userDescAnonymous is the schema descriptor for anonymous field.
	userDescAnonymous := userFields[6].Descriptor()
	// user.DefaultAnonymous holds the default value on creation for the anonymous field.
	user.DefaultAnonymous = userDescAnonymous.Default.(bool)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
	usersessionMixin := schema.UserSession{}.Mixin()
	usersessionMixinFields0 := usersessionMixin[0].Fields()
	_ = usersessionMixinFields0
	usersessionFields := schema.UserSession{}.Fields()
	_ = usersessionFields
	// usersessionDescCreateTime is the schema descriptor for create_time field.
	usersessionDescCreateTime := usersessionMixinFields0[0].Descriptor()
	// usersession.DefaultCreateTime holds the default value on creation for the create_time field.
	usersession.DefaultCreateTime = usersessionDescCreateTime.Default.(func() time.Time)
	// usersessionDescUpdateTime is the schema descriptor for update_time field.
	usersessionDescUpdateTime := usersessionMixinFields0[1].Descriptor()
	// usersession.DefaultUpdateTime holds the default value on creation for the update_time field.
	usersession.DefaultUpdateTime = usersessionDescUpdateTime.Default.(func() time.Time)
	// usersession.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	usersession.UpdateDefaultUpdateTime = usersessionDescUpdateTime.UpdateDefault.(func() time.Time)
	// usersessionDescSid is the schema descriptor for sid field.
	usersessionDescSid := usersessionFields[0].Descriptor()
	// usersession.SidValidator is a validator for the "sid" field. It is called by the builders before save.
	usersession.SidValidator = usersessionDescSid.Validators[0].(func(string) error)
	// usersessionDescLastActivity is the schema descriptor for last_activity field.
	usersessionDescLastActivity := usersessionFields[3].Descriptor()
	// usersession.DefaultLastActivity holds the default value on creation for the last_activity field.
	usersession.DefaultLastActivity = usersessionDescLastActivity.Default.(func() time.Time)
	// usersessionDescActive is the schema descriptor for active field.
	usersessionDescActive := usersessionFields[4].Descriptor()
	// usersession.DefaultActive holds the default value on creation for the active field.
	usersession.DefaultActive = usersessionDescActive.Default.(bool)
}
