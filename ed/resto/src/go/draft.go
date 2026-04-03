package main
import "fmt"

func dividir(n1 int){
    
    if( n1 == 0){
        return 
    }
    resto := n1 % 2
    div := n1 / 2
    dividir(div)
    fmt.Println(div, resto)
}
func main() {
    var n int 
    fmt.Scan(&n)
    dividir(n)
    
}