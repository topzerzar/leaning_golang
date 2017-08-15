package main

import (
	"fmt"
)

func main() {
  // print
  // fmt.Println("Hello, Gopher.")

  // Declare variable
  // var gopher string
  // gopher = "Gopher"
  // fmt.Printf("Hello, %s.\n", gopher)

  // var name = "weerapon top"
  // fmt.Printf("My name is %s. \n", name)

  // nickname := "TOPz"
  // fmt.Printf("My nickname is %s. \n", nickname)
  
  // fix compile error case declare vaialble not used 
  // _ not keep value
  // _ = nickname

  fmt.Print("Input a fruit: ")
  var fruit string
  fmt.Scanln(&fruit)

  if fruit == "" {
    fmt.Println("meh! ")
    return
  }
}