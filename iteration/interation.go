package iteration

const repeatCount = 5

func Repeat(character string) string {
	var repeted string
	for i := 0; i < repeatCount; i++ {
		repeted += character
	}
	return repeted
}
