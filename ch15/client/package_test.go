package client

import (
	"LearnGo/ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonnacciSeries(5))
	t.Log(series.Square(5))
}
