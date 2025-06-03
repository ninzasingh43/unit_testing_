package math

import "testing"

func TestAdd(t *testing.T) {
  	sum := Add(4,6)
	want :=10
	 if sum != want{
		t.Errorf("got %d, wanted %d",sum,want)
	 }

}