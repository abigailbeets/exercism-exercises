package accumulate

const testVersion = 1

func Accumulate(input []string, myFunc func(string) string) []string {
  var result []string

  for _, value := range input {
    result = append(result, myFunc(value))
  }

  return result
}
