package matchers

import (
	"reflect"

	events "github.com/hootsuite/atlantis/server/events"
	"github.com/petergtz/pegomock"
)

func AnyEventsProjectconfig() events.ProjectConfig {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(events.ProjectConfig))(nil)).Elem()))
	var nullValue events.ProjectConfig
	return nullValue
}

func EqEventsProjectconfig(value events.ProjectConfig) events.ProjectConfig {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue events.ProjectConfig
	return nullValue
}
