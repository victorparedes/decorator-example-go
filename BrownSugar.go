package main

// Hola! Soy un decorador concreto y voy "decorando" con mas funcionalidad a un componente que recibo.
//       Implemento la inteface CUP para poder agergar mi logica.
//
// Yo agrego azucar de mascabo a lo que sea que me llegue (no me importa que hayan agregado antes, ni lo que agreguen despues)
type BrownSugar struct {
	beberage Beberage
}

func (b *BrownSugar) GetCost() float64 {
	return b.beberage.GetCost() + 1.50
}

func (b *BrownSugar) GetDescription() string {
	return b.beberage.GetDescription() + ", Sugar"
}
