## Secret Hunter

> Herramienta avanzada de reconocimiento, auditoría web y caza de secretos desarrollada en **Go**. Diseñada para profesionales de la ciberseguridad, pentesters y auditores que buscan automatizar fases críticas de descubrimiento y análisis ofensivo.

---

## Características Principales

**Secret Hunter** automatiza flujos de trabajo de reconocimiento y análisis mediante una arquitectura concurrente y de alto rendimiento:

1. **[FULL SCAN] Auditoría Masiva Total:** Ejecuta las 6 fases de análisis de forma consecutiva sobre el objetivo.
2. **Enumeración de Subdominios & Directorios:** Fuerza bruta optimizada con goroutines y sincronización segura (`sync.WaitGroup`).
3. **Escaneo Universal de Stack:** Identificación de tecnologías web, servidores, cabeceras (`Server`, `X-Powered-By`) y detección de CMS (WordPress Core).
4. **Caza Profunda de Secretos y API Keys:** Análisis de código fuente mediante expresiones regulares optimizadas para detectar credenciales sensibles:
   * Google API Keys
   * AWS Access Key IDs
   * GitHub Access Tokens
   * Slack Webhooks
   * Tokens de autorización (Bearer) y patrones genéricos de contraseñas/API keys.
5. **Fuzzing de Rutas & Detección de SQLi:** Descubrimiento de endpoints críticos (`/wp-admin/`, `/wp-json/`, archivos de respaldo) e inyección automática de payloads de prueba para inyección SQL.
6. **Generador de Enlaces a CVEs y Exploit-DB:** Referenciamiento automático de componentes detectados hacia bases de datos oficiales de vulnerabilidades (**NIST NVD** y **Exploit-DB**).

---

## Requisitos previos

* **Go (Golang)**: Versión 1.20 o superior instalada en tu sistema.

---

## 📦 Instalación y Compilación

Clona el repositorio y compila un binario optimizado para producción utilizando las banderas de optimización de Go:

```bash
# Clonar el repositorio
git clone [https://github.com/tolotelin99/secret-hunter.git](https://github.com/tolotelin99/secret-hunter.git)
cd secret-hunter

# Compilar binario optimizado (sin símbolos de depuración ni rutas absolutas)
go build -ldflags="-s -w" -trimpath -o secret-hunter main.go

## Modo de Uso
Ejecuta el binario interactivo:

Bash
./secret-hunter
Al iniciar, se mostrará el menú principal con las 7 opciones disponibles:

Plaintext
========================================================
    SECRET HUNTER - ULTIMATE FULL SCAN v22.1      
========================================================
 1. [FULL SCAN] Auditoría Masiva Total (6 Fases en 1)
 2. Enumeración de Subdominios & Directorios por Fuerza Bruta
 3. Escaneo Masivo Universal (CMS, Stack & Tecnologías)
 4. Caza Profunda de Secretos, API Keys & Tokens en HTML/JS
 5. Fuzzing Masivo de Rutas & Extracción Automática de Parámetros
 6. Generador Automático de Enlaces a NIST / Exploit-DB
 7. Salir
========================================================
[>] Selecciona una opción [1-7]:

🔒 Calidad de Código y Seguridad
El proyecto mantiene altos estándares de calidad de código y análisis estático (SonarQube Quality Gate: Passed), asegurando mantenibilidad, ausencia de duplicación de literales y optimización de complejidad cognitiva.

⚠️ Aviso Legal (Disclaimer)
Esta herramienta ha sido desarrollada exclusivamente con fines educativos, de investigación y para auditorías de seguridad autorizadas. Los autores no se responsabilizan por el uso indebido o los daños causados por este software. Asegúrate de contar con el consentimiento explícito del propietario antes de auditar cualquier objetivo.

