package main

// Hola! Soy un decorador concreto y voy "decorando" con mas funcionalidad a un componente que recibo.
//       Implemento la inteface CUP para poder agergar mi logica.
//
// Yo agrego azucar blanca a lo que sea que me llegue (no me importa que hayan agregado antes, ni lo que agreguen despues)
type Sugar struct {
	beberage Beberage
}

func (s *Sugar) GetCost() float64 {
	return s.beberage.GetCost() + 1.50
}

func (s *Sugar) GetDescription() string {
	return s.beberage.GetDescription() + ", Sugar"
}
