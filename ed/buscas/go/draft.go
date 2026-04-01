package main
import "fmt"
func cacapalavra(strings []string, queries []string) []int {
    freq := make(map[string]int)

    // contar ocorrências
    for _, s := range strings {
        freq[s]++
    }

    // responder consultas
    var result []int
    for _, q := range queries {
        result = append(result, freq[q])
    }

    return result
}


func main() {
    var n1, n2 int
      fmt.Scan(&n1)
   vet1 := make([]string, n1)
  
  
    for i := 0 ; i < n1 ; i++{
        fmt.Scan(&vet1[i])
    }
    fmt.Scan(&n2)
    vet2 := make([]string, n2)
    for i := 0 ; i < n2; i++{
        fmt.Scan(&vet2[i])
    }
    cacapalavra(vet1 , vet2)
    result := cacapalavra(vet1, vet2)

	// imprimir resultado
	for i, v := range result {
    if i > 0 {
        fmt.Print(" ")
    }
    fmt.Print(v)
}
	
	fmt.Println()
}
