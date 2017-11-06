package matchers

import (
	"reflect"

	events "github.com/hootsuite/atlantis/server/events"
	"github.com/petergtz/pegomock"
)

func AnyEventsStatus() events.Status {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(events.Status))(nil)).Elem()))
	var nullValue events.Status
	return nullValue
}

func EqEventsStatus(value events.Status) events.Status {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue events.Status
	return nullValue
}
