package main

// Hola! Soy un decorador concreto y voy "decorando" con mas funcionalidad a un componente que recibo.
//       Implemento la inteface CUP para poder agergar mi logica.
//
// Yo agrego cafe a lo que sea que me llegue (no me importa que hayan agregado antes, ni lo que agreguen despues)
type Coffe struct {
	beberage Beberage
}

func (c *Coffe) GetCost() float64 {
	return c.beberage.GetCost() + 600
}

func (c *Coffe) GetDescription() string {
	return c.beberage.GetDescription() + ", coffe"
}
