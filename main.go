package main

import (
	"fmt"

	packagea "github.com/SarahFrench/teamcity-example/internal/services/packageA"
	packageb "github.com/SarahFrench/teamcity-example/internal/services/packageB"
)

func main() {
	fmt.Println("Hello World")

	a := packagea.HelloWorld_PackageA()
	fmt.Println(a)
	b := packageb.HelloWorld_PackageB()
	fmt.Println(b)

	fmt.Println("Finished")
}
