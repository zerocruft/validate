package validate

import (
	"errors"
	"reflect"
)

const (
	STRING = "string"
	INT    = "int"
)

func ValidateRequired(i interface{}) (error, []string) {
	var err error
	t := reflect.TypeOf(i)
	err = errors.New("Error: " + t.Name() + " has required fields")
	infractions := []string{}
	ivalue := reflect.ValueOf(i)

	for index := 0; index < t.NumField(); index++ {
		field := t.Field(index)
		name := field.Name
		required := field.Tag.Get("required")

		if required != "" {
			if field.Type.Name() == STRING {
				if ivalue.FieldByName(name).String() == "" {
					infractions = append(infractions, name)
				}
			}
		}
	}

	if len(infractions) == 0 {
		err = nil
		infractions = nil
	}

	return err, infractions
}
