package randompcg

type PCGState struct {
	State uint64
	Inc uint64
}

var globalState *PCGState = &PCGState{0x853c49e6748fea9b, 0xda3e39cb94b95bdb }

func Srandom(initState, initSeq uint64) {
	globalState.State = 0
	globalState.Inc = (initSeq << 1) | 1
	Random()
	globalState.State += initState
	Random()
}

func Random() uint32 {
    oldState := globalState.State
    globalState.State = oldState * 6364136223846793005 + globalState.Inc
    xorShifted := uint32(((oldState >> 18) ^ oldState) >> 27)
    rot := uint32(oldState >> 59)
    return (xorShifted >> rot) | (xorShifted << ((-rot) & 31))
}

func BoundedRandom(bound uint32) uint32 {
    threshold := -bound % bound
    var r uint32
    for {
	r = Random()
	if r >= threshold {
	    return r % bound
	}
    }
}
