// Lógica y modelos de colecciones
package collection

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Aquí irán los modelos y funciones para manejar colecciones

type Figura struct {
	ID     string  `json:"id"`
	Nombre string  `json:"nombre"`
	Marca  string  `json:"marca"`
	Serie  string  `json:"serie"`
	Precio float64 `json:"precio"`
	Tamano string  `json:"tamano"`
	Imagen string  `json:"imagen"`
	URL    string  `json:"url"`
}

type Coleccion struct {
	ID      string   `json:"id"`
	Nombre  string   `json:"nombre"`
	Figuras []Figura `json:"figuras"`
}

// Obtener una colección por ID
func GetCollectionByID(id string) *Coleccion {
	// Esta función es solo un ejemplo para usar en db.go
	return nil
}

// Obtener una figura por ID dentro de una colección
func GetFigureByID(collection *Coleccion, figureID string) *Figura {
	for _, f := range collection.Figuras {
		if f.ID == figureID {
			return &f
		}
	}
	return nil
}

// Obtener una colección por ID desde MongoDB
func GetCollectionByIDFromDB(client *mongo.Client, id string) (*Coleccion, error) {
	dbName := os.Getenv("DATABASE_NAME")
	col := client.Database(dbName).Collection("collections")
	filter := bson.M{"id": id}
	var result Coleccion
	err := col.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Obtener todas las colecciones desde MongoDB
func GetAllCollectionsFromDB(client *mongo.Client) ([]Coleccion, error) {
	dbName := os.Getenv("DATABASE_NAME")
	col := client.Database(dbName).Collection("collections")
	cursor, err := col.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	var colecciones []Coleccion
	for cursor.Next(context.Background()) {
		var c Coleccion
		if err := cursor.Decode(&c); err == nil {
			colecciones = append(colecciones, c)
		}
	}
	return colecciones, nil
}
