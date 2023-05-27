func encontrarSegundoMaiorValor(slice []int) (int, error) {
    if len(slice) < 2 {
        return 0, errors.New("O slice precisa ter pelo menos 2 elementos")
    }
    max := math.MinInt64
    segundoMax := math.MinInt64
    for _, num := range slice {
        if num > max {
            segundoMax = max
            max = num
        } else if num > segundoMax && num < max {
            segundoMax = num
        }
    }
    return segundoMax, nil
}
