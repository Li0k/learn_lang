package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

//实现Error()
func (e ErrNegativeSqrt) Error() string{

	// fmt.Sprintf(e) 打印e相当于调用 e.Error(),不转换会导致无限递归 栈溢出
	return fmt.Sprintf("cannot Sqrt negative number: %v",float64(e))
}

func Sqrt(x float64) (float64, error) {
	const E = 0.000001
	z := float64(1)
	k := float64(0)
	
	if x < 0 {
		return 0,ErrNegativeSqrt(x)	
	}
	
	for ; ;z = z - (z*z - x) / (2*z) {
		if z - k <= E && z-k >= -E {
			break
		}
		k = z
	}
	
	return z,nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
