package pangram

import "strings"

const testVersion = 1

func IsPangram(input string) bool {
  if len(input) < 26 {
    return false
  }

  alphabet := strings.Split("abcdefghijklmnopqrstuvwxyz", "")
  input = strings.ToLower(input)

  for i := 0; i < len(alphabet); i++ {
    if !strings.Contains(input, alphabet[i]) {
      return false
    }
  }

  return true
}
