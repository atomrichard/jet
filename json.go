package jet

import (
  "fmt"
  "io/ioutil"
)

//DoDaJSON is checking if a json file is existing and if makeit is true creating it
func DoDaJSON(makeit bool, filename string, path string, basepath string, basebyte []byte) (string, error) {

  filepathy := fmt.Sprintf("%s%s", path, filename)

  if(makeit) {
    //make the folder
    if err := EnsureDir(path); err != nil {
      fmt.Println("Directory creation failed with error: " + err.Error())
    }

  	if FileExists(filepathy) == false {
      var input []byte
      var err error

      if(basepath == ""){
        input = basebyte
      }else{
        input, err = ioutil.ReadFile(basepath)
        if err != nil {
          fmt.Println("ReadFile on basepath" + err.Error())
        }
      }

  		err = ioutil.WriteFile(filepathy, input, 0644)
  		if err != nil {
  			fmt.Println("WriteFile on input" + err.Error())
  		}
  	}
  }
  //megnyitni a fájlt
  file, err := ioutil.ReadFile(filepathy)
  if err != nil {
    fmt.Println("ReadFile on filepathy" + err.Error())
  }
  filestr := string(file)
  return filestr, err
}
