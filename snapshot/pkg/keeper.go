package pkg

type Keeper struct {
	Items []*Snapshot
}

func (k *Keeper) Add(s *Snapshot) {
	k.Items = append(k.Items, s)
}

func (k Keeper) Get(index int) *Snapshot {
	return k.Items[index]
}
