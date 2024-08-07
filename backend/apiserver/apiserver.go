package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/ITegs/trippify/database"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type ApiServer interface {
	Main()
}

type apiServer struct {
	db database.DB
}

func NewApiServer(DB database.DB) ApiServer {
	apiServer := &apiServer{
		db: DB,
	}

	return apiServer
}

func (api *apiServer) Main() {
	fmt.Println("Program started!")

	apiHandler := api.buildApi()

	server := http.Server{
		Addr:    ":80",
		Handler: apiHandler,
	}

	fmt.Printf("API server listening on %s\n", server.Addr)
	server.ListenAndServe()
}

type Route struct {
	Method  string
	Path    string
	Handler http.Handler
}

func (api *apiServer) buildApi() *httprouter.Router {
	router := httprouter.New()

	var routes = []*Route{
		{
			Method:  http.MethodGet,
			Path:    "/user/:username",
			Handler: http.HandlerFunc(api.GetUser),
		},
		{
			Method:  http.MethodPost,
			Path:    "/newUser",
			Handler: http.HandlerFunc(api.NewUser),
		},
		{
			Method:  http.MethodGet,
			Path:    "/trip/:tripId",
			Handler: http.HandlerFunc(api.GetTrip),
		},
		{
			Method:  http.MethodPost,
			Path:    "/newTrip",
			Handler: http.HandlerFunc(api.NewTrip),
		},
		{
			Method:  http.MethodPost,
			Path:    "/trip/:tripId/spot/add",
			Handler: http.HandlerFunc(api.AddSpotToTrip),
		},
	}

	for i := 0; i < len(routes); i++ {
		r := routes[i]
		router.Handler(r.Method, "/api"+r.Path, r.Handler)
	}

	return router
}

func (api *apiServer) GetUser(w http.ResponseWriter, r *http.Request) {
	p := httprouter.ParamsFromContext(r.Context())
	userId := p.ByName("username")

	trip, err := api.db.GetUser(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)

	err = encoder.Encode(trip)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	return
}

func (api *apiServer) NewUser(w http.ResponseWriter, r *http.Request) {
	var user database.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	err = api.db.NewUser(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	// TODO: Return the user
	return
}

func (api *apiServer) GetTrip(w http.ResponseWriter, r *http.Request) {
	p := httprouter.ParamsFromContext(r.Context())
	tripId := p.ByName("tripId")

	trip, err := api.db.GetTrip(tripId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*") // TODO: Remove
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)

	err = encoder.Encode(trip)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func (api *apiServer) NewTrip(w http.ResponseWriter, r *http.Request) {
	var trip database.Trip
	err := json.NewDecoder(r.Body).Decode(&trip)
	if err != nil || trip.Title == "" || trip.Owner == primitive.NilObjectID {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	trip.StartDate = time.Now().Format(time.RFC3339)

	err = api.db.NewTrip(&trip)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	// TODO: return the trip
	return
}

func (api *apiServer) AddSpotToTrip(w http.ResponseWriter, r *http.Request) {
	p := httprouter.ParamsFromContext(r.Context())
	tripId := p.ByName("tripId")

	var spot database.Spot
	err := json.NewDecoder(r.Body).Decode(&spot)
	if err != nil || spot.Title == "" || len(spot.Images) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	err = api.db.AddSpotToTrip(tripId, spot)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}
