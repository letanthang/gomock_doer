package match

import (
	"github.com/golang/mock/gomock"
	"reflect"
)

type ofType struct {t string}

func (o ofType) Matches(x interface{}) bool {
	return reflect.TypeOf(x).String() == o.t
}

func (o ofType) String() string {
	return "is of type " + o.t
}

func OfType(t string) gomock.Matcher  {
	return &ofType{t}
}
