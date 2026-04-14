package main
import "fmt"

func casal(vet[] int , a int){
    var qdt int  
    semninguem :=  make([]int, a)
    for i := range vet{
        if semninguem[i] < 0 {
        
        }
    }
    fmt.Println(vet)
    
}
func main() {
    var a int
    fmt.Scan(&a)
    vet := make([]int, a)
    for i := range vet{
        var t int
        fmt.Scan(&t)
        vet[i] = t

    }
    casal(vet , a)
    
}
