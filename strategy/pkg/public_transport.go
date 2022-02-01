package pkg

import "fmt"

type PublicTransportStrategy struct {
}

func (p *PublicTransportStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 40
	distance := endPoint - startPoint
	totalTime := distance * 40 // 40 - среднее время от начала до конца

	fmt.Printf("PublicTransport A: [%d] to B: [%d] avg speed [%d] km/h Distance: [%d] Total Time: [%d] min\n", startPoint, endPoint, avgSpeed, distance, totalTime)
}
