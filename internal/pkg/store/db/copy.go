package db

import (
	"errors"
	"reflect"
)

// Copy .
func Copy(dest interface{}, source interface{}) (err error) {
	destValue := indirectValue(reflect.ValueOf(dest))
	sourceValue := indirectValue(reflect.ValueOf(source))

	if !destValue.CanAddr() {
		return errors.New("copy to value is unaddressable")
	}

	if !sourceValue.IsValid() {
		return
	}

	destType := indirectType(destValue.Type())
	sourceType := indirectType(sourceValue.Type())

	if sourceType.Kind() != reflect.Struct || destType.Kind() != reflect.Struct {
		return
	}

	sourceTypeFields := deepFields(sourceType)

	for _, field := range sourceTypeFields {
		name := field.Name
		if sourceField := sourceValue.FieldByName(name); sourceField.IsValid() {
			if destField := destValue.FieldByName(name); destField.IsValid() {
				if destField.CanSet() {
					var val reflect.Value
					switch destField.Kind() {
					case reflect.String:
						if sourceField.Kind() == reflect.Struct {
							val = sourceField.FieldByName("String")
						}
					case reflect.Int32:
						if sourceField.Kind() == reflect.Struct {
							val = sourceField.FieldByName("Int32")
						}
					case reflect.Int64:
						if sourceField.Kind() == reflect.Struct {
							val = sourceField.FieldByName("Int64")
						}
					case reflect.Float64:
						if sourceField.Kind() == reflect.Struct {
							val = sourceField.FieldByName("Float64")
						}
					case reflect.Bool:
						if sourceField.Kind() == reflect.Struct {
							val = sourceField.FieldByName("Bool")
						}
					}

					if val.IsValid() && !val.IsZero() {
						if val.Type() != destField.Type() {
							if val.Type().ConvertibleTo(destField.Type()) {
								destField.Set(val.Convert(destField.Type()))
							}
						} else {
							destField.Set(val)
						}
					}
				}
			}
		}
	}

	return
}

func indirectValue(reflectValue reflect.Value) reflect.Value {
	for reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	return reflectValue
}

func indirectType(reflectType reflect.Type) reflect.Type {
	for reflectType.Kind() == reflect.Ptr {
		reflectType = reflectType.Elem()
	}
	return reflectType
}

func deepFields(reflectType reflect.Type) []reflect.StructField {
	var fields []reflect.StructField

	if reflectType = indirectType(reflectType); reflectType.Kind() == reflect.Struct {
		for i := 0; i < reflectType.NumField(); i++ {
			v := reflectType.Field(i)
			if v.Anonymous {
				fields = append(fields, deepFields(v.Type)...)
			} else {
				fields = append(fields, v)
			}
		}
	}

	return fields
}
