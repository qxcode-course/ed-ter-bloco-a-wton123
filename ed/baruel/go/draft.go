package main

import "fmt"

func main() {

    var total, qtd int
    fmt.Scan(&total)
    fmt.Scan(&qtd)  

    contador := make(map[int]int) 
    
    for i := 0; i < qtd; i++ {
        var v int
        fmt.Scan(&v)
        contador[v]++ 
    }

    repetidas := []int{}
    for fig, qtd := range contador {
        if qtd > 1 {
            for i := 0; i < qtd-1; i++ {
                repetidas = append(repetidas, fig)
            }
        }
    }

   
    faltando := []int{}

    for i := 1; i <= total; i++ {
        if contador[i] == 0 {
           
            faltando = append(faltando, i)
        }
    }

   
    if len(repetidas) == 0 {
        fmt.Println("N")
    } else {
        for i, v := range repetidas {
            if i > 0 {
                fmt.Print(" ")
            }
            fmt.Print(v)
        }
        fmt.Println()
    }

   
    if len(faltando) == 0 {
        fmt.Println("N")
    } else {
        for i, v := range faltando {
            if i > 0 {
                fmt.Print(" ")
            }
            fmt.Print(v)
        }
        fmt.Println()
    }
}