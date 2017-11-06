package matchers

import (
	"reflect"

	models "github.com/hootsuite/atlantis/server/events/models"
	"github.com/petergtz/pegomock"
)

func AnyModelsPullrequest() models.PullRequest {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(models.PullRequest))(nil)).Elem()))
	var nullValue models.PullRequest
	return nullValue
}

func EqModelsPullrequest(value models.PullRequest) models.PullRequest {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue models.PullRequest
	return nullValue
}
