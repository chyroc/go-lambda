package lambda_test

import (
	"fmt"

	"github.com/chyroc/go-lambda"
)

type exampleItem struct {
	Name string
}

func ExampleNew() {
	// new with string list
	lambda.New([]string{"1", "2"})

	// new with struct
	lambda.New([]exampleItem{{Name: "Tom"}, {Name: "Harry"}})
}

func ExampleObject_Array() {
	// new with int list, and incr every one, return incr-ed int list
	res, err := lambda.New([]int{1, 2, 3}).MapArray(func(idx int, obj interface{}) interface{} {
		return obj.(int) + 1
	}).ToIntList()
	fmt.Println("err", err)
	fmt.Println("res", res) // be: [2, 3, 4]
}
