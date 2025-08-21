// Conexión y utilidades para MongoDB
package db

import (
	"awesomeProject/internal/collection"
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// InitMongo carga la configuración y conecta a MongoDB
func InitMongo() error {
	// Cargar variables de entorno desde config.env
	err := godotenv.Load("config/config.env")
	if err != nil {
		return err
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(uri)
	c, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	client = c
	return nil
}

// GetMongoClient devuelve el cliente de MongoDB
func GetMongoClient() *mongo.Client {
	return client
}

// Insertar una colección
func InsertCollection(collection collection.Coleccion) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	dbName := os.Getenv("DATABASE_NAME")
	col := client.Database(dbName).Collection("collections")
	_, err := col.InsertOne(ctx, collection)
	return err
}

// Actualizar una colección (por id)
func UpdateCollection(id string, updated collection.Coleccion) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	dbName := os.Getenv("DATABASE_NAME")
	col := client.Database(dbName).Collection("collections")
	filter := map[string]interface{}{"id": id}
	update := map[string]interface{}{"$set": updated}
	_, err := col.UpdateOne(ctx, filter, update)
	return err
}

// Eliminar una colección (por id)
func DeleteCollection(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	dbName := os.Getenv("DATABASE_NAME")
	col := client.Database(dbName).Collection("collections")
	filter := map[string]interface{}{"id": id}
	_, err := col.DeleteOne(ctx, filter)
	return err
}

// Eliminar una figura de una colección (por id de colección y id de figura)
func DeleteFigureFromCollection(collectionID, figureID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	dbName := os.Getenv("DATABASE_NAME")
	col := client.Database(dbName).Collection("collections")
	filter := map[string]interface{}{"id": collectionID}
	update := map[string]interface{}{"$pull": map[string]interface{}{"figuras": map[string]interface{}{"id": figureID}}}
	_, err := col.UpdateOne(ctx, filter, update)
	return err
}
