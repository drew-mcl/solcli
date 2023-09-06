package printer

import (
	"fmt"
	"reflect"

	"github.com/fatih/color"
)

type PrintOptions struct {
	DisplayFields []string
	OmitEmpty     bool
	DisplayDepth  int
	OmitFields    []string
	PrintField    string
	AllFields     bool
}

func Print(data interface{}, opts PrintOptions, indent int) {
	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)

	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = v.Type()
	}

	switch v.Kind() {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fieldType := t.Field(i)
			fieldName := fieldType.Name

			// Skip fields based on options
			if !opts.AllFields && opts.PrintField != "" && fieldName != opts.PrintField {
				continue
			}

			// Indentation for visual clarity
			for j := 0; j < indent; j++ {
				fmt.Print("  ")
			}

			fmt.Printf("%s: ", color.CyanString(fieldName))

			if field.Kind() == reflect.Slice {
				fmt.Println()
				for idx := 0; idx < field.Len(); idx++ {
					Print(field.Index(idx).Interface(), opts, indent+1)
				}
			} else if field.Kind() == reflect.Struct || (field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct) {
				fmt.Println()
				Print(field.Interface(), opts, indent+1)
			} else {
				fmt.Println(field.Interface())
			}
		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			Print(v.Index(i).Interface(), opts, indent)
		}
	default:
		fmt.Println(data)
	}
}
