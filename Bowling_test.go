package bowling

import "testing"

func TestGame90(t *testing.T) {
	game := Game{
		frames: [12]Frame{
			{9, 0},
			{9, 0},
			{9, 0},
			{9, 0},
			{9, 0},
			{9, 0},
			{9, 0},
			{9, 0},
			{9, 0},
			{9, 0},
			{0, 0},
			{0, 0},
		},
	}
	got := BowlingGame(game)
	want := 90

	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}

func TestGame300(t *testing.T) {
	game := Game{
		 frames: [12]Frame{
			 {10, 0},
			 {10, 0},
			 {10, 0},
			 {10, 0},
			 {10, 0},
			 {10, 0},
			 {10, 0},
			 {10, 0},
			 {10, 0},
			 {10, 0},
			 {10, 0},
			 {10, 0},
		},
	}
	got := BowlingGame(game)
	want := 300

	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}

func TestGame150(t *testing.T) {
	game := Game{
		frames: [12]Frame{
			{5, 5},
			{5, 5},
			{5, 5},
			{5, 5},
			{5, 5},
			{5, 5},
			{5, 5},
			{5, 5},
			{5, 5},
			{5, 5},
			{5, 5},
			{0, 0},
		},
	}
	got := BowlingGame(game)
	want := 150

	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}

	func TestGame123(t *testing.T) {
		game := Game{
			frames: [12]Frame{
				{10, 0}, //20
				{5, 5},  //40
				{10, 0}, //70
				{10, 0}, //95
				{10, 0}, //114
				{5, 4},  //123
				{0, 0},
				{0, 0},
				{0, 0},
				{0, 0},
				{0, 0},
				{0, 0},
			},
		}
		got := BowlingGame(game)
		want := 123

		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
}


