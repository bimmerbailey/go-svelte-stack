package utils

import (
	"fmt"
	"reflect"
)

func getValidStructKeyValues(s struct{}) map[string]string {
	r := reflect.ValueOf(&s).Elem()
	rt := r.Type()
	finalMap := make(map[string]string)

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		rv := reflect.ValueOf(&s)
		value := reflect.Indirect(rv).FieldByName(field.Name)
		fmt.Println(field.Name, value.String())
		finalMap[field.Name] = value.String()
	}
	return finalMap
}
