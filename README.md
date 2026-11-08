# Secret Hunter v22.1

**Secret Hunter** es un framework de reconocimiento ofensivo y auditoría de seguridad automatizada desarrollado en **Go**, diseñado para identificar rápidamente tecnologías, cazar secretos/API keys en el código fuente, realizar fuzzing de rutas, inyectar parámetros vulnerables, enumerar subdominios y generar enlaces directos a repositorios de exploits (Exploit-DB y NIST NVD).

---

## Características Principales (6 Fases del Ultimate Full Scan)

1. **Reconocimiento Universal de Stack:** Identifica servidores web, cabeceras `X-Powered-By` (PHP, Java, Node.js) y detección de CMS (WordPress Core).
2. **Caza Profunda de Secretos (Regex Engine):** Escanea el código fuente en busca de API Keys de Google, AWS, tokens de GitHub, Webhooks de Slack, tokens Bearer y credenciales genéricas.
3. **Fuzzing de Rutas & Extracción de Parámetros:** Barrido de endpoints críticos de WordPress y aplicaciones web con análisis heurístico de campos de entrada.
4. **Inyección Automatizada de Parámetros:** Pruebas de payloads orientados a SQL Injection (`SQLi`) y validación de respuestas del servidor.
5. **Fuerza Bruta Concurrente:** Enumeración multihilo (`goroutines`) de subdominios activos y directorios protegidos.
6. **Generador de Enlaces de Explotación:** Enlaces directos a bases de datos oficiales de CVEs (NIST NVD) y Exploit-DB basados en los componentes detectados.

---

## Instalación y Requisitos

Asegúrate de tener instalado **Go** (versión 1.18 o superior) en tu sistema.

1. Clona o clona el repositorio localmente:
   ```bash
   git clone [https://github.com/toloteli99/secret-hunter.git](https://github.com/toloteli99/secret-hunter.git)
   cd secret-hunter
Inicializa o compila el proyecto:

Bash
go build -o secret-hunter main.go

## Uso
Puedes ejecutar la herramienta directamente mediante Go:

Bash
go run main.go
Al iniciar, se mostrará el menú interactivo con las opciones disponibles:

Plaintext
========================================================
   SECRET HUNTER - ULTIMATE FULL SCAN v22.1      
======================================================
 1. [FULL SCAN] Auditoría Masiva Total (6 Fases en 1)
 2. Enumeración de Subdominios & Directorios por Fuerza Bruta
 3. Escaneo Masivo Universal (CMS, Stack & Tecnologías)
 4. Caza Profunda de Secretos, API Keys & Tokens en HTML/JS
 5. Fuzzing Masivo de Rutas & Extracción Automática de Parámetros
 6. Generador Automático de Enlaces a NIST / Exploit-DB
 7. Salir
========================================================
## Tecnologías Utilizadas
Lenguaje: Go (Golang)

Concurrencia: Goroutines y sync.WaitGroup para escaneos de alta velocidad.

Seguridad: Expresiones regulares (Regex) y cliente HTTP con soporte para omisión de certificados TLS (InsecureSkipVerify).

⚖️ Aviso Legal
Esta herramienta ha sido desarrollada exclusivamente con fines educativos, de investigación y para auditorías de seguridad autorizadas en entornos de prueba (CTFs y laboratorios). El uso no autorizado contra sistemas sin consentimiento explícito es ilegal.
