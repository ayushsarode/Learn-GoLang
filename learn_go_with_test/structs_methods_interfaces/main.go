package main

type Reactangle struct {
	Width	float64
	Height	float64
}	

func Perimeter(rectangle Reactangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(reactangle Reactangle) float64 {

	return reactangle.Width* reactangle.Height
}
