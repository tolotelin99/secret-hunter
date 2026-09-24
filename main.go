package main

import (
    "flag"
    "fmt"
    "io"
    "net/http"
    "os"
    "regexp"
    "sync"
    "time"
)

// Firmas avanzadas (Regex) para detectar credenciales y tokens expuestos
var reglasSecretos = map[string]*regexp.Regexp{
    "Token AWS (Access Key ID)": regexp.MustCompile(`AKIA[0-9A-Z]{16}`),
    "Clave Privada RSA/SSH":    regexp.MustCompile(`-----BEGIN (RSA|OPENSSH|PRIVATE) KEY-----`),
    "Token JWT (Bearer)":        regexp.MustCompile(`eyJ[A-Za-z0-9_-]{10,}\.[A-Za-z0-9_-]{10,}\.[A-Za-z0-9_-]{10,}`),
    "Posible API Key Genérica":  regexp.MustCompile(`(?i)(api[_-]?key|secret[_-]?key|password)['"]?\s*[:=]\s*['"][a-zA-Z0-9_\-]{16,45}['"]`),
}

type Hallazgo struct {
    URL     string
    Regla   string
    Secreto string
}

func escanearURL(url string, resultados chan<- Hallazgo, wg *sync.WaitGroup) {
    defer wg.Done()

    client := &http.Client{
        Timeout: 10 * time.Second,
    }

    // Simulamos un User-Agent de navegador legítimo para evitar bloqueos básicos de WAF
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return
    }
    req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("[-] Error conectando a %s: %v\n", url, err)
        return
    }
    defer resp.Body.Close()

    cuerpo, err := io.ReadAll(resp.Body)
    if err != nil {
        return
    }

    codigoFuente := string(cuerpo)

    for nombreRegla, regex := range reglasSecretos {
        coincidencias := regex.FindAllString(codigoFuente, -1)
        for _, match := range coincidencias {
            resultados <- Hallazgo{
                URL:     url,
                Regla:   nombreRegla,
                Secreto: match,
            }
        }
    }
}

func main() {
    // Definición de argumentos por consola
    urlFlag := flag.String("u", "", "URL objetivo a escanear (ej. https://ejemplo.com)")
    flag.Parse()

    // Validar si se ingresó el parámetro obligatorio
    if *urlFlag == "" {
        fmt.Println("========================================")
        fmt.Println("  SECRET HUNTER v1.0 - AppSec Engine")
        fmt.Println("========================================")
        fmt.Println("[!] Error: Debes especificar un objetivo usando la bandera -u")
        fmt.Println("\nUso correcto:")
        fmt.Println("  go run main.go -u https://tu-objetivo.com")
        fmt.Println("\nAyuda:")
        flag.PrintDefaults()
        os.Exit(1)
    }

    fmt.Println("========================================")
    fmt.Println("  SECRET HUNTER v1.0 - AppSec Engine")
    fmt.Println("========================================")
    fmt.Printf("[*] Analizando objetivo: %s\n\n", *urlFlag)

    var wg sync.WaitGroup
    resultadosChan := make(chan Hallazgo, 100)

    inicio := time.Now()

    wg.Add(1)
    go escanearURL(*urlFlag, resultadosChan, &wg)

    go func() {
        wg.Wait()
        close(resultadosChan)
    }()

    contador := 0
    for hallazgo := range resultadosChan {
        contador++
        fmt.Printf("[!] ¡ALERTA CRÍTICA DETECTADA!\n")
        fmt.Printf("    - Tipo:     %s\n", hallazgo.Regla)
        fmt.Printf("    - Hallazgo: %s\n\n", hallazgo.Secreto)
    }

    duracion := time.Since(inicio)
    fmt.Println("========================================")
    fmt.Printf("[+] Escaneo completado en %s\n", duracion)
    fmt.Printf("[+] Total de secretos encontrados: %d\n", contador)
}