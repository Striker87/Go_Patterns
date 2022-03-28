package pkg

type Publisher struct {
	Name      string
	Consumers map[string]Consumer
}

func (p *Publisher) Subscribe(consumer Consumer) {
	p.Consumers[consumer.GetName()] = consumer
}

func (p *Publisher) Unsubscribe(consumer Consumer) {
	delete(p.Consumers, consumer.GetName())
}

func (p *Publisher) Notify() {
	for _, consumer := range p.Consumers {
		consumer.Update(p.Name)
	}
}
