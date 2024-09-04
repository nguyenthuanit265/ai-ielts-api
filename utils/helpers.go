package utils

import (
	"context"
	"main/component/models"
	"reflect"
	"strings"
)

func GetCurrentUser(ctx context.Context) (models.AuthClaim, models.AIIeltsError) {
	userCtx := ctx.Value(CurrentUser)
	currentUser, ok := userCtx.(models.AuthClaim)
	if !ok {
		return models.AuthClaim{}, models.AIIeltsError{IsError: true, Message: "Cannot parse to type User"}
	}

	return currentUser, models.AIIeltsError{}
}

func ConvertStructToMap(structInput interface{}) map[string]interface{} {
	// Create a new map to hold the result
	resultMap := make(map[string]interface{})

	// Reflect on the input to get its value and type
	val := reflect.ValueOf(structInput)
	if val.Kind() == reflect.Ptr {
		// If a pointer is passed, get the value it points to
		val = val.Elem()
	}
	typ := val.Type()

	// Iterate over all the fields of the struct
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		jsonTag := fieldType.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}
		jsonFieldName := strings.Split(jsonTag, ",")[0]

		// Check if the field is exported before adding it to the map
		if fieldType.PkgPath == "" {
			resultMap[jsonFieldName] = field.Interface()
		}
	}

	return resultMap
}
