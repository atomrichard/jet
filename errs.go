package jet

import (
  "fmt"
  "os"
)

//HardErr log error then Stop the program
func HardErr(err error) {
    fmt.Println(err)
    os.Exit(2)
}

//SoftErr log error then do nothing
func SoftErr(err error) {
    fmt.Println(err)
}
