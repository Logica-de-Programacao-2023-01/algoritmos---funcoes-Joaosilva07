func calcularSomaDigitos(numero int) (int, error) {
    if numero < 0 {
        return 0, errors.New("O número deve ser positivo")
    }
    soma := 0
    for numero > 0 {
        soma += numero % 10
        numero /= 10
    }
    return soma, nil
}
