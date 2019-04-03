package bowling

type Game struct {
	frames [10]Frame
	bonus1 int
	bonus2 int
}

type Frame struct {
	ball1 int
	ball2 int
}

func BowlingGame (game Game) int {
	var score = 0
	for i, frame := range game.frames {
		if i < 8 {
			if frame.ball1 == 10 {
				score += game.frames[i+1].ball1
				score += game.frames[i+2].ball1
			} else if frame.ball1 + frame.ball2 == 10 {
				score += game.frames[i+1].ball1
			}
		} else if i == 8 {
			if frame.ball1 == 10 {
				score += game.frames[i+1].ball1
				score += game.bonus1
			} else if frame.ball1 + frame.ball2 == 10 {
				score += game.frames[i+1].ball1
			}
		} else if i == 9 {
			if frame.ball1 == 10 {
				score += game.bonus1
				score += game.bonus2
			} else if frame.ball1 + frame.ball2 == 10 {
				score += game.bonus1
			}
		}
		score += frame.ball1
		score += frame.ball2
	}
	return score
}
