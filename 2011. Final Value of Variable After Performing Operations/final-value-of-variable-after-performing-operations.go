package main

// https://leetcode.com/problems/final-value-of-variable-after-performing-operations/

func finalValueAfterOperations(operations []string) int {
	x := 0
	for i:=0; i<len(operations); i++{
		if string(operations[i][0]) == "+" || string(operations[i][1]) == "+"{
			x++
		}else{
			x--
		}
	}
	return x
}