package main

type M struct {
	Tipo  int
	Valor float64
}

func calcular(ms []M) float64 {
	res := float64(0)
	for _, m := range ms {
		if m.Tipo == 1 {
			res -= m.Valor
		} else {
			res += m.Valor
		}
	}
	return res
}

func calcular1(ms []M) (res float64) {
	for _, m := range ms {
		if m.Tipo == 1 {
			res -= m.Valor
		} else {
			res += m.Valor
		}
	}
	return
}

func calcular2(ms []M) float64 {
	res := float64(0)
	for i := 0; i < len(ms); i++ {
		if ms[i].Tipo == 1 {
			res -= ms[i].Valor
		} else {
			res += ms[i].Valor
		}
	}
	return res
}

func calcular3(ms []M) (res float64) {

	for i := 0; i < len(ms); i++ {
		if ms[i].Tipo == 1 {
			res -= ms[i].Valor
		} else {
			res += ms[i].Valor
		}
	}
	return
}

func calcular4(ms []M) float64 {
	res := float64(0)
	for _, m := range ms {
		if m.Tipo == 1 {
			res -= m.Valor
			continue
		}
		res += m.Valor
	}
	return res
}

func calcular5(ms []M) (res float64) {
	for _, m := range ms {
		if m.Tipo == 1 {
			res -= m.Valor
			continue
		}
		res += m.Valor
	}
	return
}

func calcular6(ms []M) float64 {
	res := float64(0)
	for i := 0; i < len(ms); i++ {
		if ms[i].Tipo == 1 {
			res -= ms[i].Valor
			continue
		}
		res += ms[i].Valor
	}
	return res
}

func calcular7(ms []M) (res float64) {

	for i := 0; i < len(ms); i++ {
		if ms[i].Tipo == 1 {
			res -= ms[i].Valor
			continue
		}
		res += ms[i].Valor
	}
	return
}

func calcular8(ms []M) (res float64) {

	for i := 0; i < len(ms); i++ {
		res += ms[i].Valor
		if ms[i].Tipo == 1 {
			res -= ms[i].Valor
			res -= ms[i].Valor
		}

	}
	return
}
