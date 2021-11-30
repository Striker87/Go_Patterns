package pkg

type Builder struct {
	Collector Collector
}

func NewBuilder(collector Collector) *Builder {
	return &Builder{Collector: collector}
}

func (b *Builder) SetCollector(collector Collector) {
	b.Collector = collector
}

func (b Builder) CreateComputer() Computer {
	b.Collector.SetCore()
	b.Collector.SetMemory()
	b.Collector.SetBrand()
	b.Collector.SetGraphicCard()
	b.Collector.SetMonitor()

	return b.Collector.GetComputer()
}
