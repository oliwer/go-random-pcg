package randompcg

import "testing"
import "time"

func TestAll(t *testing.T) {
	now := uint64(time.Now().Unix())
	Srandom(now, now)
	v := BoundedRandom(255)
	//println(v)
	if v < 0 || v > 255 {
		t.Errorf("invalid BoundedRandom(255) result: %i", v)
	}
}
