func aplicarFuncao(slice []int, função func(int) int) ([]int, error) {
    if len(slice) == 0 {
        return nil, errors.New("O slice está vazio")
    }
    novoSlice := make([]int, len(slice))
    for i, num := range slice {
        novoSlice[i] = função(num)
    }
    return novoSlice, nil
}
