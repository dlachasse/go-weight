package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
  "./lib"
)

// this is a comment

func main() {
  var myunit string
  var amount string

  reader := bufio.NewReader(os.Stdin)

  fmt.Print("Enter input unit: ")
  myunit, _ = reader.ReadString('\n')
  myunit = myunit[:len(myunit)-1]

  fmt.Print("Enter input amount: ")
  amount, _ = reader.ReadString('\n')
  amount = amount[:len(amount)-1]
  amountFlt, err := strconv.ParseFloat(amount, 64)

  if err != nil {
    fmt.Println(err)
  }

  fmt.Print("Enter output unit: ")
  outUnit, _ := reader.ReadString('\n')
  outUnit = outUnit[:len(outUnit)-1]

  wgt := weight.Weight{Unit: myunit, Count: amountFlt}
  fmt.Println(weight.ConvertTo(wgt, outUnit))
}
