package err_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

var LessThanTwoError = errors.New("n should be not less than 2")
var LargerThanHundredError = errors.New("n should be not larger than 100")

func GetFibonnacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}

	if n > 100 {
		return nil, LargerThanHundredError
	}

	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonnacci(1); err != nil {
		if err == LessThanTwoError {
			fmt.Println("It is less.")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func GetFibonnacci1(str string) {
	var (
		i    int
		err  error
		list []int
	)

	if i, err = strconv.Atoi(str); err == nil {
		if list, err = GetFibonnacci(i); err == nil {
			fmt.Println(list)
		} else {
			fmt.Println("Error", err)
		}
	} else {
		fmt.Println("Error", err)
	}
}

func GetFibonnacci2(str string) {
	var (
		i    int
		err  error
		list []int
	)

	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error", err)
		return
	}

	if list, err = GetFibonnacci(i); err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(list)
}
