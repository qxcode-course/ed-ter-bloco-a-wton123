package main

import "fmt"

func main() {

    var q int
    var d string

    fmt.Scan(&q, &d)

   
    x := make([]int, q)
    y := make([]int, q)


    for i := 0; i < q; i++ {
        fmt.Scan(&x[i], &y[i])
    }

    
    antX := x[0]
    antY := y[0]

    
    if d == "L" {
        x[0]--
    } else if d == "R" {
        x[0]++
    } else if d == "U" {
        y[0]-- 
    } else if d == "D" {
        y[0]++ 
    }

  
    for i := 1; i < q; i++ {

       
        tempX := x[i]
        tempY := y[i]

        
        x[i] = antX
        y[i] = antY

        
        antX = tempX
        antY = tempY
    }

  
    for i := 0; i < q; i++ {
        fmt.Println(x[i, y[] , )
    }
}