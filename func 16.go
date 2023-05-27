func substituirVogais(str string) (string, error) {
    if len(str) == 0 {
        return "", errors.New("A string está vazia")
    }
    vogais := "aeiouAEIOU"
    novaStr := ""
    for _, char := range str {
        if strings.ContainsRune(vogais, char) {
            novaStr += "*"
        } else {
            novaStr += string(char)
        }
    }
    return novaStr, nil
}
