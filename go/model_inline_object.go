/*
 * LBA OpenUAV Questions
 *
 * Questions and answers of LBA OpenUAV for preparing to the exam.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type AnswerPayload struct {
	AnswerId int32 `json:"answerId,omitempty"`
	Comment string `json:"comment,omitempty"`
}

// AssertInlineObjectRequired checks if the required fields are not zero-ed
func AssertInlineObjectRequired(obj AnswerPayload) error {
	return nil
}

// AssertRecurseInlineObjectRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of AnswerPayload (e.g. [][]AnswerPayload), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseInlineObjectRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aInlineObject, ok := obj.(AnswerPayload)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertInlineObjectRequired(aInlineObject)
	})
}
