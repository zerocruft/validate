package validate

import (
	"errors"
	"reflect"
)

const (
	VALIDATE = "validate"
	REQUIRED = "required"
	STRING   = "string"
	INT      = "int"
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
		isValidate := field.Tag.Get(VALIDATE)

		// TODO will want to check what type of validation is happening. For now, only required is respected
		if isValidate != "" {
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
