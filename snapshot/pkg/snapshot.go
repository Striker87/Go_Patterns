package pkg

type Snapshot struct {
	state string
}

func (s Snapshot) GetSavedState() string {
	return s.state
}
