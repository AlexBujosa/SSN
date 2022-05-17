package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var validator string = "!000!666!900-999|3,01-99|2,0001-9999|4"
var symbol [4]string = [4]string{",", "|", "!","-"} 
var pos = 0
func Splitstring(text string, pos int) []string{
	return strings.Split(text, symbol[pos])
}
func removeEmpty(arr []string) []string{
	var newArr []string
	for i:=0; i< len(arr); i++{
		if arr[i] != "" {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr

}
func validate(validatorSplit []string, input string, pos int, iteration int) string{
	if len(validatorSplit) == 0 {
		validatorSplit = Splitstring(validator, pos)
		rtrn := validate(validatorSplit, input, pos, iteration)
		return rtrn
	}else if pos == 0 {
		pos += 1
		validateSpl := Splitstring(validatorSplit[iteration], pos)
		num, err := strconv.Atoi(validateSpl[1])
		if len(input) != num  || err != nil{
			return "Social Security Number not valid"
		}
		rtrn := validate(validateSpl, input, pos, iteration)
		return rtrn
	} else if pos == 1{
		pos += 1
		validateSpl := Splitstring(validatorSplit[0], pos)
		validateSpl = removeEmpty(validateSpl)
		for i:= 0; i<len(validateSpl); i++ {
			_, err := strconv.Atoi(validateSpl[i])
			if err != nil {
				valSplit := Splitstring(validateSpl[i], pos +1)
				if len(validateSpl) == 1 {
					rtrn := validate(valSplit, input, pos + 1, iteration)
					return rtrn
				}else {
					rtrn :=  validate(valSplit, input, pos, iteration)
					return rtrn
				}
			}
			if validateSpl[i] == input {
				return "Social Security Number not valid"
			}
		}
	} else {
		firstNumber, _ := strconv.Atoi(validatorSplit[0])
		secondNumber, _ := strconv.Atoi(validatorSplit[1])
		ssn, _ := strconv.Atoi(input)
		if (firstNumber <= ssn && secondNumber >= ssn) && pos == 2 {
			return "Social Security Number not valid"
		} else if (firstNumber > ssn || secondNumber < ssn) && pos != 2{
			return "Social Security Number not valid"
		}
	}
	return "Social Security Number valid"
}
func main(){
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Ingrese su numero social: ")
	reader.Scan()
	input := reader.Text()
	splitinput := strings.Split(input, "-")
	var validatorSplit []string
	var message string
	for i:= 0; i< len(splitinput); i++ {
		message = validate(validatorSplit, splitinput[i], pos, i)
		if message != "Social Security Number valid" {
			break
		}
	}
	if message != "Social Security Number valid"{
		fmt.Println("Social Security Number not valid")
	}else {
		fmt.Println("Social Security Number valid")
	}


}