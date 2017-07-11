package bob

import (
  "strings"
  "unicode"
)

const testVersion = 3

func Hey(input string) string {
  input = strings.TrimSpace(input)

  switch {
  case input == "":
    return "Fine. Be that way!"
  case any(input, unicode.IsUpper) && !any(input, unicode.IsLower):
    return "Whoa, chill out!"
  case input[len(input)-1] == '?':
    return "Sure."
  default:
    return "Whatever."
  }
}

func any(input string, myFunc func(rune) bool) bool {
  for _, item := range input {
    if myFunc(item) {
      return true
    }
  }
  return false
}
