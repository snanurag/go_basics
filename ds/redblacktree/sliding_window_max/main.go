package main

import (
	"fmt"
	treeset "github.com/emirpasic/gods/sets/treeset"
)

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3))
}

type Item struct {
	i int
}

func maxSlidingWindow(nums []int, k int) []int {

	var r []int
	s := treeset.NewWithIntComparator()
	i := 0
	for s.Size() < k && i < len(nums) {
		s.Add(nums[i])
		//fmt.Println(s)
		if s.Size() == k {
			val, ok := getMaxValue(s)
			if ok {
				r = append(r, val)
			} else {
				return nil
			}
			s.Remove(nums[i-k+1])
		}
		i++
	}
	return r
}

func getMaxValue(s *treeset.Set) (int, bool) {
	it := s.Iterator()
	it.End()
	if it.Prev() {
		return it.Value().(int), true
	} else {
		return -1, false
	}

}