// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://scap.nist.gov/schema/ocil/2.0
package inter

import (
	"encoding/xml"
)

// Element
type Ocil struct {
	XMLName xml.Name `xml:ocil`

	Generator GeneratorType `xml:"generator"`

	Document *DocumentType `xml:"document"`

	Questionnaires QuestionnairesType `xml:"questionnaires"`

	TestActions TestActionsType `xml:"test_actions"`

	Questions QuestionsType `xml:"questions"`

	Artifacts *ArtifactsType `xml:"artifacts"`

	Variables *VariablesType `xml:"variables"`

	Results *ResultsType `xml:"results"`
}

// Element
type TestAction struct {
	XMLName xml.Name `xml:test_action`

	Notes []string `xml:"notes"`
}

// Element
type QuestionTestAction struct {
	XMLName xml.Name `xml:question_test_action`

	Title *TextType `xml:"title"`

	WhenUnknown *TestActionConditionType `xml:"when_unknown"`

	WhenNotTested *TestActionConditionType `xml:"when_not_tested"`

	WhenNotApplicable *TestActionConditionType `xml:"when_not_applicable"`

	WhenError *TestActionConditionType `xml:"when_error"`
}

// Element
type BooleanQuestionTestAction struct {
	XMLName xml.Name `xml:boolean_question_test_action`

	WhenTrue TestActionConditionType `xml:"when_true"`

	WhenFalse TestActionConditionType `xml:"when_false"`
}

// Element
type ChoiceQuestionTestAction struct {
	XMLName xml.Name `xml:choice_question_test_action`

	WhenChoice []ChoiceTestActionConditionType `xml:"when_choice"`
}

// Element
type NumericQuestionTestAction struct {
	XMLName xml.Name `xml:numeric_question_test_action`
}

// Element
type StringQuestionTestAction struct {
	XMLName xml.Name `xml:string_question_test_action`

	WhenPattern []PatternTestActionConditionType `xml:"when_pattern"`
}

// Element
type Question struct {
	XMLName xml.Name `xml:question`

	QuestionText []QuestionTextType `xml:"question_text"`

	Instructions *InstructionsType `xml:"instructions"`
}

// Element
type BooleanQuestion struct {
	XMLName xml.Name `xml:boolean_question`
}

// Element
type ChoiceQuestion struct {
	XMLName xml.Name `xml:choice_question`
}

// Element
type NumericQuestion struct {
	XMLName xml.Name `xml:numeric_question`
}

// Element
type StringQuestion struct {
	XMLName xml.Name `xml:string_question`
}

// Element
type Variable struct {
	XMLName xml.Name `xml:variable`

	Description *TextType `xml:"description"`
}

// Element
type ConstantVariable struct {
	XMLName xml.Name `xml:constant_variable`

	Value string `xml:"value"`
}

// Element
type LocalVariable struct {
	XMLName xml.Name `xml:local_variable`

	Set string `xml:"set"`
}

// Element
type ExternalVariable struct {
	XMLName xml.Name `xml:external_variable`
}

// Element
type Target struct {
	XMLName xml.Name `xml:target`

	Name string `xml:"name"`
}

// Element
type User struct {
	XMLName xml.Name `xml:user`

	Organization []string `xml:"organization"`

	Position []string `xml:"position"`

	Email []string `xml:"email"`
}

// Element
type System struct {
	XMLName xml.Name `xml:system`

	Organization string `xml:"organization"`

	Ipaddress []string `xml:"ipaddress"`

	Description *TextType `xml:"description"`
}

// Element
type QuestionResult struct {
	XMLName xml.Name `xml:question_result`
}

// Element
type BooleanQuestionResult struct {
	XMLName xml.Name `xml:boolean_question_result`

	Answer bool `xml:"answer"`
}

// Element
type ChoiceQuestionResult struct {
	XMLName xml.Name `xml:choice_question_result`

	Answer ChoiceAnswerType `xml:"answer"`
}

// Element
type NumericQuestionResult struct {
	XMLName xml.Name `xml:numeric_question_result`

	Answer float64 `xml:"answer"`
}

// Element
type StringQuestionResult struct {
	XMLName xml.Name `xml:string_question_result`

	Answer string `xml:"answer"`
}

// Element
type ArtifactValue struct {
	XMLName xml.Name `xml:artifact_value`
}

// Element
type TextArtifactValue struct {
	XMLName xml.Name `xml:text_artifact_value`

	Data string `xml:"data"`
}

// Element
type BinaryArtifactValue struct {
	XMLName xml.Name `xml:binary_artifact_value`

	Data string `xml:"data"`
}

// Element
type ReferenceArtifactValue struct {
	XMLName xml.Name `xml:reference_artifact_value`

	Reference Reference `xml:"reference"`
}

// Element
type Expression struct {
	XMLName xml.Name `xml:expression`

	Value string `xml:"value"`
}

// Element
type WhenPattern struct {
	XMLName xml.Name `xml:when_pattern`
}

// Element
type WhenChoice struct {
	XMLName xml.Name `xml:when_choice`
}

// Element
type WhenRange struct {
	XMLName xml.Name `xml:when_range`
}

// Element
type WhenBoolean struct {
	XMLName xml.Name `xml:when_boolean`
}

// Element
type Reference struct {
	XMLName xml.Name `xml:reference`

	Href string `xml:"href,attr"`
}

// XSD ComplexType declarations

type OCILType struct {
	Generator GeneratorType `xml:"generator"`

	Document *DocumentType `xml:"document"`

	Questionnaires QuestionnairesType `xml:"questionnaires"`

	TestActions TestActionsType `xml:"test_actions"`

	Questions QuestionsType `xml:"questions"`

	Artifacts *ArtifactsType `xml:"artifacts"`

	Variables *VariablesType `xml:"variables"`

	Results *ResultsType `xml:"results"`
}

type QuestionnairesType struct {
	Questionnaire []QuestionnaireType `xml:"questionnaire"`
}

type QuestionnaireType struct {
	Id string `xml:"id,attr"`

	ChildOnly string `xml:"child_only,attr"`
}

type GeneratorType struct {
	ProductName string `xml:"product_name"`

	ProductVersion string `xml:"product_version"`

	Author []UserType `xml:"author"`

	SchemaVersion float64 `xml:"schema_version"`

	Timestamp string `xml:"timestamp"`

	AdditionalData *ExtensionContainerType `xml:"additional_data"`
}

type ExtensionContainerType struct {
}

type DocumentType struct {
	Title string `xml:"title"`

	Description []string `xml:"description"`

	Notice []string `xml:"notice"`
}

type TestActionsType struct {
	TestAction []TestAction `xml:"test_action"`
}

type QuestionTestActionType struct {
	QuestionRef string `xml:"question_ref,attr"`

	Id string `xml:"id,attr"`

	Title *TextType `xml:"title"`

	WhenUnknown *TestActionConditionType `xml:"when_unknown"`

	WhenNotTested *TestActionConditionType `xml:"when_not_tested"`

	WhenNotApplicable *TestActionConditionType `xml:"when_not_applicable"`

	WhenError *TestActionConditionType `xml:"when_error"`
}

type BooleanQuestionTestActionType struct {
	WhenTrue TestActionConditionType `xml:"when_true"`

	WhenFalse TestActionConditionType `xml:"when_false"`
}

type ChoiceQuestionTestActionType struct {
	WhenChoice []ChoiceTestActionConditionType `xml:"when_choice"`
}

type NumericQuestionTestActionType struct {
}

type StringQuestionTestActionType struct {
	WhenPattern []PatternTestActionConditionType `xml:"when_pattern"`
}

type TestActionRefType struct {
	Negate string `xml:"negate,attr"`

	Text string `xml:",chardata"`
}

type ChoiceTestActionConditionType struct {
	ChoiceRef []string `xml:"choice_ref"`
}

type EqualsTestActionConditionType struct {
	VarRef string `xml:"var_ref,attr"`

	Value []float64 `xml:"value"`
}

type RangeTestActionConditionType struct {
	Range []RangeType `xml:"range"`
}

type PatternTestActionConditionType struct {
	Pattern []PatternType `xml:"pattern"`
}

type PatternType struct {
	VarRef string `xml:"var_ref,attr"`

	Text string `xml:",chardata"`
}

type RangeType struct {
	Min *RangeValueType `xml:"min"`

	Max *RangeValueType `xml:"max"`
}

type TestActionConditionType struct {
	ArtifactRefs *ArtifactRefsType `xml:"artifact_refs"`

	Result string `xml:"result"`

	TestActionRef TestActionRefType `xml:"test_action_ref"`
}

type RangeValueType struct {
	Inclusive string `xml:"inclusive,attr"`

	VarRef string `xml:"var_ref,attr"`

	Text string `xml:",chardata"`
}

type QuestionsType struct {
	Question []Question `xml:"question"`

	ChoiceGroup []ChoiceGroupType `xml:"choice_group"`
}

type QuestionTextType struct {
	Sub []SubstitutionTextType `xml:"sub"`

	InnerXml string `xml:",innerxml"`
}

type QuestionType struct {
	Id string `xml:"id,attr"`

	QuestionText []QuestionTextType `xml:"question_text"`

	Instructions *InstructionsType `xml:"instructions"`
}

type BooleanQuestionType struct {
	DefaultAnswer string `xml:"default_answer,attr"`

	Model string `xml:"model,attr"`
}

type ChoiceQuestionType struct {
	DefaultAnswerRef string `xml:"default_answer_ref,attr"`
}

type NumericQuestionType struct {
	DefaultAnswer string `xml:"default_answer,attr"`
}

type StringQuestionType struct {
	DefaultAnswer string `xml:"default_answer,attr"`
}

type ChoiceType struct {
	Id string `xml:"id,attr"`

	VarRef string `xml:"var_ref,attr"`

	Text string `xml:",chardata"`
}

type ChoiceGroupType struct {
	Id string `xml:"id,attr"`

	Choice []ChoiceType `xml:"choice"`
}

type InstructionsType struct {
	Title TextType `xml:"title"`

	Step []StepType `xml:"step"`
}

type ResultsType struct {
	StartTime string `xml:"start_time,attr"`

	EndTime string `xml:"end_time,attr"`

	Title *TextType `xml:"title"`

	QuestionnaireResults *QuestionnaireResultsType `xml:"questionnaire_results"`

	TestActionResults *TestActionResultsType `xml:"test_action_results"`

	QuestionResults *QuestionResultsType `xml:"question_results"`

	ArtifactResults *ArtifactResultsType `xml:"artifact_results"`

	Targets *TargetsType `xml:"targets"`
}

type QuestionnaireResultsType struct {
	QuestionnaireResult []QuestionnaireResultType `xml:"questionnaire_result"`
}

type TestActionResultsType struct {
	TestActionResult []TestActionResultType `xml:"test_action_result"`
}

type QuestionResultsType struct {
	QuestionResult []QuestionResult `xml:"question_result"`
}

type QuestionnaireResultType struct {
	QuestionnaireRef string `xml:"questionnaire_ref,attr"`

	Result string `xml:"result,attr"`

	ArtifactResults *ArtifactResultsType `xml:"artifact_results"`
}

type TestActionResultType struct {
	TestActionRef string `xml:"test_action_ref,attr"`

	Result string `xml:"result,attr"`

	ArtifactResults *ArtifactResultsType `xml:"artifact_results"`
}

type QuestionResultType struct {
	QuestionRef string `xml:"question_ref,attr"`

	Response string `xml:"response,attr"`
}

type BooleanQuestionResultType struct {
	Answer bool `xml:"answer"`
}

type ChoiceQuestionResultType struct {
	Answer ChoiceAnswerType `xml:"answer"`
}

type NumericQuestionResultType struct {
	Answer float64 `xml:"answer"`
}

type StringQuestionResultType struct {
	Answer string `xml:"answer"`
}

type ChoiceAnswerType struct {
	ChoiceRef string `xml:"choice_ref,attr"`
}

type ArtifactsType struct {
	Artifact []ArtifactType `xml:"artifact"`
}

type ArtifactType struct {
	Id string `xml:"id,attr"`

	Persistent string `xml:"persistent,attr"`

	Title TextType `xml:"title"`

	Description TextType `xml:"description"`
}

type ArtifactRefsType struct {
	ArtifactRef []ArtifactRefType `xml:"artifact_ref"`
}

type ArtifactRefType struct {
	Idref string `xml:"idref,attr"`

	Required string `xml:"required,attr"`
}

type ArtifactResultsType struct {
	ArtifactResult []ArtifactResultType `xml:"artifact_result"`
}

type ArtifactValueType struct {
}

type EmbeddedArtifactValueType struct {
	MimeType string `xml:"mime_type,attr"`
}

type TextArtifactValueType struct {
	Data string `xml:"data"`
}

type BinaryArtifactValueType struct {
	Data string `xml:"data"`
}

type ReferenceArtifactValueType struct {
	Reference Reference `xml:"reference"`
}

type ArtifactResultType struct {
	ArtifactRef string `xml:"artifact_ref,attr"`

	Timestamp string `xml:"timestamp,attr"`

	ArtifactValue ArtifactValue `xml:"artifact_value"`

	Provider string `xml:"provider"`

	Submitter UserType `xml:"submitter"`
}

type TargetsType struct {
	Target []Target `xml:"target"`
}

type UserType struct {
	Organization []string `xml:"organization"`

	Position []string `xml:"position"`

	Email []string `xml:"email"`
}

type SystemTargetType struct {
	Organization string `xml:"organization"`

	Ipaddress []string `xml:"ipaddress"`

	Description *TextType `xml:"description"`
}

type VariablesType struct {
	Variable []Variable `xml:"variable"`
}

type VariableType struct {
	Id string `xml:"id,attr"`

	Datatype string `xml:"datatype,attr"`

	Description *TextType `xml:"description"`
}

type ConstantVariableType struct {
	Value string `xml:"value"`
}

type LocalVariableType struct {
	QuestionRef string `xml:"question_ref,attr"`

	Set string `xml:"set"`
}

type ExternalVariableType struct {
}

type SetExpressionBaseType struct {
	Value string `xml:"value"`
}

type SetExpressionPatternType struct {
	Pattern string `xml:"pattern,attr"`
}

type SetExpressionChoiceType struct {
	ChoiceRef string `xml:"choice_ref,attr"`
}

type SetExpressionRangeType struct {
	Min string `xml:"min,attr"`

	Max string `xml:"max,attr"`
}

type SetExpressionBooleanType struct {
	Value string `xml:"value,attr"`
}

type VariableSetType struct {
	Expression []Expression `xml:"expression"`
}

type SubstitutionTextType struct {
	VarRef string `xml:"var_ref,attr"`
}

type ReferenceType struct {
	Href string `xml:"href,attr"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type StepType struct {
	IsDone string `xml:"is_done,attr"`

	IsRequired string `xml:"is_required,attr"`

	Description *TextType `xml:"description"`

	Reference []ReferenceType `xml:"reference"`

	Step []StepType `xml:"step"`
}

type ItemBaseType struct {
	Revision string `xml:"revision,attr"`

	Notes []string `xml:"notes"`
}

type NamedItemBaseType struct {
	Name string `xml:"name"`
}

type CompoundTestActionType struct {
	Title *TextType `xml:"title"`

	Description *TextType `xml:"description"`

	References *ReferencesType `xml:"references"`

	Actions OperationType `xml:"actions"`
}

type ReferencesType struct {
	Reference []ReferenceType `xml:"reference"`
}

type OperationType struct {
	Operation string `xml:"operation,attr"`

	Negate string `xml:"negate,attr"`

	TestActionRef []TestActionRefType `xml:"test_action_ref"`
}

type TextType struct {
	XmlLang string `xml:"lang,attr"`

	Text string `xml:",chardata"`
}
