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

func ExampleObject_Chunk() {
	// Split the list into shorter length lists
	res, err := lambda.New([]int{1, 2, 3, 4, 5}).Chunk(2).ToIntListSlice()
	fmt.Println("err:", err)
	fmt.Println("res:", res)
	// output:
	// err: <nil>
	// res: [[1 2] [3 4] [5]]
}

func ExampleObject_Compact() {
	// Remove 0-valued elements from the list
	res, err := lambda.New([]interface{}{0, 1, false, true, "", "str"}).Compact().ToInterfaceSlice()
	fmt.Println("err:", err)
	fmt.Println("res:", res)
	// output:
	// err: <nil>
	// res: [1 true str]
}

func ExampleObject_Flatten() {
	// Flatten the list
	res, err := lambda.New([][]int{{1, 2}, {2, 3}, {4}}).Flatten().ToIntSlice()
	fmt.Println("err:", err)
	fmt.Println("res:", res)
	// output:
	// err: <nil>
	// res: [1 2 2 3 4]
}

func ExampleObject_Reverse() {
	// Reverse list
	res, err := lambda.New([]int{1, 2, 3, 4}).Reverse().ToIntSlice()
	fmt.Println("err:", err)
	fmt.Println("res:", res)
	// output:
	// err: <nil>
	// res: [4 3 2 1]
}

func ExampleObject_Uniq() {
	// Remove duplicate elements in the list
	res, err := lambda.New([]int{1, 2, 1, 3, 2, 3, 4}).Uniq().ToIntSlice()
	fmt.Println("err:", err)
	fmt.Println("res:", res)
	// output:
	// err: <nil>
	// res: [1 2 3 4]
}

func ExampleObject_MapList() {
	// Traverse the elements of the list, and after each element is processed, the returned elements form a new list
	res, err := lambda.New([]int{1, 2, 3}).MapList(func(idx int, obj interface{}) interface{} {
		return obj.(int) + 1
	}).ToIntSlice()
	fmt.Println("err:", err)
	fmt.Println("res:", res)
	// output:
	// err: <nil>
	// res: [2 3 4]
}

func ExampleObject_FilterList() {
	// Traverse the elements of the list, each element is added to a new list or not, and a new list is returned
	res, err := lambda.New([]int{1, 2, 3, 4}).FilterList(func(idx int, obj interface{}) bool {
		return obj.(int)%2 == 0
	}).ToIntSlice()
	fmt.Println("err:", err)
	fmt.Println("res:", res)
	// output:
	// err: <nil>
	// res: [2 4]
}
