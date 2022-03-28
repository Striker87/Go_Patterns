package pkg

type Subject interface {
	Subscribe(consumer Consumer)
	Unsubscribe(consumer Consumer)
	Notify()
}
