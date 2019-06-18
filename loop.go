package main
import (
    "fmt"
    "math"
)

func sqrt(x float64) string{
    if x < 0 {
	//平方根算法與正數相同，差別在一個虛數i
	return sqrt(-x)+"i"
    }
    return fmt.Sprint(math.Sqrt(x))
}

func main(){
    sum := 0
    for i := 0; i < 10 ; i++{
	sum += i
    }

    fmt.Println(sum)
    sum2 := 1
    for ;sum2 < 1000;{
	sum2 += sum2
    }

    fmt.Println(sum2)
    fmt.Println(sqrt(2), sqrt(-4))

}



