func filtrarStringsMaiusculas(slice []string) ([]string, error) {
    if len(slice) == 0 {
        return nil, errors.New("O slice está vazio")
    }
    novoSlice := []string{}
    for _, str := range slice {
        if len(str) > 0 && unicode.IsUpper(rune(str[0])) {
            novoSlice = append(novoSlice, str)
        }
    }
    return novoSlice, nil
}
