// package math

// import "testing"

// func TestAdd(t *testing.T) {
//   	sum := Add(4,6)
// 	want :=10
// 	 if sum != want{
// 		t.Errorf("got %d, wanted %d",sum,want)
// 	 }

// }

package math

import (
	"fmt"
	"testing"
)

type addTest struct {
	frst_num, sec_num, expected_res int
}

var addTests = []addTest{

	addTest{2, 3, 5},
	addTest{4, 8, 12},
	addTest{6, 9, 15},
	addTest{3, 10, 13},
}

func TestAdd(t *testing.T) {
	for _, test := range addTests {
		if output := Add(test.frst_num, test.sec_num); output != test.expected_res {
			t.Errorf("output %d not equal to expected %d", output, test.expected_res)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 6)
	}
}

func ExampleAdd() {
	fmt.Println(Add(4, 6))
	// Output: 10
}
