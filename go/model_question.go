/*
 * LBA OpenUAV Questions
 *
 * Questions and answers of LBA OpenUAV for preparing to the exam.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Question struct {

	Id int32 `json:"id,omitempty"`

	Text string `json:"text,omitempty"`

	QuestionType string `json:"questionType,omitempty"`

	Structure QuestionStructure `json:"structure,omitempty"`

	Answers []QuestionAnswer `json:"answers,omitempty"`
}

// AssertQuestionRequired checks if the required fields are not zero-ed
func AssertQuestionRequired(obj Question) error {
	if err := AssertQuestionStructureRequired(obj.Structure); err != nil {
		return err
	}
	for _, el := range obj.Answers {
		if err := AssertQuestionAnswerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseQuestionRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Question (e.g. [][]Question), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseQuestionRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aQuestion, ok := obj.(Question)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertQuestionRequired(aQuestion)
	})
}