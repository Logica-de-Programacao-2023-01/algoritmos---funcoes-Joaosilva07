func encontrarPosicaoElemento(slice []int, valor int) int {
    for i, num := range slice {
        if num == valor {
            return i
        }
    }
    return -1
}
