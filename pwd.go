package main
import (
    "fmt"
    "path/filepath"
)

func main() {
    path, err:= filepath.Abs(".")
    if (err != nil){
        fmt.Println("Could not determine current dir : ", err)
    } else {
        fmt.Println(path) 
    }
}
