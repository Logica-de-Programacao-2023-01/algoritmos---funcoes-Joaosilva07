func verificarNumeroPrimo(numero int) (bool, error) {
    if numero < 0 {
        return false, errors.New("O número deve ser positivo")
    }
    if numero < 2 {
        return false, nil
    }
    for i := 2; i*i <= numero; i++ {
        if numero%i == 0 {
            return false, nil
        }
    }
    return true, nil
}
