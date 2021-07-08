package main

import ("fmt"
        "time"
        "math/rand"
        "strings"
        "os"
)

func main(){
  // Generate hash
  var hash = randomStr(64)

  // Save hash if needed
  //saveString(hash)

  // Print timestamp and hash
  for range time.Tick(time.Second * 5) {
    currentTime := time.Now()
    fmt.Println(currentTime.Format("2006-01-02T01:04:05.705Z") + " : " + hash)
  }
}

// Source code here: https://yourbasic.org/golang/generate-random-string/
func randomStr(length int) string {
  rand.Seed(time.Now().UnixNano())
  chars := []rune("!#€%&/()=?$|[]abcdefghijklmnopqrstuvwxyzåäöABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
    "abcdefghijklmnopqrstuvwxyzåäö0123456789")
  var b strings.Builder
  for i := 0; i < length; i++ {
    b.WriteRune(chars[rand.Intn(len(chars))])
  }
  str := b.String()
	return string(str)
}

func saveString(str string){
  f, err := os.Create("data.txt")
  defer f.Close()
  if err != nil {
          fmt.Println(err)
      }

  f.WriteString(str)
}
