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

  // ARRAY
  // var a [5]int
  // a[0] = 10
  // a[1] = 20
  // a[2] = 30

  // fmt.Println(a)
  // fmt.Println(len(a))

  // for i := 0; i < len(a); i++ {
  //   fmt.Println(a[i])
  // }

  // for i := range a {
  //   fmt.Println(a[i])
  // }

  // for _, v := range a {
  //   fmt.Println(v)
  // }

  // b := [3] int{1, 2, 3}
  // fmt.Println(b)

  // var c [2][3]int
  // for i := 0; i < len(c); i++ {
  //   for j := 0; j < len(c[i]); j++ {
  //     c[i][j] = j
  //   }
  // }

  // fmt.Println(c)

  // // array slide
  // d := make([]int, 5)
  // d[0] = 10
  // d[2] = 20

  // fmt.Println(d)

  // d = append(d, 90)
  // fmt.Println(d)

  // fmt.Println(d[2:4])

  // fmt.Println(d[0:4])
  // // equal
  // fmt.Println(d[:4])

  // f := d[2:4]
  // fmt.Println(f)

  // f[0] = 10
  // f[1] = 20

  // fmt.Println(d)
  // fmt.Println(f)

  // f = append(f, 20)
  // fmt.Println(d)
  // fmt.Println(f)

  // MAP
  a := make(map[string]string)
  a["hello"] = "gopher"
  a["name"] = "weerapon"
  a["x"] = "test"
  fmt.Println(a)

  // check key in array
  x, check := a["x"]

  if x, check := a["x"]; check {
    fmt.Println(x)
  }

  fmt.Println(check)
  fmt.Println(x)

  // delete key
  delete(a, "x")
  fmt.Println(a)

  for key, value := range a {
    fmt.Println(key, " : ", value)
  }

  b := map[string]string {
    "hello": "go",
    "go": "lang",
  }

  fmt.Println(b)
}