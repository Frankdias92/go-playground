package	main 

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Endereço representa o resultado da consulta de CEP
type Endereco struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	Fonte       string `json:"fonte"` // Guarda a fonte da API que retornou o resultado
}

// Cliente HTTP compartilhado com timeout
var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func main() {
	// Configurando o servidor HTTP
	http.HandleFunc("/cep/", buscaCEPHandler)
	fmt.Println("Servidor iniciado em :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func buscaCEPHandler(w http.ResponseWriter, r *http.Request) {
	// Extrair o CEP da URL. Exemplo: /cep/01001000
	if len(r.URL.Path) < 6 {
		http.Error(w, "CEP não informado", http.StatusBadRequest)
		return
	}
	
	cep := r.URL.Path[5:] // Remove "/cep/" do início
	
	// Criando um contexto com cancelamento
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Garante que todos os recursos serão liberados
	
	// Canal para receber o resultado da primeira API que responder
	resultChan := make(chan Endereco, 1)
	
	// Inicia as consultas concorrentes
	go consultaViaCEP(ctx, cep, resultChan)
	go consultaBrasilAPI(ctx, cep, resultChan)
	
	// Aguarda o primeiro resultado ou timeout geral
	select {
	case endereco := <-resultChan:
		// Cancelamos o contexto para interromper a outra requisição
		cancel()
		
		// Retorna o resultado como JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(endereco)
		
	case <-time.After(5 * time.Second):
		// Timeout geral caso ambas as APIs demorem muito
		http.Error(w, "Timeout ao consultar APIs de CEP", http.StatusGatewayTimeout)
	}
}

// Consulta a API ViaCEP
func consultaViaCEP(ctx context.Context, cep string, resultChan chan<- Endereco) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Printf("Erro ao criar requisição para ViaCEP: %v", err)
		return
	}
	
	resp, err := httpClient.Do(req)
	if err != nil {
		// Verificamos se o erro foi causado pelo cancelamento do contexto
		if ctx.Err() == context.Canceled {
			log.Println("Requisição para ViaCEP cancelada")
			return
		}
		log.Printf("Erro ao consultar ViaCEP: %v", err)
		return
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		log.Printf("ViaCEP respondeu com código %d", resp.StatusCode)
		return
	}
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Erro ao ler resposta do ViaCEP: %v", err)
		return
	}
	
	var endereco Endereco
	if err := json.Unmarshal(body, &endereco); err != nil {
		log.Printf("Erro ao parsear resposta do ViaCEP: %v", err)
		return
	}
	
	// Adiciona a fonte da informação
	endereco.Fonte = "ViaCEP"
	
	// Verifica novamente se o contexto foi cancelado antes de enviar o resultado
	select {
	case <-ctx.Done():
		log.Println("Contexto cancelado antes de enviar resultado do ViaCEP")
		return
	case resultChan <- endereco:
		log.Println("Resultado do ViaCEP enviado primeiro")
	}
}

// Consulta a API BrasilAPI
func consultaBrasilAPI(ctx context.Context, cep string, resultChan chan<- Endereco) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Printf("Erro ao criar requisição para BrasilAPI: %v", err)
		return
	}
	
	resp, err := httpClient.Do(req)
	if err != nil {
		// Verificamos se o erro foi causado pelo cancelamento do contexto
		if ctx.Err() == context.Canceled {
			log.Println("Requisição para BrasilAPI cancelada")
			return
		}
		log.Printf("Erro ao consultar BrasilAPI: %v", err)
		return
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		log.Printf("BrasilAPI respondeu com código %d", resp.StatusCode)
		return
	}
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Erro ao ler resposta do BrasilAPI: %v", err)
		return
	}
	
	var brasilAPIResp struct {
		Cep        string `json:"cep"`
		State      string `json:"state"`
		City       string `json:"city"`
		Neighborhood string `json:"neighborhood"`
		Street     string `json:"street"`
	}
	
	if err := json.Unmarshal(body, &brasilAPIResp); err != nil {
		log.Printf("Erro ao parsear resposta do BrasilAPI: %v", err)
		return
	}
	
	// Converte para nosso formato padrão
	endereco := Endereco{
		Cep:        brasilAPIResp.Cep,
		Logradouro: brasilAPIResp.Street,
		Bairro:     brasilAPIResp.Neighborhood,
		Localidade: brasilAPIResp.City,
		UF:         brasilAPIResp.State,
		Fonte:      "BrasilAPI",
	}
	
	// Verifica novamente se o contexto foi cancelado antes de enviar o resultado
	select {
	case <-ctx.Done():
		log.Println("Contexto cancelado antes de enviar resultado do BrasilAPI")
		return
	case resultChan <- endereco:
		log.Println("Resultado do BrasilAPI enviado primeiro")
	}
}
