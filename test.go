package main

import (
	"reflect"
	"testing"
)

func expect(t *testing.T, a interface{}, b interface{}) {
	if !reflect.DeepEqual(b, a) {
		t.Errorf("Expected %#v (type %v) - Got %#v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}
