package bowling

import "testing"

func TestGame90(t *testing.T) {
	game := Game{
		frames: [12]Frame{
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
			Frame{0, 0},
			Frame{0, 0},
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
		 frames: [12]Frame{
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
			Frame{10, 0},
			Frame{10, 0},
		},
	}
	got := BowlingGame(game)
	want := 300

	if (got != want) {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}

func TestGame150(t *testing.T) {
	game := Game{
		frames: [12]Frame{
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
			Frame{5, 5},
			Frame{0, 0},
		},
	}
	got := BowlingGame(game)
	want := 150

	if (got != want) {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}

	func TestGame123(t *testing.T) {
		game := Game{
			frames: [12]Frame{
				Frame{10, 0}, 	//20
				Frame{5, 5},		//40
				Frame{10, 0},		//70
				Frame{10, 0},		//95
				Frame{10, 0},		//114
				Frame{5, 4},		//123
				Frame{0, 0},
				Frame{0, 0},
				Frame{0, 0},
				Frame{0, 0},
				Frame{0, 0},
				Frame{0, 0},
			},
		}
		got := BowlingGame(game)
		want := 123

		if (got != want) {
			t.Errorf("got '%d' want '%d'", got, want)
		}
}


