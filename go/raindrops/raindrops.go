package raindrops

import "strconv"


const testVersion = 3

func Convert(number int) string {
  returnString := ""

  if number % 3 == 0 {
    returnString += "Pling"
  }

  if number % 5 == 0 {
    returnString += "Plang"
  }

  if number % 7 == 0 {
    returnString += "Plong"
  }

  if number % 3 != 0 && number % 5 != 0 && number % 7 != 0 {
    returnString += strconv.FormatInt(int64(number), 10)
  }

  return returnString
}
