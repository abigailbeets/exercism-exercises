package hamming

import "errors"

const testVersion = 6

func Distance(a, b string) (int, error) {
  //error case
  if len(a) != len(b) {
    return 0, errors.New("Strands must be the same length")
  }

  //case for no error
  count := 0
  for i := 0; i < len(a); i++ {
    if a[i] != b[i] {
      count ++
    }
  }

  return count, nil
}
