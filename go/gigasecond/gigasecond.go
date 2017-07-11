package gigasecond

import "time"

const testVersion = 4
const gigaSecond = "1000000000s"

func AddGigasecond(birthday time.Time) time.Time {
  duration, _ := time.ParseDuration(gigaSecond)
  return birthday.Add(duration)
}
