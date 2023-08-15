package printer

import (
	"fmt"
	"reflect"

	"github.com/fatih/color"
	"golang.org/x/exp/slog"
)

type PrintOptions struct {
	DisplayFields []string
	OmitEmpty     bool
	DisplayDepth  int
	OmitFields    []string
	PrintField    string
	AllFields     bool
}

func Print(data interface{}, opts PrintOptions) {
	slog.Debug("Attempting to print data", slog.Any("data", data))

	v := reflect.ValueOf(data)

	if v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		fmt.Println("Print() received a non-struct type.")
		return
	}

	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		if !opts.AllFields && opts.PrintField != "" && field.Name != opts.PrintField {
			continue
		}

		color.Cyan("%s: ", field.Name)
		printRecursive(value, opts, 0)
	}
}

func printRecursive(v reflect.Value, opts PrintOptions, depth int) {
	if !opts.AllFields && depth > opts.DisplayDepth {
		return
	}

	switch v.Kind() {
	case reflect.Ptr:
		printRecursive(v.Elem(), opts, depth+1)
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i)
			value := v.Field(i)

			if opts.OmitEmpty && isEmptyValue(value) {
				continue
			}

			omit := false
			for _, omitField := range opts.OmitFields {
				if field.Name == omitField {
					omit = true
					break
				}
			}

			if omit {
				continue
			}

			if len(opts.DisplayFields) > 0 {
				display := false
				for _, displayField := range opts.DisplayFields {
					if field.Name == displayField {
						display = true
						break
					}
				}

				if !display {
					continue
				}
			}

			color.Cyan("%s: ", field.Name)
			printRecursive(value, opts, depth+1)
		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			printRecursive(v.Index(i), opts, depth+1)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			fmt.Println(key.Interface(), ": ")
			printRecursive(v.MapIndex(key), opts, depth+1)
		}
	default:
		fmt.Println(v.Interface())
	}
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
