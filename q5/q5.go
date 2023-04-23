package q5

func ConvertTemperature(temp float64, fromScale string, toScale string) (float64, error) {
    switch from {
    case Celsius:
        switch to {
        case F:
            return temp*9/5 + 32, nil
        case K:
            return temp + 273.15, nil
        default:
            return 0, fmt.Errorf("escala de destino inválida")
        }
    case Fahrenheit:
        switch to {
        case C:
            return (temp - 32) * 5 / 9, nil
        case K:
            return (temp-32)*5/9 + 273.15, nil
        default:
            return 0, fmt.Errorf("escala de destino inválida")
        }
    case K:
        switch to {
        case C:
            return temp - 273.15, nil
        case F:
            return (temp-273.15)*9/5 + 32, nil
        default:
            return 0, fmt.Errorf("escala de destino inválida")
        }
    default:
        return 0, fmt.Errorf("escala de origem inválida")
    }
}
