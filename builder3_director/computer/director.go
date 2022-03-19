package computer

type Director struct {
	b BuilderI
}

func NewDirector(b BuilderI) *Director {
	return &Director{
		b: b,
	}
}

func (d *Director) SetBuilder(b BuilderI) {
	d.b = b
}

func (d *Director) BuildComputer() Computer {
	return d.b.Build()
}
