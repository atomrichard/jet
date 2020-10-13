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
//ParamCheck sends back the URL parameters value
func ParamCheck(param string, r *http.Request, fallback string) (string, bool) {
  keys, ok := r.URL.Query()[param]

      if !ok || len(keys[0]) < 1 {
          return fallback, false
      }
      return keys[0], true
}
