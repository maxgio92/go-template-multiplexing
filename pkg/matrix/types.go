package matrix

type Point interface{}

type Column struct {
	OrdinateIndex int
	Points        []Point
}
