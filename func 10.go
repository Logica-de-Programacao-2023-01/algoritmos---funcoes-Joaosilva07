func ordenarSlice(slice []int) ([]int, error) {
    if len(slice) == 0 {
        return nil, errors.New("O slice está vazio")
    }
    novoSlice := make([]int, len(slice))
    copy(novoSlice, slice)
    sort.Ints(novoSlice)
    return novoSlice, nil
}
