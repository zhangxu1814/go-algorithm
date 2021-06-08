package integer_to_roman

func intToRoman(num int) string {
	stringM := []string{"", "M", "MM", "MMM"}
	stringC := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	stringX := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	stringI := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	return stringM[num/1000] + stringC[num/100%10] + stringX[num/10%10] + stringI[num%10]
}
