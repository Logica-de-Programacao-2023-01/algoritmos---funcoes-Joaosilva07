func verificarPresencaNumero(slice []int, numero int) (bool, error) {
    if len(slice) == 0 {
        return false, errors.New("O slice está vazio")
    }
    for _, num := range slice {
        if num == numero {
            return true, nil
        }
    }
