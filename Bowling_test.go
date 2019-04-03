package bowling

import "testing"

func TestGame90(t *testing.T) {
	game := Game{
		frames: [10]Frame{
			Frame{9, 0},
			Frame{9, 0},
			Frame{9, 0},
			Frame{9, 0},
			Frame{9, 0},
			Frame{9, 0},
			Frame{9, 0},
			Frame{9, 0},
			Frame{9, 0},
			Frame{9, 0},
		},
	}
	got := BowlingGame(game)
	want := 90

	if (got != want) {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}

func TestGame300(t *testing.T) {
	game := Game{
		 frames: [10]Frame{
			Frame{10, 0},
			Frame{10, 0},
			Frame{10, 0},
			Frame{10, 0},
			Frame{10, 0},
			Frame{10, 0},
			Frame{10, 0},
			Frame{10, 0},
			Frame{10, 0},
			Frame{10, 0},
		},
		bonus1: 10,
		bonus2: 10,
	}
	got := BowlingGame(game)
	want := 300

	if (got != want) {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}

func TestGame150(t *testing.T) {
	game := Game{
		frames: [10]Frame{
			Frame{5, 5},
			Frame{5, 5},
			Frame{5, 5},
			Frame{5, 5},
			Frame{5, 5},
			Frame{5, 5},
			Frame{5, 5},
			Frame{5, 5},
			Frame{5, 5},
			Frame{5, 5},
		},
		bonus1: 5,
		bonus2: 0,
	}
	got := BowlingGame(game)
	want := 150

	if (got != want) {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}

