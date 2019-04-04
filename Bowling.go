package bowling

type Game struct {
	frames [12]Frame
}

type Frame struct {
	ball1 int
	ball2 int
}

func BowlingGame (game Game) int {
	var score = 0
	for i := 0; i < 10; i++ {
		var frame = game.frames[i]
		var frameBonus = 0

		if isStrike(frame){
			frameBonus = game.frames[i+1].ball1
			if isStrike(game.frames[i+1]) {
				frameBonus += game.frames[i+2].ball1
			} else {
				frameBonus += game.frames[i+1].ball2
			}
		} else if isSpare(frame){
			frameBonus = game.frames[i+1].ball1
		}

		score += frame.ball1 + frame.ball2 + frameBonus
	}
	return score
}

func isStrike(frame Frame) bool {
	return frame.ball1 == 10
}

func isSpare(frame Frame) bool {
	return !isStrike(frame) && frame.ball1 + frame.ball2 == 10
}
