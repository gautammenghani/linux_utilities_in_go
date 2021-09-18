package main 

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

func ls (args []string) (string, error) {
    files, err := ioutil.ReadDir(".")
    if (err!=nil) {
        log.Fatal(err)
    } 
    for _, file := range files {
        if (file.IsDir()) { 
            fmt.Println("Dire : ", file.Name())
        } else {
            fmt.Println("File : ", file.Name())
        }
    }
    return "", nil
}

func main(){
    files, err := ls(os.Args[1:])
    if (err!=nil){
        log.Fatal(err)
    } 
    fmt.Println(files)
}
