package main

func getRomainFromFraction(fraction int) string {
	switch fraction {
	case 1000:
		return "M"
	case 900:
		return "CM"
	case 500:
		return "D"
	case 400:
		return "CD"
	case 100:
		return "C"
	case 90:
		return "XC"
	case 50:
		return "L"
	case 40:
		return "XL"
	case 10:
		return "X"
	case 9:
		return "IX"
	case 5:
		return "V"
	case 4:
		return "IV"
	default:
		return "I"
	}
}

func intToRoman(num int) string {

	out := ""

	fractions := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	for _, v := range fractions {
		if num/v > 0 {
			roman := getRomainFromFraction(v)
			for i := 0; i < num/v; i++ {
				out += roman
			}
			num = num % v
		}
	}
	return out
}
