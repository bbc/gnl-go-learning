package bowling

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBowlingGame(t *testing.T) {
	Convey("GIVEN we have instantiated a new Game", t, func() {
		_ = &Game{}
	})
}
