package _6_athematics

import (
	"math"
	"time"
)

const (
	secondsInHalfClock = 30
	secondsInClock = 2 * secondsInHalfClock
	minuteInHalfClock = 30
	minuteInClock = 2 * minuteInHalfClock
	hourInHalfClock = 6
	hourInClock = 2 * hourInHalfClock
)

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondInRadius(t))
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minuteInRadius(t))
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hourInRadius(t))
}

func secondInRadius(t time.Time) float64 {
	return math.Pi / (secondsInHalfClock / float64(t.Second()))
}

func minuteInRadius(t time.Time) float64 {
	return secondInRadius(t) / minuteInClock +
		math.Pi / (minuteInHalfClock / float64(t.Minute()))
}

func hourInRadius(t time.Time) float64 {
	return minuteInRadius(t) / hourInClock +
		math.Pi / (hourInHalfClock / float64(t.Hour()))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}