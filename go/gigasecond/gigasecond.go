/*

Add a gigasecond (1,000,000,000 seconds) to a given date/time

Usage:

now := time.Now()
nowPlusOneGigasecond := AddGigasecond(now)

*/
package gigasecond

import "time"

const GIGASECOND time.Duration = time.Second * 1000000000

// Add 1 gigasecond (1,000,000,000 seconds) to a given date/time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(GIGASECOND)
}
