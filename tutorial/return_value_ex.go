package main

func main()  {
	v, s := hello(10)
	println(v)
	println(s)
	// v := hello(10)
}

func hello(n int) (int, string)  {
	return n - 1, "asdf"
}
