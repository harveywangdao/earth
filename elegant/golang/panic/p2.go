package main

func do1() {
	defer func() {
		if r := recover(); r != nil {
			_ = r
		}
	}()
	panic("12345")
}

func main() {
	do1()
}
