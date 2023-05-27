func aplicarFuncaoESomar(slice []int, funcao func(int) int) (int, error) {
    if len(slice) == 0 {
        return 0, errors.New("O slice está vazio")
    }
    soma := 0
    for _, num := range slice {
        soma += funcao(num)
    }
    return soma, nil
}
