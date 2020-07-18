package main

const (
	a = iota
	b = iota
	c = iota
)

const (
	a2 = iota
)

func naming() {
	const myConstant int = 1234
	printValAndType(myConstant, myConstant)
}

func printValAndTypeForSingleVal(e interface{}) {
	printValAndType(e, e)
}

func iotas() {
	printValAndType(a, a)
	printValAndType(b, b)
	printValAndType(c, c)
}

func main() {
	println("____CONSTANTS_____")
	naming()
	iotas()
}
