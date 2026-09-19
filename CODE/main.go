package main

import (
	w "CODE/WORLDMODEL/V1"
	"fmt"
)

func main() {
	data := w.DataJSON{
		Running: true,
		Visible: false,
		Counter: 42,
		Score:   100,
	}

	w.SaveData(data)

	fmt.Println("A")
}
