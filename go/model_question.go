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

	QuestionId int32 `json:"id,omitempty"`

	QuestionText string `json:"questionText,omitempty"`

	QuestionType string `json:"questionType,omitempty"`

	StructureText string `json:"questionStructure,omitempty"`

	Url string `json:"url,omitempty"`

	Comment string `json:"comment,omitempty"`

	StructureId int32 `json:"structureId,omitempty"`

	Answers []QuestionAnswer `json:"answers,omitempty"`
}

// AssertQuestionRequired checks if the required fields are not zero-ed
func AssertQuestionRequired(obj Question) error {
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
