package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-playground/validator/v10"
)

type GroupCreateBody struct {
	ParentID string `json:"parentID,omitempty" validate:"uuid"`
}

func main() {
	v := validator.New()
	g := GroupCreateBody{}

	if err := v.Struct(&g); err != nil {
		fmt.Printf("err %#v\n", err) // add stack trace
		spew.Dump(err)

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {

			fmt.Println("err.Namespace()", err.Namespace())
			fmt.Println("err.Field()", err.Field())
			fmt.Println("err.StructNamespace()", err.StructNamespace())
			fmt.Println("err.StructField()", err.StructField())
			fmt.Println("err.Tag()", err.Tag())
			fmt.Println("err.ActualTag()", err.ActualTag())
			fmt.Println("err.Kind()", err.Kind())
			fmt.Println("err.Type()", err.Type())
			fmt.Println("err.Value()", err.Value())
			fmt.Println("err.Param()", err.Param())
			fmt.Println()
		}

		// from here you can create your own error messages in whatever language you wish

		return
	}
}
