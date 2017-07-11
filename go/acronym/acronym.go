package acronym

import "strings"

const testVersion = 3

func Abbreviate(input string) string {
  resultString := ""
  input = strings.Replace(input, "-", " ", 10)
  inputArray := strings.Split(input, " ")

  for i := 0; i < len(inputArray); i++ {
    resultString += strings.ToUpper(inputArray[i][0:1])
  }

  return resultString
}
