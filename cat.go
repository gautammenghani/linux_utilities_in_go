package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "path/filepath"
    "errors"
    "strings"
    "strconv"
)

func NumberNonBlankLines(contents string) (string) {
  lines := strings.Split(contents, "\n")
  cntr := 1
  var modifiedContents string
  for _, line := range lines {
    if(len(line)>0) {
      modifiedContents += (strconv.Itoa(cntr) + " " + line + "\n")
      cntr += 1
     } else {
        modifiedContents += "\n"
     }
  }
  return modifiedContents
}

func ReadFile(fileNames []string, switches []string) (string, error){
    var contents string
    for _, file := range fileNames {
      //Get absolute file path
      fileName, err := filepath.Abs(file)
      if (err != nil) {
          errMsg := fmt.Sprintf ("File not found: %s\n", err)
          fmt.Println (errMsg)
          os.Exit(0)
      }
      data, err := ioutil.ReadFile(fileName)
      if (err != nil) {
          return contents, errors.New(fmt.Sprintf("Could not read file : %s\n", err))
      } else {
          contents += "\n" + "[+] Reading file : "+fileName + "\n" + string(data)
      }
    }
    //modify the contents as per the switches
    for _, param := range switches {
      switch param{
        case "b": 
          contents = NumberNonBlankLines(contents)
        default: 
        fmt.Println(param + " switch is not supported yet")
      }
    }
    return contents, nil
}

func ParseCommand(userInput []string) (string, error) {
  var switches[]string
  var files[]string
  for _,elem := range userInput {
    if(string(elem[0])=="-") {
      //get the switches
      t := strings.Split(elem[1:], "")
      for _,st := range t {
        switches = append(switches, st)
      }
    } else {
      //get the file names
      files = append(files,elem)
    }
  }
  contents, err := ReadFile(files, switches)
  if (err != nil) {
      fmt.Println (contents)
      errMsg := fmt.Sprintf ("Error occurred when reading file: %s\n", err)
      fmt.Println (errMsg)
      os.Exit(0)
  } 
  return contents, nil  
}

func main() {
    // Std input
    if (len(os.Args) == 1 || os.Args[1]=="-") {
        var input string
        for {
            fmt.Scanln(&input)
            fmt.Println(input)
        } 
    } else {
      contents, err := ParseCommand(os.Args[1:])
      if (err != nil) {
          fmt.Println (contents)
          errMsg := fmt.Sprintf ("Error occurred when reading file: %s\n", err)
          fmt.Println (errMsg)
          os.Exit(0)
      }
      io.WriteString(os.Stdout, contents)    
    }
}
