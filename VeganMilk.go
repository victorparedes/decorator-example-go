package main

// Hola! Soy un decorador concreto y voy "decorando" con mas funcionalidad a un componente que recibo.
//       Implemento la inteface CUP para poder agergar mi logica.
//
// Yo agrego leche de almendras a lo que sea que me llegue (no me importa que hayan agregado antes, ni lo que agreguen despues)
type VeganMilk struct {
	beberage Beberage
}

func (m *VeganMilk) GetCost() float64 {
	return m.beberage.GetCost() + 160
}

func (m *VeganMilk) GetDescription() string {
	return m.beberage.GetDescription() + ", almond milk(vegan)"
}
