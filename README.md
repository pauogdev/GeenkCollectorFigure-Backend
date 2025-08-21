# GeekCollectorFigures Backend 🖼️

API REST en **Go** para gestionar colecciones de figuras (Anime/Manga, Marvel/DC, Funko Pop, entre otras).  
Su objetivo es proporcionar un backend robusto y reutilizable, ideal para integrar con frontends web o móviles, facilitando la búsqueda, organización y acceso comunitario a colecciones de figuras.

## Estructura UML

![UML Figuras](uml%20api%20fuguras.png)

## ​ Contenido del proyecto
- Endpoints REST para crear, consultar y organizar colecciones y figuras.
- Lógica de negocio modular presentada en `internal/`.
- Configuración flexible con soporte para MongoDB.
- Scripts de scraping para obtener datos desde tiendas.
- Tests unitarios y documentación básica.
- Imagen UML para comprender la arquitectura (
  ver `docs/uml api figuras.png`).

## ​​ Requisitos
- [Go](https://go.dev/) 1.20 o superior.
- MongoDB funcionando (local o remoto).
- IDE recomendado: **GoLand** o **VS Code** con extensiones para Go.

## ​​ Cómo ejecutar
1. Clona este repositorio:
   ```bash
   git clone https://github.com/pauogdev/GeenkCollectorFigure-Backend.git
   cd GeenkCollectorFigure-Backend


## Licencia
Apache 2.0
