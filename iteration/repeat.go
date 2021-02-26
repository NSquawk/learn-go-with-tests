package iteration

const repeatCount = 5

func Repeat(toRepeat string, times int) string {
	repeated := ""
	for i := 0; i < times; i++ {
		repeated += toRepeat
	}
	return repeated
}
