func calcularMedia(slice []int) float64 {
    total := 0
    for _, num := range slice {
        total += num
    }
    return float64(total) / float64(len(slice))
}
