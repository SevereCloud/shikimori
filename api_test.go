package shikimori_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/SevereCloud/shikimori"
)

func isEmpty(object any) bool {
	// get nil case out of the way
	if object == nil {
		return true
	}

	objValue := reflect.ValueOf(object)

	switch objValue.Kind() { //nolint:exhaustive
	// collection types are empty when they have no element
	case reflect.Chan, reflect.Map, reflect.Slice:
		return objValue.Len() == 0
	// pointers are empty if nil or if the value they point to is empty
	case reflect.Ptr:
		if objValue.IsNil() {
			return true
		}

		deref := objValue.Elem().Interface()

		return isEmpty(deref)
	// for all other types, compare against the zero value
	// array types are empty when they match their zero-initialized state
	default:
		zero := reflect.Zero(objValue.Type())

		return reflect.DeepEqual(object, zero.Interface())
	}
}

func NotEmpty(t *testing.T, obj any) bool {
	t.Helper()

	if isEmpty(obj) {
		t.Errorf("Should NOT be empty, but was %v", obj)

		return false
	}

	return true
}

func NoError(t *testing.T, err error) bool {
	t.Helper()

	if err != nil {
		t.Error(err)

		return false
	}

	return true
}

// func needAccessToken(t *testing.T) {
// 	t.Helper()

// 	if authShiki == nil {
// 		t.Skip("ACCESS_TOKEN empty")
// 	}
// }

var shiki, authShiki *shikimori.API //nolint:gochecknoglobals

func TestMain(m *testing.M) {
	shiki = shikimori.NewAPI()

	if token := os.Getenv("ACCESS_TOKEN"); token != "" {
		authShiki = shikimori.NewAPI()
		authShiki.AccessToken = token
	}

	runTests := m.Run()
	os.Exit(runTests)
}
