package panic_recover_test

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from", err)
		}
	}()

	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}
