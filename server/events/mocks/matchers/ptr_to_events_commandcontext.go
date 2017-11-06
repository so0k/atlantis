package matchers

import (
	"reflect"

	events "github.com/hootsuite/atlantis/server/events"
	"github.com/petergtz/pegomock"
)

func AnyPtrToEventsCommandcontext() *events.CommandContext {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*events.CommandContext))(nil)).Elem()))
	var nullValue *events.CommandContext
	return nullValue
}

func EqPtrToEventsCommandcontext(value *events.CommandContext) *events.CommandContext {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *events.CommandContext
	return nullValue
}
