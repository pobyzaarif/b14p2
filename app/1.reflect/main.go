package main

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-playground/validator/v10"
)

func main() {
	spew.Dump()

	initValidator := validator.New()

	// learnBasicReflect()
	// learnStructReflect()
	learnValidateStructWithValidator(initValidator)
}

func learnBasicReflect() {
	var freeNumber interface{}
	freeNumber = 1
	fmt.Println(freeNumber)

	reflectA := reflect.ValueOf(freeNumber)
	fmt.Println(reflectA.Type())

	fmt.Println(reflectA.Int())

	var flexibleValue interface{}
	flexibleValue = 2
	fmt.Println(flexibleValue)
}

type User struct {
	Email    string `required:"true" validate:"required,email"`
	Fullname string `min:"8" max:"64"`
	Age      int    `min:"18"`
}

func learnStructReflect() {
	newUser := User{
		// Email:    "learn@gmail.com",
		Fullname: "learner",
		Age:      1,
	}

	userValue := reflect.ValueOf(newUser)
	fmt.Println(userValue)

	userType := reflect.TypeOf(newUser)
	fmt.Println(userType)

	userField := userType.NumField()
	fmt.Println(userField)

	fmt.Println(userType.Field(0))
	fmt.Println(userType.Field(0).Name)
	// spew.Dump(userType.Field(0))

	fmt.Println(userType.Field(0).Tag.Get("required"))

	fmt.Println(ValidateStruct(newUser))
}

func ValidateStruct(s interface{}) error {
	t := reflect.TypeOf(s)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(s).Field(i).Interface()
			if value == "" {
				err := errors.New("some field is required, and your value is empty or 0")
				// return fmt.Errorf("%s is required", field.Name)
				return err
			}
		}
		if field.Tag.Get("min") != "" {
			// TODO get min value
			// Compare value length is greater than min value, if less than min value then return error
		}
	}
	return nil
}

func ValidateStructWithValidator(v *validator.Validate, u User) error {
	err := v.Struct(u)
	if err != nil {
		return err
	}

	return nil
}

func learnValidateStructWithValidator(v *validator.Validate) {
	newUser := User{
		Email:    "learner@gmail.com",
		Fullname: "learner",
		Age:      1,
	}
	fmt.Println(ValidateStructWithValidator(v, newUser))

}
