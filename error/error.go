package errs

import (
	"fmt"
	"reflect"
)

func PanicErr(err error) {
	if err != nil {
		fmt.Println(reflect.TypeOf(err).String())
		fmt.Println(reflect.ValueOf(err).String())
		panic(err)
	}
}

