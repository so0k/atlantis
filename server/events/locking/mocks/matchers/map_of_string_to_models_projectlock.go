package matchers

import (
	"reflect"

	"github.com/petergtz/pegomock"

	models "github.com/hootsuite/atlantis/server/events/models"
)

func AnyMapOfStringToModelsProjectlock() map[string]models.ProjectLock {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(map[string]models.ProjectLock))(nil)).Elem()))
	var nullValue map[string]models.ProjectLock
	return nullValue
}

func EqMapOfStringToModelsProjectlock(value map[string]models.ProjectLock) map[string]models.ProjectLock {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue map[string]models.ProjectLock
	return nullValue
}
