package main

import (
	"fmt"
	"strconv"
)

func main() {

	var (
		problem  string
		solution string
	)

	fmt.Println("Welcome to the best calculator ever!")
	for i := 1; i > 0; i++ {
		fmt.Print("Enter your problem (format: 1+1, 2-1,etc.): ")
		fmt.Scanln(&problem)
		solution = calculate(problem)
		fmt.Println("Solution:", solution)
	}
}

func calculate(problem string) string {

	var (
		solution      string
		parsedProblem [3]string
	)

	parsedProblem = problemParse(problem)

	if parsedProblem[1] == "+" {
		solution = add(parsedProblem)
	} else if parsedProblem[1] == "-" {
		solution = substract(parsedProblem)
	} else if parsedProblem[1] == "*" {
		solution = multiply(parsedProblem)
	} else if parsedProblem[1] == "/" {
		solution = divide(parsedProblem)
	} else {
		solution = "Operation invalid"
	}

	return solution
}

func problemParse(unparsedProblem string) [3]string {

	var parsedProblem [3]string

	for i, r := range unparsedProblem {
		parsedProblem[i] = string(r)
	}

	return parsedProblem
}

func add(problem [3]string) string {

	var (
		numX     int
		numY     int
		numZ     int
		solution string
		err      error
	)

	numX, err = strconv.Atoi(problem[0])
	if err != nil {
		fmt.Println("Error while converting:", err)
	}

	numY, err = strconv.Atoi(problem[2])
	if err != nil {
		fmt.Println("Error while converting:", err)
	}

	numZ = numX + numY

	solution = strconv.Itoa(numZ)

	return solution
}

func substract(problem [3]string) string {

	var (
		numX     int
		numY     int
		numZ     int
		solution string
		err      error
	)

	numX, err = strconv.Atoi(problem[0])
	if err != nil {
		fmt.Println("Error while converting:", err)
	}

	numY, err = strconv.Atoi(problem[2])
	if err != nil {
		fmt.Println("Error while converting:", err)
	}

	numZ = numX - numY

	solution = strconv.Itoa(numZ)

	return solution
}

func multiply(problem [3]string) string {

	var (
		numX     int
		numY     int
		numZ     int
		solution string
		err      error
	)

	numX, err = strconv.Atoi(problem[0])
	if err != nil {
		fmt.Println("Error while converting:", err)
	}

	numY, err = strconv.Atoi(problem[2])
	if err != nil {
		fmt.Println("Error while converting:", err)
	}

	numZ = numX * numY

	solution = strconv.Itoa(numZ)

	return solution
}

func divide(problem [3]string) string {

	var (
		numX     int
		numY     int
		numZ     int
		solution string
		err      error
	)

	numX, err = strconv.Atoi(problem[0])
	if err != nil {
		fmt.Println("Error while converting:", err)
	}

	numY, err = strconv.Atoi(problem[2])
	if err != nil {
		fmt.Println("Error while converting:", err)
	}

	numZ = numX / numY

	solution = strconv.Itoa(numZ)

	return solution
}
