/*
 * LBA OpenUAV Questions
 *
 * Questions and answers of LBA OpenUAV for preparing to the exam.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
)



// DefaultApiRouter defines the required methods for binding the api requests to a responses for the DefaultApi
// The DefaultApiRouter implementation should parse necessary information from the http request,
// pass the data to a DefaultApiServicer to perform the required actions, then write the service results to the http response.
type DefaultApiRouter interface {
	GetAllQuestions(http.ResponseWriter, *http.Request)
	GetOneQuestion(http.ResponseWriter, *http.Request)
	GetQuestionById(http.ResponseWriter, *http.Request)
	SetQuestionAnswered(http.ResponseWriter, *http.Request)
}


// DefaultApiServicer defines the api actions for the DefaultApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DefaultApiServicer interface {
	GetAllQuestions(context.Context) (ImplResponse, error)
	GetOneQuestion(context.Context, string) (ImplResponse, error)
	GetQuestionById(context.Context, int32) (ImplResponse, error)
	SetQuestionAnswered(context.Context, int32, AnswerPayload) (ImplResponse, error)
}
