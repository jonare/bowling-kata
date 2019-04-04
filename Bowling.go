package bowling

type Game struct {
	frames [12]Frame
}

type Frame struct {
	ball1 int
	ball2 int
}

func (frame *Frame) isStrike() bool {							//"attach methods to types ~= objektorientering
	return frame.ball1 == 10
}

func (frame *Frame) isSpare() bool {
	return !frame.isStrike() && frame.ball1 + frame.ball2 == 10
}

func BowlingGame (game Game) int {
	var score = 0
	for i := 0; i < 10; i++ {
		var frame = game.frames[i]
		var frameBonus = 0

		if frame.isStrike(){
			frameBonus = game.frames[i+1].ball1
			if game.frames[i+1].isStrike(){
				frameBonus += game.frames[i+2].ball1
			} else {
				frameBonus += game.frames[i+1].ball2
			}
		} else if frame.isSpare(){
			frameBonus = game.frames[i+1].ball1
		}

		score += frame.ball1 + frame.ball2 + frameBonus
	}
	return score
}