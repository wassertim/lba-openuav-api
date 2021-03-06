/*
 * LBA OpenUAV Questions
 *
 * Questions and answers of LBA OpenUAV for preparing to the exam.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type QuestionStructure struct {

	Id int32 `json:"id,omitempty"`

	Text string `json:"text,omitempty"`
}

// AssertQuestionStructureRequired checks if the required fields are not zero-ed
func AssertQuestionStructureRequired(obj QuestionStructure) error {
	return nil
}

// AssertRecurseQuestionStructureRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of QuestionStructure (e.g. [][]QuestionStructure), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseQuestionStructureRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aQuestionStructure, ok := obj.(QuestionStructure)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertQuestionStructureRequired(aQuestionStructure)
	})
}
