package main
import "fmt"
func acharmatar(v map[int]int , n int, a int){

i := a 
    for  i < n{
        v[i + 1] = 0
        a = i + 2
        /*if v[i] > 0 {
             v[i] = 0

        }*/
        i++
    }
} 
func main() {
    t := 0
    a := 0
    n := make(map[int]int)
    
    fmt.Scan(&t)
    for i := 0 ; i < t ; i++{
        n[i] = i + 1
    }
     fmt.Scan(&a)
    acharmatar(n, t, a)
    fmt.Println(n, a)
}