package try_test

import (
	"fmt"
	. "github.com/guoapeng/try"
)

func ExampleRuntimeError() {
	Try{func(a, b, c int) (int, error) {
		return a / b, nil
	}}.Catch(func(ex error) {
		fmt.Println("error happens", ex)
	}).Go(4, 0, 3)
	// Output:
	// error happens runtime error: integer divide by zero
}

func ExampleCheckError() {
	Try{func(a, b, c int) (int, error) {
		if a%b != 0 {
			return 0, fmt.Errorf("%d is not aliquot by %d", a, b)
		}
		return a / b, nil
	}}.Catch(func(ex error) {
		fmt.Println("error happens", ex)
	}).Go(5, 2, 3)
	// Output:
	// error happens 5 is not aliquot by 2
}

func ExampleUpdateLocalVariable() {
	local := 100
	Try{func(a, b, c int) {
		local = a * b * c
	}}.Catch(func(ex error) {
		fmt.Println("error happens", ex)
	}).Go(5, 5, 5)
	fmt.Println(local)
	// Output:
	// 125
}

func ExampleCheckHandlePanic() {
	Try{func(a, b int) (int, error) {
		panic("mandatory service down")
		return a / b, nil
	}}.Catch(func(ex error) {
		fmt.Println("error happens", ex)
	}).Go(5, 2)
	// Output:
	// error happens mandatory service down
}

func ExampleHandleMultipleDatType() {
	Try{func(a string, b, c int) {
		fmt.Println(a, b/c)
	}}.Catch(func(ex error) {
		fmt.Println("error happens", ex)
	}).Go("b/c =", 8, 2)
	// Output:
	// b/c = 4
}

func ExampleGetValueFromFunc() {
	fn := Try{func(a, b int) int {
		return a / b
	}}.Catch(func(ex error) {
		fmt.Println("error happens", ex)
	})

	r := fn.Go(8, 8)
	fmt.Printf("8/8 = %v\n", r[0].Interface())
	r = fn.Go(8, 0) // error happened, r = nil
	// Output:
	// 8/8 = 1
	// error happens runtime error: integer divide by zero
}

func ExampleHowToReturnValue() {
	r := divide(8, 4)
	fmt.Println("8/4 =", r)
	//Output:
	// 8/4 = 2
}

func ExampleHowToThrowException() {
	Try{func() {
		r := divideThrowsRuntimeExption(8, 0)
		fmt.Println("8/4 =", r)
	}} .Catch(func(ex error) {
		fmt.Println("error handled at point B:", ex)
	}).Go()

	//Output:
	// error handled at point B: it can't be handle at point A - runtime error: integer divide by zero
}

func divideThrowsRuntimeExption(x, y int) (z int) {
	Try{func(a, b int) {
		z = a / b
		return
	}}.Catch(func(ex error) {
		if err, ok := ex.(error); ok {
			panic(fmt.Errorf("it can't be handle at point A - %s", err))
		}
		fmt.Println("error happens", ex)
	}).Go(x, y)
	return
}

func divide(x, y int) (z int) {
	Try{func(a, b int) {
		z = a / b
		return
	}}.Catch(func(ex error) {
		fmt.Println("error happens", ex)
	}).Go(x, y)
	return
}
