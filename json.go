package jet

import (
  "fmt"
  "io/ioutil"
)

//DoDaJSON is checking if a json file is existing and if makeit is true creating it
func DoDaJSON(makeit bool, filename string, path string, defaultpath string) (string) {

  filepathy := fmt.Sprintf("%s%s", path, filename)

  if(makeit) {
    //make the folder
    if err := EnsureDir(path); err != nil {
      fmt.Println("Directory creation failed with error: " + err.Error())
    }

  	if FileExists(filepathy) == false {

  		input, err := ioutil.ReadFile(defaultpath)
  		if err != nil {
  			fmt.Println(err)
  		}

  		err = ioutil.WriteFile(filepathy, input, 0644)
  		if err != nil {
  			fmt.Println("Error creating")
  			fmt.Println(err)
  		}
  	}
  }
  //megnyitni a fájlt
  file, err := ioutil.ReadFile(filepathy)
  if err != nil {
    fmt.Println(err)
  }
  filestr := string(file)
  return filestr
}
