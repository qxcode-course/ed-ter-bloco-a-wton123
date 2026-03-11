package main
import "fmt"
func main() {
    fmt.Println("faltando 7 dias")
    var n int;
    fmt.Println("wellington digite 1 para verdadeiro e 2 para falso:")
    fmt.Println("estar ansioso?")
    fmt.Scan(&n)
    if n == 1{
        fmt.Println("fique tranquilo nn e so vc")
    }else{
        fmt.Println("isso meu garoto")
    }
    fmt.Println("pensou numa musica pra escutarem?")
    
        // var musicas [5]string := [5]string{"dg", "ffg", "etr", "fgf", "fgh"}
         //var musicas [5]string = [5]string{"dg","ffg","etr", "fgf","fgh"}
        //musicas := {"nosso","declaração","","","","","","","","","",""}
        musicas := [5]string{"dg", "ffg", "etr", "fgf", "fgh"}
        for i := 0 ; i > 4 ; i++ {
            
        }
        fmt.Println(musicas)  
}
