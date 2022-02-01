package pkg

import "fmt"

type RoadStrategy struct {
}

func (r *RoadStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 30
	trafficJam := 2
	distance := endPoint - startPoint
	totalTime := distance * 40 * trafficJam // 40 - среднее время от начала до конца

	fmt.Printf("Road A: [%d] to B: [%d] avg speed [%d] km/h Traffic jam: [%d] Distance: [%d] Total Time: [%d] min\n", startPoint, endPoint, avgSpeed, trafficJam, distance, totalTime)
}
