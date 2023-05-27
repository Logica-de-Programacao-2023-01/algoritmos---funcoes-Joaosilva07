func obterIntersecao(slice1, slice2 []int) ([]int, error) {
    if len(slice1) == 0 || len(slice2) == 0 {
        return nil, errors.New("Pelo menos um dos slices está vazio")
    }
    intersecao := []int{}
    mapa := make(map[int]bool)
    for _, num := range slice1 {
        mapa[num] = true
    }
    for _, num := range slice2 {
        if mapa[num] {
            intersecao = append(intersecao, num)
        }
    }
    return intersecao, nil
}
