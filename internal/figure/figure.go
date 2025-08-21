// Lógica y modelos de figuras
package figure

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Aquí irán los modelos y funciones para manejar figuras

// Estructura para definir el modelo de figura global
type Figure struct {
	ID     string  `json:"id"`
	Nombre string  `json:"nombre"`
	Marca  string  `json:"marca"`
	Serie  string  `json:"serie"`
	Precio float64 `json:"precio"`
	Tamano string  `json:"tamano"`
	Imagen string  `json:"imagen"`
	URL    string  `json:"url"`
}

// Obtener todas las figuras globales desde MongoDB
func GetAllFiguresFromDB(client *mongo.Client) ([]map[string]interface{}, error) {
	dbName := os.Getenv("DATABASE_NAME")
	col := client.Database(dbName).Collection("figures")
	cursor, err := col.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	var figuras []map[string]interface{}
	for cursor.Next(context.Background()) {
		var f map[string]interface{}
		if err := cursor.Decode(&f); err == nil {
			figuras = append(figuras, f)
		}
	}
	return figuras, nil
}

func buildFigureFilter(nombre, marca, serie, precio string) bson.M {
	filter := bson.M{}
	if nombre != "" {
		filter["nombre"] = nombre
	}
	if marca != "" {
		filter["marca"] = marca
	}
	if serie != "" {
		filter["serie"] = serie
	}
	if precio != "" {
		filter["precio"] = precio
	}
	return filter
}

// Buscar figuras por filtros
func FindFiguresByFilter(client *mongo.Client, nombre, marca, serie, precio string) ([]map[string]interface{}, error) {
	dbName := os.Getenv("DATABASE_NAME")
	col := client.Database(dbName).Collection("figures")
	filter := buildFigureFilter(nombre, marca, serie, precio)
	cursor, err := col.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	var figuras []map[string]interface{}
	for cursor.Next(context.Background()) {
		var f map[string]interface{}
		if err := cursor.Decode(&f); err == nil {
			figuras = append(figuras, f)
		}
	}
	return figuras, nil
}
