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

  // IF - ELSE
  // fmt.Print("Input a fruit: ")
  // var fruit string
  // fmt.Scanln(&fruit)

  // if fruit == "" {
  //   fmt.Println("meh! ")
  //   return
  // }

  // switch fruit {
  //   case "apple":
  //     fmt.Println("Apple")
  //   case "banana":
  //     fmt.Println("Banana")
  //   case "lemon":
  //     fmt.Println("Lemon")
  //   default:
  //     fmt.Println("Default")
  // }

  // fmt.Println("Input Score: ")
  // var score int
  // fmt.Scanln(&score)

  // if score < 50 {
  //   fmt.Println("F")
  // } else if score < 60 {
  //   fmt.Println("D")
  // } else if score < 70 {
  //   fmt.Println("C")
  // } else if score < 80 {
  //   fmt.Println("B")
  // } else {
  //   fmt.Println("A")
  // }

  // switch {
  //   case score < 50:
  //     fmt.Println("F")
  //   case score < 60:
  //     fmt.Println("D")
  //   case score < 70:
  //     fmt.Println("C")
  //   case score < 80:
  //     fmt.Println("B")
  //   default:
  //     fmt.Println("A")
  // }

  var a [5]int
  a[0] = 10
  a[1] = 20
  a[2] = 30

  fmt.Println(a)
  fmt.Println(len(a))

  for i := 0; i < len(a); i++ {
    fmt.Println(a[i])
  }
}