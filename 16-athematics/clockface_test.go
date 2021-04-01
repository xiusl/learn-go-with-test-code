package _6_athematics

import (
	"math"
	"testing"
	"time"
)

func TestSecondInRadians(t *testing.T) {
	testCases := []struct{
		time time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30 ), math.Pi},
		{simpleTime(0, 0, 0 ), 0},
		{simpleTime(0, 0, 45 ), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7 ), (math.Pi / 30) * 7},
	}

	for _, tc := range testCases {
		t.Run(testName(tc.time), func(t *testing.T) {
			got := secondInRadius(tc.time)
			if !roughlyEqualFloat(got, tc.angle) {
				t.Fatalf("Wanted %v radians, but got %v", tc.angle, got)
			}
		})
	}
}

func TestSecondHandPoint(t *testing.T) {
	testCases := []struct{
		time time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, tc := range testCases {
		t.Run(testName(tc.time), func(t *testing.T) {
			got := secondHandPoint(tc.time)
			if !roughlyEqualPoint(got, tc.point) {
				t.Fatalf("Wanted %v Point, but got %v", tc.point, got)
			}
		})
	}
}

func TestMinuteInRadians(t *testing.T) {
	testCases := []struct{
		time time.Time
		angle float64
	}{
		{simpleTime(0, 30, 0 ), math.Pi},
		{simpleTime(0, 0, 0 ), 0},
		{simpleTime(0, 0, 7 ), (math.Pi / (30 * 60)) * 7},
	}

	for _, tc := range testCases {
		t.Run(testName(tc.time), func(t *testing.T) {
			got := minuteInRadius(tc.time)
			if !roughlyEqualFloat(got, tc.angle) {
				t.Fatalf("Wanted %v radians, but got %v", tc.angle, got)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	testCases := []struct{
		time time.Time
		point Point
	}{
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}

	for _, tc := range testCases {
		t.Run(testName(tc.time), func(t *testing.T) {
			got := minuteHandPoint(tc.time)
			if !roughlyEqualPoint(got, tc.point) {
				t.Fatalf("Wanted %v Point, but got %v", tc.point, got)
			}
		})
	}
}

func TestHourInRadians(t *testing.T) {
	testCases := []struct{
		time time.Time
		angle float64
	}{
		{simpleTime(30, 0, 0 ), math.Pi},
		{simpleTime(0, 0, 0 ), 0},
		{simpleTime(0, 1, 30 ), math.Pi / ((6 * 60 * 60) / 90)},
	}

	for _, tc := range testCases {
		t.Run(testName(tc.time), func(t *testing.T) {
			got := hourInRadius(tc.time)
			if !roughlyEqualFloat(got, tc.angle) {
				t.Fatalf("Wanted %v radians, but got %v", tc.angle, got)
			}
		})
	}
}

func TestHourHandPoint(t *testing.T) {
	testCases := []struct{
		time time.Time
		point Point
	}{
		{simpleTime(30, 0, 0), Point{0, -1}},
		{simpleTime(45, 0, 0), Point{-1, 0}},
	}

	for _, tc := range testCases {
		t.Run(testName(tc.time), func(t *testing.T) {
			got := hourHandPoint(tc.time)
			if !roughlyEqualPoint(got, tc.point) {
				t.Fatalf("Wanted %v Point, but got %v", tc.point, got)
			}
		})
	}
}

func roughlyEqualFloat(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat(a.X, b.X) && roughlyEqualFloat(a.Y, b.Y)
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(2021, time.April, 1, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("17:09:00")
}