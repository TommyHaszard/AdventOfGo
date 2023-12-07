package main

type Data struct {
	distance int
	time     int
}

func main() {
	problem1()
	problem2()
}

func (s *Data) UpdateDistance(distance int) {
	s.distance = distance
}
