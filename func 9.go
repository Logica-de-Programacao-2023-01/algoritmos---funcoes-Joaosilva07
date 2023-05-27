func obterPalavras(str string) ([]string, error) {
    if len(str) == 0 {
        return nil, errors.New("A string está vazia")
    }
    palavras := strings.Fields(str)
    return palavras, nil
}
