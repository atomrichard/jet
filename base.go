package jet

import (
  "os"
)

func fileExists(filepath string) bool {
    info, err := os.Stat(filepath)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}

func ensureDir(dirName string) error {

    err := os.Mkdir(dirName, 0755)

    if err == nil || os.IsExist(err) {
      return nil
    }
		return err
}
