package buffer

type Location struct {
	X, Y int
}

type Cursor struct {
	Location
	Selection [2]Location
}
