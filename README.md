# GeekCollectorFigures

API para la gestión de colecciones de figuras (Anime/Manga, Marvel/DC, Funkopop, etc.)

## Estructura UML

![UML Figuras](uml%20api%20fuguras.png)

## Descripción
Esta API permite crear colecciones y añadir figuras, almacenando los datos en MongoDB como JSON.

## Carpetas principales
- **api/**: Rutas y controladores
- **internal/**: Lógica de negocio y modelos
- **config/**: Configuración
- **scripts/**: Scraping de tiendas
- **docs/**: Documentación
- **test/**: Tests unitarios

## Instalación
1. Instala Go y MongoDB
2. Configura las variables en `config/config.env`
3. Ejecuta el proyecto:
   ```bash
   go run main.go
   ```

## Licencia
Apache 2.0