/**
 * Copyright (C) 2019 Piyush Bhangale
 *
 * This file is part of hastebin-golang.
 *
 * hastebin-golang is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * hastebin-golang is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with hastebin-golang.  If not, see <http://www.gnu.org/licenses/>.
 */

package database

import (
	"context"
	"fmt"
	"hastebin-golang/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var c *mongo.Client

// InitiateClient - Initialize the database client
func InitiateClient() bool {
	clientOptions := options.Client().ApplyURI(config.Get("mongoURI"))
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
		return false
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println("[Database] Successfully connected!")
	c = client
	return true
}

func CreateDocument(content string) string {
	collection := c.Database(config.Get("dbName")).Collection("files")
	result, err := collection.InsertOne(context.TODO(), RawDocument{CreatedAt: time.Now().Unix(), FileContent: content})
	if err != nil {
		log.Fatal(err)
		return ""
	}

	// Convert ObjectID to string
	return result.InsertedID.(primitive.ObjectID).Hex()
}

func GetDocument(key string) (string, bool) {
	var doc Document
	collection := c.Database(config.Get("dbName")).Collection("files")
	result := collection.FindOne(context.TODO(), bson.M{"_id": key})
	err := result.Decode(&doc)
	if err != nil {
		return "", false
	}
	return doc.FileContent, true
}
