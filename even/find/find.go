package find

type Sam struct {
	j int 
}

func Odd(a int) bool {
	if a%2 == 0 {
		return true
	} else {
		return false
	}
}

func (d *Sam)Add(a, b int) int {
	d.j=a + b + d.j
	return d.j
}
