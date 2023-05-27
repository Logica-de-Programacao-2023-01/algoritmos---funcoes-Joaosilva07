func filtrarStringsMaisDe5Caracteres(slice []string) ([]string, error) {
    if len(slice) == 0 {
        return nil, errors.New("O slice está vazio")
    }
    novoSlice := []string{}
    for _, str := range slice {
        if len(str) > 5 {
            novoSlice = append(novoSlice, str)
        }
    }
    return novoSlice, nil
}
