package main

// Hola! Yo soy un componente concreto ya que tengo logica implementada y soy quien sera "decorado".
//       Si no tuviera logia implementada, me llamaria "componente abstracto". Podria ser mas flexible y escalable
//       pero para propositos de simplificar la prueba me quedare con algo de logica.
//
// Soy una taza vacia a la cual podemos empezar a agregarle diferentes tipos de bebidas y topicos.
type Cup struct{}

func (c *Cup) GetCost() float64 {
	return 2.50
}

func (c *Cup) GetDescription() string {
	return "Cup"
}
