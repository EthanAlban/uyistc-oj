package codesfiles

import (
	"fmt"
	"testing"
)

func TestAppend_0(t *testing.T) {
	base := make([]int, 0, 1)
	base = append(base, 100)
	fmt.Printf("len(base) = %v, cap(base) = %v\n", len(base), cap(base))
	// equivalent
	// base := []int{100}
	a := append(base, 200)
	b := append(base, 300)
	fmt.Printf("base = %v\na = %v\nb = %v\n", base, a, b)
	fmt.Printf("addr:\nbase = %p\n   a = %p\n   b = %p\n", base, a, b)
	fmt.Printf("len(a) = %v, cap(a) = %v\n", len(a), cap(a))
	fmt.Printf("len(b) = %v, cap(b) = %v\n", len(b), cap(b))
}

func TestAppend_1(t *testing.T) {
	base := make([]int, 0, 2)
	base = append(base, 100)
	fmt.Printf("len(base) = %v, cap(base) = %v\n", len(base), cap(base))
	a := append(base, 200)
	b := append(base, 300)
	fmt.Printf("base = %v\na = %v\nb = %v\n", base, a, b)
	fmt.Printf("addr:\nbase = %p\n   a = %p\n   b = %p\n", base, a, b)
	fmt.Printf("len(a) = %v, cap(a) = %v\n", len(a), cap(a))
	fmt.Printf("len(b) = %v, cap(b) = %v\n", len(b), cap(b))
}

func TestSliceAppend(t *testing.T) {
	ret := [][]int{}
	ret2 := [][]int{}
	slices := [][]int{{2, 2, 3}, {7}, {8}}

	for i := 0; i < len(slices); i++ {
		ret = append(ret, append([]int(nil), slices[i]...))
		ret2 = append(ret2, slices[i])
		fmt.Println("------------------------------")
		//fmt.Println(slices)
		fmt.Printf("ret: cap:%v len:%v val:%v\n", cap(ret), len(ret), ret)
		fmt.Printf("ret: cap:%v len:%v val:%v\n", cap(ret2), len(ret2), ret2)
	}

}
