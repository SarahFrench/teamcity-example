package main

import (
	"fmt"

	packagea "github.com/SarahFrench/teamcity-example/internal/services/packageA"
	packageb "github.com/SarahFrench/teamcity-example/internal/services/packageB"
	packagec "github.com/SarahFrench/teamcity-example/internal/services/packageC"
)

func main() {
	fmt.Println("Hello World")

	a := packagea.HelloWorld_PackageA()
	fmt.Println(a)
	b := packageb.HelloWorld_PackageB()
	fmt.Println(b)
	c := packagec.HelloWorld_PackageC()
	fmt.Println(c)

	fmt.Println("Finished")
}
