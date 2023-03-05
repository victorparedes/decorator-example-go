package main

// Hola! Soy un decorador concreto y voy "decorando" con mas funcionalidad a un componente que recibo.
//       Implemento la inteface CUP para poder agergar mi logica.
//
// Yo agrego TE a lo que sea que me llegue (no me importa que hayan agregado antes, ni lo que agreguen despues)
type Tea struct {
	beberage Beberage
}

func (t *Tea) GetCost() float64 {
	return t.beberage.GetCost() + 300
}

func (t *Tea) GetDescription() string {
	return t.beberage.GetDescription() + ", TEA Early grey"
}
