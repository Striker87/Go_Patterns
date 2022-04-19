package pkg

type StationManager struct {
	PlatformFree bool
	Queue        []Vehicle
}

func NewStationManager() *StationManager {
	return &StationManager{PlatformFree: true}
}

func (s *StationManager) CanArrive(v Vehicle) bool {
	if s.PlatformFree {
		s.PlatformFree = false
		return true
	}
	s.Queue = append(s.Queue, v)
	return false
}

func (s *StationManager) NotifyAboutGo() {
	if s.PlatformFree {
		s.PlatformFree = false
	}
	if len(s.Queue) > 0 {
		firstTrainInQueue := s.Queue[0]
		s.Queue = s.Queue[1:]
		firstTrainInQueue.PermitArrive()
	}
}
