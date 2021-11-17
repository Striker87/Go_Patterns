package pkg

type HomePC struct {
	Cpu     int
	Gpu     int
	Wrapper Wrapper
}

func (pc HomePC) GetPrice() float64 {
	return pc.Wrapper.GetPrice() * float64(pc.Cpu) * float64(pc.Gpu) // берем базову комплектацию, умножаем на кол-во процессоров и видеокарт
}
