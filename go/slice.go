func Pic(dx,dy int) [][]uint8 {
	a := make([][]uint8,dy)

	for i := range a{
		b := make([]uint8,dx)
		for j := range b{
			b[j] = uint8(i^j)
		}

		a[i] = b
 	}

 	return a
}