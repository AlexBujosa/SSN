package main

import (
	"strings"
	"testing"
)

func TestCase1(t *testing.T){
	want := "Social Security Number valid"
	input := "001-29-1829"
	splitinput := strings.Split(input, "-")
	var validatorSplit []string
	var message string
	for i:= 0; i< len(splitinput); i++ {
		message = validate(validatorSplit, splitinput[i], pos, i)
		if message != "Social Security Number valid" {
			break
		}
	}
	if want != message {
		t.Errorf("The value expected is %q and it return %q", want, message)
	}
}
func TestCase2(t *testing.T){
	want := "Social Security Number not valid"
	input := "000-34-2817"
	splitinput := strings.Split(input, "-")
	var validatorSplit []string
	var message string
	for i:= 0; i< len(splitinput); i++ {
		message = validate(validatorSplit, splitinput[i], pos, i)
		if message != "Social Security Number valid" {
			break
		}
	}
	if want != message {
		t.Errorf("The value expected is %q and it return %q", want, message)
	}
}
func TestCase3(t *testing.T){
	want := "Social Security Number not valid"
	input := "666-48-3899"
	splitinput := strings.Split(input, "-")
	var validatorSplit []string
	var message string
	for i:= 0; i< len(splitinput); i++ {
		message = validate(validatorSplit, splitinput[i], pos, i)
		if message != "Social Security Number valid" {
			break
		}
	}
	if want != message {
		t.Errorf("The value expected is %q and it return %q", want, message)
	}
}
func TestCase4(t *testing.T){
	want := "Social Security Number not valid"
	input := "901-23-2749"
	splitinput := strings.Split(input, "-")
	var validatorSplit []string
	var message string
	for i:= 0; i< len(splitinput); i++ {
		message = validate(validatorSplit, splitinput[i], pos, i)
		if message != "Social Security Number valid" {
			break
		}
	}
	if want != message {
		t.Errorf("The value expected is %q and it return %q", want, message)
	}
}
func TestCase5(t *testing.T){
	want := "Social Security Number not valid"
	input := "567-00-4837"
	splitinput := strings.Split(input, "-")
	var validatorSplit []string
	var message string
	for i:= 0; i< len(splitinput); i++ {
		message = validate(validatorSplit, splitinput[i], pos, i)
		if message != "Social Security Number valid" {
			break
		}
	}
	if want != message {
		t.Errorf("The value expected is %q and it return %q", want, message)
	}
}
func TestCase6( t *testing.T){
	want := "Social Security Number not valid"
	input := "567-99-0000"
	splitinput := strings.Split(input, "-")
	var validatorSplit []string
	var message string
	for i:= 0; i< len(splitinput); i++ {
		message = validate(validatorSplit, splitinput[i], pos, i)
		if message != "Social Security Number valid" {
			break
		}
	}
	if want != message {
		t.Errorf("The value expected is %q and it return %q", want, message)
	}
}
func TestCase7(t *testing.T){
	want := "Social Security Number valid"
	input := "567-99-0001"
	splitinput := strings.Split(input, "-")
	var validatorSplit []string
	var message string
	for i:= 0; i< len(splitinput); i++ {
		message = validate(validatorSplit, splitinput[i], pos, i)
		if message != "Social Security Number valid" {
			break
		}
	}
	if want != message {
		t.Errorf("The value expected is %q and it return %q", want, message)
	}
}