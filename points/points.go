package points

import "fmt"

func GetPoints() {
	x := 10        // Variável normal
	take(&x)       // Passa o endereço da variável x para a função take
	fmt.Println(x) // Deve imprimir 100

	y := "try to change me"
	takeString(&y) // Passa o endereço da string y para a função takeString
	fmt.Println(y) // Deve imprimir "I dont change"

	xPtr := create()   // Cria um inteiro e obtém seu ponteiro
	fmt.Println(*xPtr) // Imprime o valor do inteiro através do ponteiro

	var z int = 0  // Inicializa uma variável z
	deRef(&z)      // Passa o endereço da variável z para a função deRef
	fmt.Println(z) // Deve imprimir 100 agora, pois z foi modificado pela função deRef
}

func take(x *int) {
	*x = 100 // Modifica o valor do inteiro apontado por x
}

func takeString(y *string) {
	*y = "I dont change" // Modifica a string apontada por y
}

func create() *int {
	stack := new(int) // Aloca um novo inteiro no heap e retorna seu ponteiro
	*stack = 10       // Define o valor do inteiro como 10
	return stack      // Retorna o ponteiro para o inteiro alocado no heap
}

func deRef(x *int) {
	*x = 100 // Modifica o valor do inteiro apontado por x
}
