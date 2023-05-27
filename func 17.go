func ordenarSliceStrings(slice []string) ([]string, error) {
    if len(slice) == 0 {
        return nil, errors.New("O slice está vazio")
    }
    novoSlice := make([]string, len(slice))
    copy(novoSlice, slice)
    sort.Strings(novoSlice)
    return novoSlice, nil
}
