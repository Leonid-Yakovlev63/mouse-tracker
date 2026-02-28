package domain

import "fmt"

type (
	Cords struct {
		X int `json:"x"`
		Y int `json:"y"`
	}
)

func (c *Cords) PrintCords() {
	fmt.Println("x:", c.X, "y:", c.Y)
}
