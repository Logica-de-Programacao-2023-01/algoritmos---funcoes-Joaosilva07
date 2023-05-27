func contarVogais(str string) int {
    vogais := []rune{'a', 'e', 'i', 'o', 'u'}
    count := 0
    for _, char := range strings.ToLower(str) {
        for _, vogal := range vogais {
            if char == vogal {
                count++
                break
            }
        }
    }
    return count
}
