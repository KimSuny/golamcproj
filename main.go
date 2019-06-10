package main

import (
	"fmt"
	"lamcproj/lamcproj"
)

func main() {
	//Longitude, Latitude
	//Latitude : 37.479884
	//Longitude : 126.9863715
	x, y := lamcproj.Map_conv(126.9863715, 37.4764771)

	fmt.Println(x, y)
}
