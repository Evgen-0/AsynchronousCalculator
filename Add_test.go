package main

import "testing"

func TestAdd(t *testing.T) {
	t.Run("Correct", func(t *testing.T) {
		t.Parallel()
		expression := "10+10"
		result := 20
		realResult := Add(expression)
		if realResult != result {
			t.Errorf("expected result: %d != %d real result", result, realResult)
		}
	})
	t.Run("Bad", func(t *testing.T) {
		t.Parallel()
		expression := "10+k"
		result := Add(expression)
		if result == 0 {

		}
	})
	t.Run("panic", func(t *testing.T) {
		t.Parallel()
		expression := "1"
		result := 0
		realResult := Add(expression)
		if realResult != result {
			t.Errorf("panic")
		}

	})

}
