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
		realResult := Add(expression)
		result := 0
		if realResult != result {
			t.Errorf("expected realResult: %d != %d real realResult", realResult, result)
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
