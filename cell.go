package main

const (
	dead3 = 3 << 1
	live2 = (3 << 1) | 1
	live3 = (4 << 1) | 1
)

func isAlive(cell int8) bool {
	return (cell & 1) != 0
}

func neighbourCount(cell int8) int8 {
	return cell &^ 1
}

func changeState(cell *int8) {
	*cell ^= 1
}

func isChanging(cell int8) bool {
	return cell == dead3 ||
		(isAlive(cell) && (cell < live2 || cell > live3))
}
