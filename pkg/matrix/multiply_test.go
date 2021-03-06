package matrix

import "testing"

func TestMultiply(t *testing.T) {
	m := Matrix{
		{"1"},
	}

	ch := make(chan string)
	go m.Multiply(ch)
	multiplication := <-ch

	if multiplication != "1\n" {
		t.Errorf("Expected multiplication response to be the only element in the matrix.")
	}

	m = Matrix{
		{"1", "2"},
		{"3", "4"},
	}

	ch = make(chan string)
	go m.Multiply(ch)
	multiplication = <-ch

	if multiplication != "24\n" {
		t.Errorf("Expected multiplication response to be equal to 24.")
	}

	m = Matrix{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	ch = make(chan string)
	go m.Multiply(ch)
	multiplication = <-ch

	if multiplication != "362880\n" {
		t.Errorf("Expected multiplication response to be equal to 362880")
	}
}
