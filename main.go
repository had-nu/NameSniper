package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "os"
    "time"

    "github.com/joho/godotenv"
    "github.com/common-nighthawk/go-figure"
)

// Configurações globais
var (
    googleAPIKey string
    googleCX     string
    googleURL    string
    limiteDiario = 100             // Limite gratuito da API do Goolge
    contadorFile = "consultas.json" // Arquivo para rastrear consultas
    timeoutAPI	 = 15 * time.Second
)

// Estrutura para os resultados da API do Google
type GoogleSearchResponse struct {
    Items []struct {
        Title   string `json:"title"`
        Link    string `json:"link"`
        Snippet string `json:"snippet"`
    } `json:"items"`
}

// Estrutura para o contador de consultas
type Contador struct {
    Data      string `json:"data"`      // Data no formato "YYYY-MM-DD"
    Consultas int    `json:"consultas"` // Número de consultas feitas no dia
}

func main() {
	// Exibe o título estilizado
    banner := figure.NewFigure("NameSniper", "slant", true)
    banner.Print()
    fmt.Println("\nVersão 1.0 - Desenvolvido por hadnu\n")
    
    // Carrega variáveis de ambiente do arquivo .env
    if err := godotenv.Load(); err != nil {
        fmt.Println("Erro ao carregar .env:", err)
        os.Exit(1)
    }

    // Carrega as variáveis de ambiente
    googleAPIKey = os.Getenv("GOOGLE_API_KEY")
    googleCX = os.Getenv("GOOGLE_CX")
    googleURL = os.Getenv("GOOGLE_URL")

    // Verifica se as variáveis de ambiente estão definidas
    if googleAPIKey == "" || googleCX == "" || googleURL == "" {
        fmt.Println("Erro: Variáveis de ambiente GOOGLE_API_KEY, GOOGLE_CX e GOOGLE_URL são obrigatórias.")
        os.Exit(1)
    }

    // Verifica os argumentos da linha de comando
    if len(os.Args) < 2 {
        fmt.Println("Uso: namesnipe [primeiro_nome] [sobrenome]")
        os.Exit(1)
    }

    nome := os.Args[1]
    sobrenome := ""
    if len(os.Args) > 2 {
        sobrenome = os.Args[2]
    }

    query := url.QueryEscape(nome) // Codifica caracteres especiais
    if sobrenome != "" {
        query = url.QueryEscape(fmt.Sprintf("%s %s", nome, sobrenome))
    }
    fmt.Printf("Buscando por: %s %s\n", nome, sobrenome)

    // Verifica o limite de consultas antes de executar
    if !podeFazerConsulta() {
        fmt.Println("Limite diário de 100 consultas gratuitas atingido. Tente novamente amanhã.")
        os.Exit(1)
    }

    resultados := buscarGoogle(query)
    if len(resultados) > 0 {
        fmt.Println("Resultados encontrados: ")
        for i, resultado := range resultados {
            fmt.Printf("%d. %s\n   URL: %s\n   Snippet: %s\n", i+1, resultado.Title, resultado.Link, resultado.Snippet)
        }
    } else {
        fmt.Println("Nenhum resultado encontrado.")
    }

    // Atualiza o contador após a consulta
    atualizarContador()
}

func buscarGoogle(query string) []struct {
    Title   string `json:"title"`
    Link    string `json:"link"`
    Snippet string `json:"snippet"`
} {
    url := fmt.Sprintf("%s?key=%s&cx=%s&q=%s", googleURL, googleAPIKey, googleCX, query)
    fmt.Println("URL da requisição: ", url) // Debug: Mostra a URL

    // Configura um cliente HTTP com timeout
    client := &http.Client{Timeout: timeoutAPI}
    resp, err := client.Get(url)
    if err != nil {
        fmt.Println("Erro na requisição:", err)
        return nil
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Printf("Erro na API: %s\n", resp.Status)
        body, _ := ioutil.ReadAll(resp.Body)
        fmt.Println("Resposta bruta:", string(body))
        return nil
    }

    var searchResponse GoogleSearchResponse
    if err := json.NewDecoder(resp.Body).Decode(&searchResponse); err != nil {
        fmt.Println("Erro ao decodificar resposta: ", err)
        return nil
    }

    return searchResponse.Items
}

// Verifica se ainda há consultas disponíveis no dia
func podeFazerConsulta() bool {
    contador := carregarContador()
    hoje := time.Now().UTC().Format("2006-01-02") // Usa UTC

    // Se a data mudou, reseta o contador
    if contador.Data != hoje {
        contador.Data = hoje
        contador.Consultas = 0
        salvarContador(contador)
    }

    // Verifica o limite
    if contador.Consultas >= limiteDiario {
        return false
    }

    // Avisa se está na última consulta
    if contador.Consultas == limiteDiario-1 {
        fmt.Println("Atenção: Esta é a última consulta gratuita do dia!")
    }

    return true
}

// Atualiza o contador após cada consulta
func atualizarContador() {
    contador := carregarContador()
    contador.Consultas++
    salvarContador(contador)
    fmt.Printf("Consultas usadas hoje: %d/%d\n", contador.Consultas, limiteDiario)
}

// Carrega o contador do arquivo
func carregarContador() Contador {
    data, err := ioutil.ReadFile(contadorFile)
    if err != nil {
        // Se o arquivo não existe, retorna um contador zerado
        return Contador{Data: time.Now().UTC().Format("2006-01-02"), Consultas: 0}
    }

    var contador Contador
    if err := json.Unmarshal(data, &contador); err != nil {
        return Contador{Data: time.Now().UTC().Format("2006-01-02"), Consultas: 0}
    }
    return contador
}

// Salva o contador no arquivo
func salvarContador(contador Contador) {
    data, err := json.Marshal(contador)
    if err != nil {
        fmt.Println("Erro ao salvar contador: ", err)
        return
    }
    if err := ioutil.WriteFile(contadorFile, data, 0644); err != nil {
        fmt.Println("Erro ao escrever arquivo: ", err)
    }
}