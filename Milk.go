package main

// Hola! Soy un decorador concreto y voy "decorando" con mas funcionalidad a un componente que recibo.
//       Implemento la inteface CUP para poder agergar mi logica.
//
// Yo agrego leche a lo que sea que me llegue (no me importa que hayan agregado antes, ni lo que agreguen despues)
type Milk struct {
	beberage Beberage
}

func (m *Milk) GetCost() float64 {
	return m.beberage.GetCost() + 80
}

func (m *Milk) GetDescription() string {
	return m.beberage.GetDescription() + ", Milk"
}
