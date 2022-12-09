package bowling

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBowlingGame(t *testing.T) {
	Convey("GIVEN we have instantiated a new Game", t, func() {
		g := &Game{}

		Convey("WHEN we throw a gutter game", func() {
			rollMany(g, 20, 0)

			Convey("THEN the score is zero.", func() {
				So(g.Score(), ShouldEqual, 0)
			})
		})

		Convey("WHEN we throw all ones", func() {
			rollMany(g, 20, 1)

			Convey("THEN the score is twenty.", func() {
				So(g.Score(), ShouldEqual, 20)
			})
		})

		Convey("WHEN we throw a Spare followed by 3", func() {
			rollSpare(g)
			g.Roll(3)
			rollMany(g, 17, 0)

			Convey("THEN the score is sixteen.", func() {
				So(g.Score(), ShouldEqual, 16)
			})
		})

		Convey("WHEN we throw a Strike followed by 3 and 4", func() {
			rollStrike(g)
			g.Roll(3)
			g.Roll(4)

			rollMany(g, 16, 0)

			Convey("THEN the score is twenty-four.", func() {
				So(g.Score(), ShouldEqual, 24)
			})
		})

		Convey("WHEN we throw a Perfect game", func() {
			rollMany(g, 12, 10)

			Convey("THEN the score is three hundred.", func() {
				So(g.Score(), ShouldEqual, 300)
			})
		})
	})
}

func rollStrike(g *Game) {
	g.Roll(10)
}

func rollSpare(g *Game) {
	g.Roll(5)
	g.Roll(5) // Spare
}

func rollMany(g *Game, rolls int, pins int) {
	for i := 0; i < rolls; i++ {
		g.Roll(pins)
	}
}
