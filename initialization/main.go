package initialization

import "fmt"

type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	}
	return fmt.Sprintf("coucou")
}

func Main() {
	fmt.Println("KB =", KB)
	fmt.Println("MB =", MB)
	fmt.Println("GB =", GB)
	fmt.Println("YB =", YB)
}
