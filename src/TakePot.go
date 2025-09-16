package main

import "fmt"

func TakePot(c Character) {
	if c.Pv+50 > c.Pvmax {
		c.Pv = c.Pvmax
	} else {
		c.Pv += 50
	}
	fmt.Println(c.Pv, "HP /", c.Pvmax, "HP")
}
