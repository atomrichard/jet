package jet

import (
  "os"
)

//FileExists is check if the file exists
func FileExists(filepath string) bool {
    info, err := os.Stat(filepath)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}

//EnsureDir is check and make folder
func EnsureDir(dirName string) error {

    err := os.Mkdir(dirName, 0755)

    if err == nil || os.IsExist(err) {
      return nil
    }
		return err
}
