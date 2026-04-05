package main
import "fmt"



   // imprime "n² = ... = ?"

   // resultado = função(n-1)

   // calcula n²

   // imprime resultado final

   // retorna n²



func raizquadadra(n int)int{
    if n == 1{
        fmt.Println("1^2 =" ,n) 
        return 1
    }
    
   fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = ?\n",n,n-1,n-1)

    
    raizquadadra(n - 1)
    
  
    fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = %d\n",n,n-1,n-1,n*n)
    return 1
}
func main() {
    var n int 
    fmt.Scan(&n)
    

   
    raizquadadra(n)
}