package store

import (
	"math/rand"
	"time"
)

func genCode() string {
	const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const n = 8
	byts := make([]byte, n)

	// Inicializa a seed para gerar números aleatórios diferentes em cada execução
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		byts[i] = characters[rand.Intn(len(characters))]
	}

	return string(byts)
}
