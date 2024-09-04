package utils

import (
	"gopkg.in/guregu/null.v4"
	"reflect"
	"time"
)

func DecodeNullTimeHook(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	if f == reflect.TypeOf(null.Time{}) && t == reflect.TypeOf(time.Time{}) {
		t, ok := data.(null.Time)
		if ok && t.Valid {
			return t.Time, nil
		}
		return time.Time{}, nil
	}
	return data, nil
}
