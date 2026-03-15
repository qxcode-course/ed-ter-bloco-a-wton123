package main
import "fmt"
func main() {
    var nome string
    var ida int
    fmt.Scanf("%s",&nome)
    fmt.Scanf("%d",&ida)
    if ida < 12{
             fmt.Println(nome,"eh crianca")
    }else if ida < 18{
         fmt.Println(nome,"eh jovem")
    }else if ida < 65{
             fmt.Println(nome,"eh adulto")
    }else if ida < 1000{
             fmt.Println(nome,"eh idoso")
    }else{
         fmt.Println(nome,"eh mumia")
    }
   
}
