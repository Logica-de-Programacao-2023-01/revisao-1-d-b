package q5

func ConvertTemperature(temp float64, fromScale string, toScale string) (float64, error) {
	var finalTemp float64
	switch fromScale {
	case "C":
		switch toScale {
		case "F":
			finalTemp = temp*9/5 + 32
		case "K":
			finalTemp = temp + 273.15
		default:
			return 0, fmt.Errorf("escala inválida")
		}
	case "F":
		switch toScale {
		case "C":
			finalTemp = (temp - 32) * 5 / 9
		case "K":
			finalTemp = (temp-32)*5/9 + 273.15
		default:
			return 0, fmt.Errorf("escala inválida")
		}
	case "K":
		switch toScale {
		case "C":
			finalTemp = temp - 273.15
		case "F":
			finalTemp = (temp-273.15)*9/5 + 32
		default:
			return 0, fmt.Errorf("escala inválida")
		}
	default:
		return 0, fmt.Errorf("escala inválida")
	}
	return finalTemp, nil
}
