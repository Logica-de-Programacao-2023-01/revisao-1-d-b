package q5

func ConvertTemperature(temp float64, fromScale string, toScale string) (float64, error) {
	var err error
	var convert float64

	if fromScale == "C" && toScale == "F" {
		convert := temp*9/5 + 32
		return convert, err
	} else if fromScale == "C" && toScale == "K" {
		convert := temp + 273.15
		return convert, err
	} else if fromScale == "F" && toScale == "C" {
		convert := (temp - 32) * 5 / 9
		return convert, err
	} else if fromScale == "F" && toScale == "K" {
		convert := (temp-32)*5/9 + 273.15
		return convert, err
	} else if fromScale == "K" && toScale == "C" {
		convert := temp - 273.15
		return convert, err
	} else if fromScale == "K" && toScale == "F" {
		convert := (temp-273.15)*9/5 + 32
		return convert, err
	} else {
		return convert, err
	}
}
