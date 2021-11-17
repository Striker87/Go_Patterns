package pkg

type ServerPC struct {
	Cpu     int
	Memory  int
	Wrapper Wrapper
}

func (pc ServerPC) GetPrice() float64 {
	return pc.Wrapper.GetPrice() * float64(pc.Cpu) * float64(pc.Memory) // берем базову комплектацию, умножаем на кол-во процессоров и оперативки
}
