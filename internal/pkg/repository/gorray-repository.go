package repository

import (
	"fmt"
	"strconv"
	"strings"
)

type GorrayRepository struct {}

var gorrayRepository * GorrayRepository

func GetGorrayRepository() *GorrayRepository{
	if gorrayRepository == nil {
		gorrayRepository = &GorrayRepository{}
	}
	return gorrayRepository
}


func formatMatrix(records [][]string) string{
	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return response
}

func (g *GorrayRepository) Echo(slice [][]string) string{
	return formatMatrix(slice)
}

func (g *GorrayRepository) Transpose(slice [][]string) string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}

	return formatMatrix(result)
}

func (g *GorrayRepository) Flatten(slice [][]string) []string {
	xl := len(slice[0])
	yl := len(slice)
	var result []string
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result = append(result, slice[i][j])
		}
	}	
	return result
}

func (g *GorrayRepository) Sum(slice [][]string) int {
	xl := len(slice[0])
	yl := len(slice)
	result := 0
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			num, _ := strconv.Atoi(slice[i][j])
			result += num
		}
	}
	return result
}

func (g *GorrayRepository) Multiply(slice [][]string) int {
	xl := len(slice[0])
	yl := len(slice)
	result := 1
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			num, _ := strconv.Atoi(slice[i][j])
			result *= num
		}
	}
	return result
}