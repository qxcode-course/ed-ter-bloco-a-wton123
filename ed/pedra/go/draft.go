package main
import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    var ant, p int
    var ind = -1

    for i := 0; i < n; i++ {
        var a, b int
        fmt.Scan(&a, &b)

        if a < 10 || b < 10 {
            continue
        }

        if a < b {
            p = b - a
        } else {
            p = a - b
        }

        if ind == -1 {
            ant = p
            ind = i
        } else if p < ant {
            ant = p
            ind = i
        }
    }

    if ind == -1 {
        fmt.Println("sem ganhador")
    } else {
        fmt.Println(ind)
    }
}