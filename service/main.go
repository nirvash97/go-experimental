package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb+srv://aegisx1:papth0391@experimental-01.8lsgx.mongodb.net/?retryWrites=true&w=majority&appName=Experimental-01"

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Printf("[%s] %s %s \n", time.Now().UTC().Format(timeFormat), r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func main() {
	/// Mongo DB Connect
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	var echo bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&echo); err != nil {
		panic(err)
	}
	// ==================MongoDb ===========================
	mux := http.NewServeMux()
	mux.HandleFunc("/favicon.ico", http.NotFound)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, "Hello world main")

	})

	mux.HandleFunc("/schedule", func(w http.ResponseWriter, r *http.Request) {

		var scheduleList []Schedule

		scheduleDetail := Schedule{Date: time.Now().UTC(), Display: false, Meeting: []ScheduleMeeting{
			{MeetingTime: "10.00", Status: "Available"},
			{MeetingTime: "11.00", Status: "Available"},
			{MeetingTime: "12.00", Status: "Blocked"}},
		}
		scheduleList = append(scheduleList, scheduleDetail)
		jsonStructure, err := json.Marshal(scheduleList)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonStructure)
		}

	})

	mux.HandleFunc("/movie/languages/", func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Path[len("/movie/languages/"):]
		filter := bson.D{{Key: "languages", Value: bson.D{{Key: "$in", Value: bson.A{lang}}}}}

		db := client.Database("sample_mflix")
		collection := db.Collection("movies")
		option := options.Find()
		option.SetLimit(10)
		option.SetProjection(bson.D{
			{Key: "title", Value: 1},
			{Key: "year", Value: 1},
			{Key: "runtime", Value: 1},
			{Key: "imdb", Value: 1},
			{Key: "released", Value: 1},
			{Key: "languages", Value: 1},
			{Key: "tomatoes", Value: -1},
		})
		cursor, err := collection.Find(context.Background(), filter, option)
		if err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer cursor.Close(context.TODO())
		var movies []bson.M
		if err = cursor.All(context.TODO(), &movies); err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		decodeData, err := json.Marshal(movies)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(decodeData)
		}

	})

	mux.HandleFunc("/person/", func(w http.ResponseWriter, r *http.Request) {
		var personData Person
		id := r.URL.Path[len("/person/"):]
		//fmt.Fprintf(w, "Hello , %s!", id)
		if id == "1" {
			personData = Person{Name: "Tossaporn Meesiri", Age: 26, Height: 160, Weight: 62.3}
			jsonPerson, err := json.Marshal(personData)
			if err != nil {
				fmt.Fprint(w, err)

			} else {
				w.Header().Set("Content-Type", "application/json")
				w.Write(jsonPerson)
			}
		} else {
			fmt.Fprint(w, "id not found")
		}

	})

	mux.HandleFunc("/person", func(w http.ResponseWriter, r *http.Request) {
		var personData Person
		id := r.URL.Path[len("/person/"):]
		fmt.Fprintf(w, "Your id is %s", id)
		if id == fmt.Sprintf("%d", 1) {
			personData = Person{Name: "Tossaporn Meesiri", Age: 26, Height: 160, Weight: 62.3}
		}

		jsonPerson, err := json.Marshal(personData)
		if err != nil {
			fmt.Fprint(w, err)

		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonPerson)
		}

	})

	mux.HandleFunc("/header", func(w http.ResponseWriter, r *http.Request) {
		header := r.Header
		jsonHeader, err := json.Marshal(header)
		if err != nil {
			fmt.Fprint(w, err)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonHeader)
		}

	})
	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})
	mux.HandleFunc("/greeting/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Path[len("/greeting/"):]
		fmt.Fprintf(w, "Hello , %s!", name)
	})

	mux.HandleFunc("/qp", func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("page")

		if param == "" {
			http.Error(w, "Param 'page' is required for this action", http.StatusBadRequest)
			return
		}
		intValue, err := strconv.Atoi(param)
		if err != nil {
			http.Error(w, "Param 'page' must be number", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Param field page : %d", intValue)
	})

	server := &http.Server{
		Addr:         ":8081",
		Handler:      middleware(mux),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	errServer := server.ListenAndServe()
	if errServer != nil {
		fmt.Println(err)
	}

}
