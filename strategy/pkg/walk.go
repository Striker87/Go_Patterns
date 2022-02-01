package pkg

import "fmt"

type WalkStrategy struct {
}

func (w *WalkStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 4
	distance := endPoint - startPoint
	totalTime := distance * 60 // 60 - среднее время от начала до конца

	fmt.Printf("Walk A: [%d] to B: [%d] avg speed [%d] km/h Distance: [%d] Total Time: [%d] min\n", startPoint, endPoint, avgSpeed, distance, totalTime)
}
