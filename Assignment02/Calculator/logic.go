package Calculator

func Add(number1, number2 float32) float32 {
	return number1 + number2
}

func Sub(number1, number2 float32) float32 {
	return number1 - number2
}

func Multi(number1, number2 float32) float32 {
	return number1 * number2
}
func Div(number1, number2 int) (float32, int) {

	quotient, remainder := number1/number2, number1%number2
	fquo := float32(quotient)
	return fquo, remainder

}
