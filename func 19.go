func obterNumerosPrimos(numero int) ([]int, error) {
    if numero < 0 {
        return nil, errors.New("O número deve ser positivo")
    }
    primos := []int{}
    for i := 2; i <= numero; i++ {
        if verificarNumeroPrimo(i) {
            primos = append(primos, i)
        }
    }
    return primos, nil
}
