package main

// https://leetcode.com/problems/integer-to-roman/

type Roman struct {
	Value        int
	Presentation string
}

func intToRoman(num int) string {
	res := ""
	romans := []Roman{
		Roman{Value: 1, Presentation: "I"},
		Roman{Value: 4, Presentation: "IV"},
		Roman{Value: 5, Presentation: "V"},
		Roman{Value: 9, Presentation: "IX"},
		Roman{Value: 10, Presentation: "X"},
		Roman{Value: 40, Presentation: "XL"},
		Roman{Value: 50, Presentation: "L"},
		Roman{Value: 90, Presentation: "XC"},
		Roman{Value: 100, Presentation: "C"},
		Roman{Value: 400, Presentation: "CD"},
		Roman{Value: 500, Presentation: "D"},
		Roman{Value: 900, Presentation: "CM"},
		Roman{Value: 1000, Presentation: "M"},
	}

	for num != 0 {
		for i := 12; i >= 0; i-- {
			if num >= romans[i].Value {
				res += romans[i].Presentation
				num -= romans[i].Value
				break
			}
		}
	}
	return res
}
