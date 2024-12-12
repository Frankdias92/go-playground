package points

import "fmt"

func Exercice() {
	var a int = 10
	var ptr *int = &a                                  // ptr aponta para a
	fmt.Println("Valor de a:", a)                      // 10
	fmt.Println("Endereço de a:", &a)                  // Endereço na memória
	fmt.Println("Valor do ponteiro:", ptr)             // Endereço onde ptr aponta
	fmt.Println("Valor apontado pelo ponteiro:", *ptr) // 10

	*ptr = 20                          // Modifica o valor de a através do ponteiro
	fmt.Println("Novo valor de a:", a) // 20
}
