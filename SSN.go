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
		
	}
	return "Social Security Number valid"
}
func main(){
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Ingrese su numero social: ")
	reader.Scan()
	input := reader.Text()
	splitinput := Splitstring(input, pos)
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