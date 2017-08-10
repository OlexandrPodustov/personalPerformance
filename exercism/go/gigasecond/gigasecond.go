package gigasecond

// import path for the time package from the standard library
import "time"

const testVersion = 4

func AddGigasecond(it time.Time) time.Time {
	return it.Add(time.Second * 1000 * 1000 * 1000)

}
