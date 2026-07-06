package main
import "fmt"

func backtracking(nums []int, idx int, somaa int, alvo int)bool{
   
    if somaa == alvo {
        return true
    }
     if idx == len(nums) {
        return false
    }
    return backtracking(nums, idx+1,somaa + nums[idx],alvo) || backtracking(nums, idx+1,somaa,alvo)
   
}
func main() {
    var n , m int 
    fmt.Scan(&n,&m)
    nums := make([]int, n)
    for i := 0 ; i < n ; i++{
        fmt.Scan(&nums[i])
    }
    fmt.Println(backtracking(nums, 0, 0, m))
}