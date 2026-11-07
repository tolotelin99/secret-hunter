package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

func printBanner() {
	fmt.Println("========================================================")
	fmt.Println("   [>_<] SECRET HUNTER - ULTIMATE FULL SCAN v22.1      ")
	fmt.Println("========================================================")
	fmt.Println(" 1. [FULL SCAN] Auditoría Masiva Total (6 Fases en 1)")
	fmt.Println(" 2. Enumeración de Subdominios & Directorios por Fuerza Bruta")
	fmt.Println(" 3. Escaneo Masivo Universal (CMS, Stack & Tecnologías)")
	fmt.Println(" 4. Caza Profunda de Secretos, API Keys & Tokens en HTML/JS")
	fmt.Println(" 5. Fuzzing Masivo de Rutas & Extracción Automática de Parámetros")
	fmt.Println(" 6. Generador Automático de Enlaces a NIST / Exploit-DB")
	fmt.Println(" 7. Salir")
	fmt.Println("========================================================")
}

func getExploitLinks(query string) string {
	encoded := url.QueryEscape(query)
	return fmt.Sprintf("Exploit-DB: https://www.exploit-db.com/search?q=%s | NIST NVD: https://nvd.nist.gov/vuln/search/results?query=%s", encoded, encoded)
}

func auditUniversalTech(targetURL string) {
	if !strings.HasPrefix(targetURL, "http://") && !strings.HasPrefix(targetURL, "https://") {
		targetURL = "https://" + targetURL
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	fmt.Printf("\n[+] [FASE 1/6] Reconocimiento universal de Stack en: %s\n", targetURL)
	resp, err := client.Get(targetURL)
	if err != nil {
		fmt.Printf("[-] Error al conectar: %s\n", err)
		return
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	htmlContent := string(bodyBytes)
	headers := resp.Header

	server := headers.Get("Server")
	if server != "" {
		fmt.Printf("    • Servidor Web: %s\n", server)
	} else {
		fmt.Println("    • Servidor Web: Oculto / Desconocido")
	}

	poweredBy := headers.Get("X-Powered-By")
	if poweredBy != "" {
		fmt.Printf("    • X-Powered-By: %s\n", poweredBy)
	}

	if strings.Contains(htmlContent, "wp-content") || strings.Contains(htmlContent, "wp-includes") {
		fmt.Println("    • CMS Detectado: WordPress Core")
	} else {
		fmt.Println("    • CMS: No detectado o personalizado.")
	}
}

func huntSecrets(targetURL string) {
	if !strings.HasPrefix(targetURL, "http://") && !strings.HasPrefix(targetURL, "https://") {
		targetURL = "https://" + targetURL
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	fmt.Printf("\n[+] [FASE 2/6] Buscando API Keys y Secretos en código fuente...\n")
	resp, err := client.Get(targetURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	content := string(body)

	patterns := map[string]*regexp.Regexp{
		"Google API Key":       regexp.MustCompile(`AIza[0-9A-Za-z-_]{35}`),
		"AWS Access Key ID":    regexp.MustCompile(`AKIA[0-9A-Z]{16}`),
		"GitHub Access Token":  regexp.MustCompile(`ghp_[0-9a-zA-Z]{36}`),
		"Slack Webhook":        regexp.MustCompile(`https://hooks\.slack\.com/services/T[a-zA-Z0-9_]{8}/B[a-zA-Z0-9_]{8}/[a-zA-Z0-9_]{24}`),
		"Authorization Bearer": regexp.MustCompile(`(?i)bearer\s+[a-zA-Z0-9_\-\.]{20,}`),
		"Generic API Key":      regexp.MustCompile(`(?i)(api_key|apikey|secret|passwd|password|access_token)\s*[:=]\s*['"][a-zA-Z0-9_\-\.]{8,64}['"]`),
	}

	foundSecrets := 0
	for secretName, re := range patterns {
		matches := re.FindAllString(content, -1)
		if len(matches) > 0 {
			for _, match := range matches {
				fmt.Printf("    [!] ¡Secreto encontrado (%s): %s\n", secretName, match)
				foundSecrets++
			}
		}
	}

	if foundSecrets == 0 {
		fmt.Println("    • No se detectaron API Keys ni tokens expuestos en el HTML principal.")
	}
}

func auditAndExtractEndpoints(baseURL string) []string {
	if !strings.HasPrefix(baseURL, "http://") && !strings.HasPrefix(baseURL, "https://") {
		baseURL = "https://" + baseURL
	}
	baseURL = strings.TrimSuffix(baseURL, "/")

	paths := []string{
		"/", "/index.php", "/wp-admin/", "/wp-login.php", 
		"/wp-json/wp/v2/users", "/wp-json/wp/v2/posts", "/xmlrpc.php",
		"/search", "/query.php", "/api/",
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	fmt.Printf("\n[+] [FASE 3/6] Fuzzing de rutas y extracción automática de parámetros...\n")
	var discoveredEndpoints []string

	for _, p := range paths {
		target := baseURL + p
		resp, err := client.Get(target)
		if err != nil {
			continue
		}
		
		bodyBytes, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		if resp.StatusCode == 200 || resp.StatusCode == 403 {
			fmt.Printf("    [+] Endpoint Activo: %s (Status: %d)\n", target, resp.StatusCode)
			discoveredEndpoints = append(discoveredEndpoints, target)
			
			// Uso opcional del body para verificar formularios o entradas
			if strings.Contains(string(bodyBytes), "name=") {
				fmt.Printf("        ↳ [i] Campos de entrada detectados en HTML.\n")
			}
		}
	}

	return discoveredEndpoints
}

func auditExtractedParameters(endpoints []string) {
	fmt.Printf("\n[+] [FASE 4/6] Inyectando payloads automáticos en parámetros detectados...\n")
	
	if len(endpoints) == 0 {
		fmt.Println("    • No hay endpoints activos para inyectar.")
		return
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	commonParams := []string{"id", "s", "p", "page", "query", "search", "cat"}
	sqliPayloads := []string{"' OR '1'='1", "1' UNION SELECT NULL,NULL--"}

	vulnFound := false
	for _, ep := range endpoints {
		for _, param := range commonParams {
			for _, payload := range sqliPayloads {
				separator := "?"
				if strings.Contains(ep, "?") {
					separator = "&"
				}
				testURL := fmt.Sprintf("%s%s%s=%s", ep, separator, param, url.QueryEscape(payload))
				
				resp, err := client.Get(testURL)
				if err != nil {
					continue
				}
				body, _ := io.ReadAll(resp.Body)
				resp.Body.Close()
				bodyStr := string(body)

				if strings.Contains(bodyStr, "sql syntax") || strings.Contains(bodyStr, "mysql_fetch") {
					fmt.Printf("    [!] [CRITICAL] ¡Posible Inyección SQL detectada!\n")
					fmt.Printf("        - URL: %s\n", testURL)
					vulnFound = true
				}
			}
		}
	}

	if !vulnFound {
		fmt.Println("    • Escaneo de parámetros finalizado: No se detectaron errores de base de datos directos.")
	}
}

func runSubdomainBruteforce(domain string) {
	domain = strings.TrimPrefix(domain, "http://")
	domain = strings.TrimPrefix(domain, "https://")
	domain = strings.TrimSuffix(domain, "/")

	fmt.Printf("\n[+] [FASE 5/6] Enumeración de Subdominios y Directorios por Fuerza Bruta...\n")

	subdomains := []string{"www", "mail", "ftp", "admin", "api", "dev", "staging"}
	directories := []string{"admin", "login", "wp-admin", "backup", "api", "uploads"}

	client := &http.Client{
		Timeout: 4 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	var wg sync.WaitGroup
	var mutex sync.Mutex

	for _, sub := range subdomains {
		wg.Add(1)
		go func(sub string) {
			defer wg.Done()
			subURL := fmt.Sprintf("https://%s.%s", sub, domain)
			resp, err := client.Get(subURL)
			if err == nil {
				resp.Body.Close()
				mutex.Lock()
				fmt.Printf("    [+] Subdominio Activo: %s\n", subURL)
				mutex.Unlock()
			}
		}(sub)
	}
	wg.Wait()

	baseURL := fmt.Sprintf("https://%s", domain)
	for _, dir := range directories {
		wg.Add(1)
		go func(dir string) {
			defer wg.Done()
			dirURL := fmt.Sprintf("%s/%s", baseURL, dir)
			resp, err := client.Get(dirURL)
			if err == nil {
				resp.Body.Close()
				if resp.StatusCode == 200 || resp.StatusCode == 403 {
					mutex.Lock()
					fmt.Printf("    [+] Endpoint Encontrado: %s (Status: %d)\n", dirURL, resp.StatusCode)
					mutex.Unlock()
				}
			}
		}(dir)
	}
	wg.Wait()
}

func runExploitLinksGenerator(domain string) {
	fmt.Printf("\n[+] [FASE 6/6] Generando Enlaces de Explotación y CVEs (NIST & Exploit-DB)...\n")
	
	components := []string{"wordpress", "php", "apache", "mysql"}
	for _, comp := range components {
		fmt.Printf("    • Componente [%s]:\n", comp)
		fmt.Printf("      -> %s\n", getExploitLinks(comp))
	}
}

func runFullScan(targetURL string) {
	fmt.Printf("\n========================================================")
	fmt.Printf("\n[*] INICIANDO AUDITORÍA ULTIMATE FULL SCAN EN: %s\n", targetURL)
	fmt.Printf("========================================================")

	auditUniversalTech(targetURL)
	huntSecrets(targetURL)
	endpoints := auditAndExtractEndpoints(targetURL)
	auditExtractedParameters(endpoints)
	runSubdomainBruteforce(targetURL)
	runExploitLinksGenerator(targetURL)

	fmt.Printf("\n========================================================")
	fmt.Printf("\n[*] ¡AUDITORÍA ULTIMATE FULL SCAN COMPLETADA CON ÉXITO!\n")
	fmt.Printf("========================================================")
}

func queryNISTAndExploits(query string) {
	encoded := url.QueryEscape(query)
	fmt.Println("\n========================================================")
	fmt.Println("         RECURSOS OFICIALES Y REPOSITORIOS DE EXPLOITS  ")
	fmt.Println("========================================================")
	fmt.Printf("[+] NIST NVD CVE Search:      https://nvd.nist.gov/vuln/search/results?query=%s\n", encoded)
	fmt.Printf("[+] Exploit-DB Database:      https://www.exploit-db.com/search?q=%s\n", encoded)
	fmt.Printf("[+] GitHub Advisory Database: https://github.com/advisories?query=%s\n", encoded)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printBanner()
		fmt.Print("[>] Selecciona una opción [1-7]: ")
		if !scanner.Scan() {
			break
		}
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Print("[?] Dominio o URL objetivo para Full Scan (ej. example.com): ")
			if scanner.Scan() {
				runFullScan(strings.TrimSpace(scanner.Text()))
			}
		case "2":
			fmt.Print("[?] Dominio objetivo para subdominios y directorios (ej. example.com): ")
			if scanner.Scan() {
				runSubdomainBruteforce(strings.TrimSpace(scanner.Text()))
			}
		case "3":
			fmt.Print("[?] URL objetivo para Stack: ")
			if scanner.Scan() {
				auditUniversalTech(strings.TrimSpace(scanner.Text()))
			}
		case "4":
			fmt.Print("[?] URL objetivo para secretos: ")
			if scanner.Scan() {
				huntSecrets(strings.TrimSpace(scanner.Text()))
			}
		case "5":
			fmt.Print("[?] URL objetivo para fuzzing y extracción: ")
			if scanner.Scan() {
				endpoints := auditAndExtractEndpoints(strings.TrimSpace(scanner.Text()))
				auditExtractedParameters(endpoints)
			}
		case "6":
			fmt.Print("[?] Componente o tecnología a buscar: ")
			if scanner.Scan() {
				queryNISTAndExploits(strings.TrimSpace(scanner.Text()))
			}
		case "7":
			fmt.Println("[*] Saliendo de Secret Hunter...")
			return
		default:
			fmt.Println("[-] Opción no válida.")
		}

		fmt.Println("\n\nPresiona ENTER para continuar...")
		scanner.Scan()
	}
}
