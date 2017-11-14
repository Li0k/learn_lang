import (
	"fmt"
)

func fibonaci() fun() int {
	a,b := 0,1

	return func () int {
		a,b = b,a+b
		return a
	}
}

func fib(i int ) int {
	if i < 2 {
		return 1
	}
	
	return fib(i-1) + fib(i-2)
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(
			f(),
			fib(i),
		)
	}
	
}