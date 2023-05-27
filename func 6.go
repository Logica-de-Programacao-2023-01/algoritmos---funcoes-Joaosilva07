func concatenarComSeparador(slice []string) (string, error) {
    if len(slice) == 0 {
        return "", errors.New("O slice está vazio")
    }
    return strings.Join(slice, ", "), nil
}
