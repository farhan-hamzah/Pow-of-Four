package main
import (
	"fmt"
	"math"
)
func main(){
	var n, x float64
	var b bool
	var i int
	i = 0
	fmt.Scan(&n)
	for x <= n&&!b{
		x = math.Pow(4, float64(i))
		if x == n {
			b = true
		}
		i++
	}
	fmt.Print(b)
}