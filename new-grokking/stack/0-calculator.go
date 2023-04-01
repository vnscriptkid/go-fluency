package stack

import "container/list"

func calculator(expression string) int {
	// 12 - ( 6 + 2 ) + 5
	//                  ^

	// stack [ 12 -1 ]

	// cur 5
	// temp 9
	// sign 1

	//////////////////

	// ( 8 + 100 ) + ( 13 - 8 - ( 2 + 1 ) )
	//                                    ^

	// stack [108 1]

	// cur 2
	// temp 110
	// sign 1

	// Generalize
	// if ( ->
	// 1) push temp, sign to stack (if exist)
	// 2) reset cur, temp, sign

	stack := list.New()

	cur := 0
	temp := 0
	sign := 1

	for i, c := range expression {

		switch c {
		case '+':
			sign = 1
			cur = 0
		case '-':
			sign = -1
			cur = 0
		case '(':
			{
				if temp == 0 && sign == 1 {
					continue
				}
				stack.PushBack(temp)
				stack.PushBack(sign)

				cur = 0
				temp = 0
				sign = 1
			}
		case ')':
			{
				if stack.Len() >= 2 {
					cur = temp

					sign = stack.Remove(stack.Back()).(int)
					temp = stack.Remove(stack.Back()).(int)

					temp = temp + sign*cur
				}
			}
		case ' ':
			continue
		default:
			// it's a digit
			if isDigit(c) {
				digit := int(c - '0')
				cur = cur*10 + digit

				if i+1 == len(expression) || !isDigit(rune(expression[i+1])) {
					temp = temp + cur*sign
				}
			}
		}
	}

	return temp
}

func isDigit(x rune) bool {
	return x >= '0' && x <= '9'
}
