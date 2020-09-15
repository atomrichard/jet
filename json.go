package jet

import (
  "fmt"
  "io/ioutil"
  "os"
	"path/filepath"
	"strings"
)

//DoDaJSON is checking if a json file is existing and if makeit is true creating it
func DoDaJSON(makeit bool, filename string, path string, defaultpath string) (string) {

  if(makeit) {
    //make the folder
    if err := ensureDir(path); err != nil {
      fmt.Println("Directory creation failed with error: " + err.Error())
    }

    filepath := fmt.Sprintf("%s/%s.json", path, filename)

  	if fileExists(filepath) == false {

  		input, err := ioutil.ReadFile(defaultpath)
  		if err != nil {
  			fmt.Println(err)
  		}

  		err = ioutil.WriteFile(filepath, input, 0644)
  		if err != nil {
  			fmt.Println("Error creating")
  			fmt.Println(err)
  		}
  	}
  }
  //megnyitni a fájlt
  file, err := ioutil.ReadFile(filepath)
  if err != nil {
    fmt.Println(err)
  }
  filestr := string(file)
  return filestr
}
