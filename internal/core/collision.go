package core

func checkCollision(x1 int, y1 int, w1 int, h1 int, x2 int, y2 int, w2 int, h2 int) bool {
	return (max(x1, x2) < min(x1+w1, x2+w2)) && (max(y1, y2) < min(y1+h1, y2+h2))
}
