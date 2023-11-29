package find

type Sam struct {
	J int 
}

func Odd(a int) bool {
	if a%2 == 0 {
		return true
	} else {
		return false
	}
}

func (d *Sam)Add(a, b int) int {
	d.J=a + b + d.J
	return d.J
}
