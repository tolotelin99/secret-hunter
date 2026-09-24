# Secret Hunter

Herramienta de reconocimiento de alta velocidad desarrollada en Golang, orientada a la auditoría de seguridad ofensiva y cumplimiento bajo el estándar OWASP A05 (Security Misconfiguration).

La aplicación está diseñada para escanear el código fuente de sitios web y aplicaciones en busca de credenciales expuestas, claves de API, tokens de autenticación y llaves criptográficas olvidadas en producción.

## Características Principales

* Motor Concurrente: Implementación de rutinas paralelas en Go (Goroutines y Channels) para maximizar la velocidad de análisis en múltiples objetivos.
* Análisis por Expresiones Regulares (Regex): Detección automatizada de firmas críticas (Tokens AWS, claves privadas RSA/SSH, tokens JWT y claves genéricas).
* Interfaz CLI Dinámica: Soporte nativo para argumentos de consola, permitiendo auditorías rápidas contra objetivos específicos.

## Especificaciones Técnicas

* Lenguaje: Go (Golang)
* Dependencias: Librería estándar de Go (`net/http`, `regexp`, `sync`). Arquitectura "Zero-External-Dependencies".

## Instalación y Despliegue

Clonar el repositorio e inicializar la herramienta desde la terminal: bash
# 1. Clonar el repositorio
git clone https://github.com/tolotelin99/secret-hunter.git

# 2. Acceder al directorio
cd secret-hunter

# 3. Ejecutar la herramienta contra un objetivo
go run main.go -u [https://ejemplo.com](https://ejemplo.com)

## Aviso Legal
Este software ha sido desarrollado con fines exclusivamente educativos y para su uso en entornos controlados o auditorías autorizadas. El autor no asume responsabilidad alguna por el uso indebido de esta herramienta contra infraestructuras sin consentimiento previo.