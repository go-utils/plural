package plural_test

import (
	"testing"

	"github.com/go-utils/plural"
)

func TestPlural(t *testing.T) {
	t.Run("normal", func(tr *testing.T) {
		words := [][]string{
			{"cat", "cats"},
			{"dog", "dogs"},
			{"pig", "pigs"},
		}

		for _, pair := range words {
			actual := plural.Convert(pair[0])
			expected := pair[1]
			if actual != expected {
				tr.Fatalf("%s: unexpected, actual: `%v`, expected: `%v`", tr.Name(), actual, expected)
			}
		}
	})

	t.Run("o/s/x/ch/sh", func(tr *testing.T) {
		words := [][]string{
			{"tomato", "tomatoes"},
			{"axis", "axises"},
			{"box", "boxes"},
			{"church", "churches"},
			{"dish", "dishes"},
		}

		for _, pair := range words {
			actual := plural.Convert(pair[0])
			expected := pair[1]
			if actual != expected {
				tr.Fatalf("%s: unexpected, actual: `%v`, expected: `%v`", tr.Name(), actual, expected)
			}
		}
	})

	t.Run("f/fe", func(tr *testing.T) {
		words := [][]string{
			{"leaf", "leaves"},
			{"knife", "knives"},
		}

		for _, pair := range words {
			actual := plural.Convert(pair[0])
			expected := pair[1]
			if actual != expected {
				tr.Fatalf("%s: unexpected, actual: `%v`, expected: `%v`", tr.Name(), actual, expected)
			}
		}
	})

	t.Run("y", func(tr *testing.T) {
		words := [][]string{
			{"play", "plays"},
			{"study", "studies"},
		}

		for _, pair := range words {
			actual := plural.Convert(pair[0])
			expected := pair[1]
			if actual != expected {
				tr.Fatalf("%s: unexpected, actual: `%v`, expected: `%v`", tr.Name(), actual, expected)
			}
		}
	})
}
