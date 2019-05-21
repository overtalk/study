package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/caarlos0/env"
)

type ReflectTestDemo struct {
	A string `env:"A" allowempty:"true" required:"aaa,bbb"`
	B int    `env:"B" allowempty:"true"`
	C bool   `env:"C" allowempty:"true"`
}

type ReflectTest struct {
	AAA string          `env:"AAA" allowempty:"false"`
	BBB int             `env:"BBB" allowempty:"true"`
	CCC bool            `env:"CCC" allowempty:"false" required:"aaa,bbb"`
	ddd ReflectTestDemo `env:"ddd" allowempty:"true"`
}

func main() {
	test := ReflectTest{}
	if err := env.Parse(&test); err != nil {
		log.Fatalf("failed to parse from env : %v", err)
	}
	ReflectAnalyse(test)
}

func ReflectAnalyse(test interface{}) {
	refType := reflect.TypeOf(test)
	refValue := reflect.ValueOf(test)

	log.Println("value of the struct : ", refValue)
	log.Println("type of the value : ", refType)

	for i := 0; i < refType.NumField(); i++ {
		refField := refValue.Field(i)
		if refField.Kind() == reflect.Ptr && !refField.IsNil() {
			continue
		}
		refTypeField := refType.Field(i)
		allowEmpty := refTypeField.Tag.Get("allowempty") == "true"
		fmt.Printf("filed %v is allow empty : %v\n", refTypeField.Name, allowEmpty)

		fmt.Println("required : ", strings.Split(refTypeField.Tag.Get("required"), ","))

		if refField.Kind() == reflect.Struct {
			fmt.Printf("filed %s is a struct\n", refTypeField.Name)
		}
	}
}
