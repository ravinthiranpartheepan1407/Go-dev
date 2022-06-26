package main

import (
	"fmt"
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	data := "xyz@gmail.com"
	if data != "xyz@gmail.com" {
		t.Errorf("The email (%v)=false is wrong", data)
	} else {
		fmt.Println("True")
	}
}

func TestIsAFruit(t *testing.T) {
	// fruits := []string{"Apple", "Banana", "Watermelon"}
	fruits := "Apple"

	if fruits != "Apples" {
		t.Fail()
	} else {
		t.Logf(fruits)
	}
}

func TestIsValidTable(t *testing.T) {
	collection := []struct {
		name   string
		number int
	}{
		{"Ravinthiran", 1},
		{"Suren", 2},
	}

	for _, data := range collection {
		outName := data.name
		// outId := data.number

		if outName != "Ravinthiran" && outName != "Suren" {
			t.Errorf("There is not data updated(%v)=false", outName)
		} else {
			t.Logf(outName)
		}
	}
}
