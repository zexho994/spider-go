package structs

import (
	"fmt"
	"testing"
)

func TestSkipList_randomLevel(t *testing.T) {
	sl := SkipList{level: 10}
	n1, n2, n3, n4, n5, n6, n7, n8, n9, n10, n64 := 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
	for i := 0; i < 10000; i++ {
		l := sl.randomLevel()
		switch l {
		case 1:
			n1++
		case 2:
			n2++
		case 3:
			n3++
		case 4:
			n4++
		case 5:
			n5++
		case 6:
			n6++
		case 7:
			n7++
		case 8:
			n8++
		case 9:
			n9++
		case 10:
			n10++
		default:
			n64++
		}
	}
	fmt.Printf("n1 -> %v\n", n1)
	fmt.Printf("n2 -> %v\n", n2)
	fmt.Printf("n3 -> %v\n", n3)
	fmt.Printf("n4 -> %v\n", n4)
	fmt.Printf("n5 -> %v\n", n5)
	fmt.Printf("n6 -> %v\n", n6)
	fmt.Printf("n7 -> %v\n", n7)
	fmt.Printf("n8 -> %v\n", n8)
	fmt.Printf("n9 -> %v\n", n9)
	fmt.Printf("n10 -> %v\n", n10)
	fmt.Printf("n64 -> %v\n", n64)
}
