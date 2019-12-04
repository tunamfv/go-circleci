package main

import (
	"errors"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("invalid parameter a", func(t *testing.T) {
		a := 1
		b := 6

		wantError := ErrParameterGreaterThanFive
		_, err := sum(a, b)


		if !errors.Is(wantError, err){
			t.Errorf("want error:%v but got:%v", wantError, err)
		}
	})

	t.Run("invalid parameter b", func(t *testing.T) {
		a := 6
		b := 5

		wantError := ErrParameterGreaterThanFive
		_, err := sum(a, b)


		if !errors.Is(wantError, err){
			t.Errorf("want error:%v but got:%v", wantError, err)
		}
	})

	t.Run("success", func(t *testing.T) {
		a := 6
		b := 6

		want := 12
		got, err := sum(a, b)

		if err != nil {
			t.Errorf("Sum() got an error: %v", err)
		}

		if want != *got {
			t.Errorf("want: %d but got: %d", want, got)
		}
	})
}