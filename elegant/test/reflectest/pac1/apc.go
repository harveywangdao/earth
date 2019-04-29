package pac1

import (
	"fmt"
)

type BBB struct {
	BB int
	bb int
}

func (b BBB) name() {
	fmt.Println("BBB name", b.BB, b.bb)
}

func (b BBB) Name() {
	fmt.Println("BBB Name", b.BB, b.bb)
}
