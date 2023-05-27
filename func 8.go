func filtrarPares(slice []int) ([]int, error) {
    if len(slice) == 0 {
        return nil, errors.New("O slice está vazio")
    }
    novoSlice := []int{}
    for _, num := range slice {
        if num%2 == 0 {
            novoSlice = append(novoSlice, num)
        }
    }
    return novoSlice, nil
}
