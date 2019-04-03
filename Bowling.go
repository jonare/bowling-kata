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
		var frameBonus = 0

		if i < 8 { 								//normal frames
			if isStrike(frame){
				frameBonus = game.frames[i+1].ball1 + game.frames[i+2].ball1
			} else if isSpare(frame){
				frameBonus = game.frames[i+1].ball1
			}
		} else if i == 8 {						//second to last
			if isStrike(frame) {
				frameBonus = game.frames[i+1].ball1 + game.bonus1
			} else if isSpare(frame) {
				frameBonus = game.frames[i+1].ball1
			}
		} else if i == 9 {						//last
			if isStrike(frame) {
				frameBonus = game.bonus1 + game.bonus2
			} else if isSpare(frame){
				frameBonus = game.bonus1
			}
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
