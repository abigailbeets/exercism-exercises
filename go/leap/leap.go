package leap

const testVersion = 3

func IsLeapYear(x int) bool {
  if x % 4 == 0 {
    if x % 100 == 0 {
      if x % 400 == 0 {
        return true
      }
      return false
    }
    return true
  }
  return false
}
