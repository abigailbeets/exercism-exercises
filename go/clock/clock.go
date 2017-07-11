package clock

import "fmt"

const testVersion = 4

const hrs = 24
const mins = 60

type Clock struct {
  minutes int
}

func New(hour, minute int) Clock {
  totalMinutes := (hour * 60 + minute) % 1440
  if totalMinutes < 0 {
    totalMinutes += 1440
  }

  return Clock{totalMinutes}
}

func (clock Clock) String() string {
  minutes := clock.minutes
  hours := 0
  hours = minutes / mins

  if hours == 24 {
    hours = 0
  }

  minutes %= mins

  return fmt.Sprintf("%02d:%02d", hours, minutes)
}

func (clock Clock) Add(minutes int) Clock {
  totalMinutes := clock.minutes + minutes
  hours := totalMinutes / 60
  minutes = totalMinutes % 60

  return New(hours, minutes)
}
