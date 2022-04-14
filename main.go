package main

import (
	"fmt"
	"strings"
)

func main() {
	input := [][]string{
		{"Alan", "James"},
		{"Alex", "Tiffany"},
		{"Ray", "Tiffany"},
		{"James", "Daniel"},
		{"Tiffany", "Daniel"},
	}

	// leaderName := getLeastCommonLeader(input, "Alex", "Ray") // Expect output: Tiffany
	// leaderName := getLeastCommonLeader(input, "James", "Tiffany") // Expect output: Daniel
	// leaderName := getLeastCommonLeader(input, "Ray", "Tiffany") // Expect output: Daniel
	// leaderName := getLeastCommonLeader(input, "Alex", "Tiffany") // Expect output: Daniel

	leaderName := getLeastCommonLeader(input, "Alex", "Tiffany") // Daniel

	fmt.Println(leaderName)
}

func getLeastCommonManager(employees [][]string, employeeOne string, employeeTwo string) string {

	return "Choose a method"
}

func getLeastCommonLeader(employees [][]string, employeeOne string, employeeTwo string) string {
	var leaders []string
	for _, employee := range employees {
		if strings.EqualFold(employee[0], employeeOne) || strings.EqualFold(employee[0], employeeTwo) {
			leaders = append(leaders, employee[1])
		}
	}

	if len(leaders) == 2 {
		if strings.EqualFold(leaders[0], leaders[1]) {
			return leaders[0]
		} else {
			return getLeastCommonLeader(employees, leaders[0], leaders[1])
		}
	}

	if len(leaders) == 1 {
		return leaders[0]
	}

	return "None leader was found"
}
