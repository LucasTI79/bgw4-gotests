package calculator

import "fmt"

func Add(num1, num2 int) (int, error) {
	result := num1 + num2
	return result, nil
}

func Division(numerador, denominador int) (int, error) {
	if denominador == 0 {
		panic("o denominador deve ser diferente de 0")
		return 0, fmt.Errorf("o denominador deve ser diferente de 0")
	}
	result := numerador / denominador
	return result, nil
}
