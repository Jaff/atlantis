// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"github.com/petergtz/pegomock"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
)

func AnyModelsProjectResult() models.ProjectResult {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(models.ProjectResult))(nil)).Elem()))
	var nullValue models.ProjectResult
	return nullValue
}

func EqModelsProjectResult(value models.ProjectResult) models.ProjectResult {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue models.ProjectResult
	return nullValue
}
