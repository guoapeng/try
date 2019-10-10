package try_test

import (
	"fmt"
	. "github.com/guoapeng/try"
)

func ExampleRuntimeError() {
	Try{ func(a, b, c int) (int,error) {
		return a/b, nil
	}}.Catch(func (ex error){
		fmt.Println("error happens", ex)
	}).Go(4, 0, 3)
	// Output:
	// error happens runtime error: integer divide by zero
}

func ExampleCheckError() {
	Try{ func(a, b, c int) (int,error) {
		if a % b != 0 {
			return 0, fmt.Errorf("%d is not aliquot by %d", a, b)
		}
		return a/b, nil
	}}.Catch(func (ex error){
		fmt.Println("error happens", ex)
	}).Go(5, 2, 3)
	// Output:
	// error happens 5 is not aliquot by 2
}

func ExampleUpdateLocalVariable() {
	local := 100
	Try{ func(a, b, c int)  {
		local = a*b*c
	}}.Catch(func (ex error){
		fmt.Println("error happens", ex)
	}).Go(5, 5, 5)
	fmt.Println(local)
	// Output:
	// 125
}

func ExampleCheckHandlePanic() {
	Try{ func(a, b int) (int,error) {
		panic("mandatory service down")
		return a/b, nil
	}}.Catch(func (ex error){
		fmt.Println("error happens", ex)
	}).Go(5, 2)
	// Output:
	// error happens mandatory service down
}


func ExampleHandleMultipleDatType() {
	Try{ func(a string , b, c int) {
		fmt.Println(a, b/c)
	}}.Catch(func (ex error){
		fmt.Println("error happens", ex)
	}).Go("b/c =", 8, 2)
	// Output:
	// b/c = 4
}
