package matchers

import (
	"reflect"

	events "github.com/hootsuite/atlantis/server/events"
	"github.com/petergtz/pegomock"
)

func AnyEventsPreexecuteresult() events.PreExecuteResult {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(events.PreExecuteResult))(nil)).Elem()))
	var nullValue events.PreExecuteResult
	return nullValue
}

func EqEventsPreexecuteresult(value events.PreExecuteResult) events.PreExecuteResult {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue events.PreExecuteResult
	return nullValue
}
