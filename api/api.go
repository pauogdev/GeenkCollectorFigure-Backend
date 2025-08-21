// Rutas y controladores HTTP
package api

import (
	"awesomeProject/internal/collection"
	"awesomeProject/internal/db"
	"awesomeProject/internal/figure"
	"encoding/json"
	"net/http"
	"strconv"
)

const ContentTypeHeader = "Content-Type"
const ContentTypeJSON = "application/json"

// Datos mock para pruebas
var colecciones = []collection.Coleccion{}
var figurasGlobales = []figure.Figure{} // Figuras disponibles para buscar y añadir

// GET /collections: Listar todas las colecciones desde MongoDB
func GetCollections(w http.ResponseWriter, r *http.Request) {
	client := db.GetMongoClient()
	colecciones, err := collection.GetAllCollectionsFromDB(client)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error consultando colecciones: " + err.Error()))
		return
	}
	w.Header().Set(ContentTypeHeader, ContentTypeJSON)
	json.NewEncoder(w).Encode(colecciones)
}

// POST /collections: Crear una nueva colección
func CreateCollection(w http.ResponseWriter, r *http.Request) {
	var nueva collection.Coleccion
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error en el formato de la colección"))
		return
	}
	colecciones = append(colecciones, nueva)
	w.Header().Set(ContentTypeHeader, ContentTypeJSON)
	json.NewEncoder(w).Encode(nueva)
}

// POST /collections/{id}/figures: Añadir figura a una colección
func AddFigureToCollection(w http.ResponseWriter, r *http.Request) {
	// Aquí deberías obtener el id de la colección desde la URL
	// y añadir la figura al array de figuras de esa colección
	w.Write([]byte("Endpoint para añadir figura a una colección"))
}

// filtrarFiguras filtra un slice de figuras según los parámetros dados
func filtrarFiguras(figuras []figure.Figure, nombre, marca, serie, precio string) []figure.Figure {
	filtradas := []figure.Figure{}
	for _, f := range figuras {
		if !figuraCumpleFiltros(f, nombre, marca, serie, precio) {
			continue
		}
		filtradas = append(filtradas, f)
	}
	return filtradas
}

// GET /collections/{id}/figures: Listar figuras de una colección desde MongoDB
func ListFiguresInCollection(w http.ResponseWriter, r *http.Request) {
	client := db.GetMongoClient()
	id := r.URL.Query().Get("id") // Temporal, idealmente usar mux para path param
	col, err := collection.GetCollectionByIDFromDB(client, id)
	if err != nil || col == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Colección no encontrada"))
		return
	}
	nombre := r.URL.Query().Get("nombre")
	marca := r.URL.Query().Get("marca")
	serie := r.URL.Query().Get("serie")
	precio := r.URL.Query().Get("precio")
	figurasFiltradas := filtrarFiguras(col.Figuras, nombre, marca, serie, precio)
	w.Header().Set(ContentTypeHeader, ContentTypeJSON)
	json.NewEncoder(w).Encode(figurasFiltradas)
}

// GET /figuras?nombre=...&marca=...&serie=...&precio=...: Buscar figuras globales usando MongoDB
func ListGlobalFigures(w http.ResponseWriter, r *http.Request) {
	nombre := r.URL.Query().Get("nombre")
	marca := r.URL.Query().Get("marca")
	serie := r.URL.Query().Get("serie")
	precio := r.URL.Query().Get("precio")
	client := db.GetMongoClient()
	figuras, err := figure.FindFiguresByFilter(client, nombre, marca, serie, precio)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error consultando figuras: " + err.Error()))
		return
	}
	w.Header().Set(ContentTypeHeader, ContentTypeJSON)
	json.NewEncoder(w).Encode(figuras)
}

// figuraCumpleFiltros encapsula la lógica de filtrado para cumplir con estándares de calidad
func figuraCumpleFiltros(f figure.Figure, nombre, marca, serie, precio string) bool {
	if nombre != "" && f.Nombre != nombre {
		return false
	}
	if marca != "" && f.Marca != marca {
		return false
	}
	if serie != "" && f.Serie != serie {
		return false
	}
	if precio != "" {
		return cumplePrecio(f.Precio, precio)
	}
	return true
}

// cumplePrecio separa la lógica de comparación de precio
func cumplePrecio(precioFigura float64, precio string) bool {
	if p, err := strconv.ParseFloat(precio, 64); err == nil {
		if precioFigura != p {
			return false
		}
	}
	return true
}
