package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/ITegs/trippify/apiserver/middleware"
	"github.com/ITegs/trippify/database"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
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
		Addr:    ":3000",
		Handler: apiHandler,
	}

	fmt.Printf("API server listening on %s\n", server.Addr)

	// Add CORS support (Cross Origin Resource Sharing)
	handler := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(server.Handler)
	err := http.ListenAndServe(server.Addr, handler)
	if err != nil {
		fmt.Println("SERVER FAILED: ", err)
	}
}

type Middleware interface {
	Handler(httprouter.Handle) httprouter.Handle
}

type MiddlewareFunc func(httprouter.Handle) httprouter.Handle

func (m MiddlewareFunc) Handler(next httprouter.Handle) httprouter.Handle {
	return m(next)
}

type Route struct {
	Method     string
	Path       string
	Handler    httprouter.Handle
	Middleware []Middleware
}

func (api *apiServer) buildApi() *httprouter.Router {
	router := httprouter.New()

	var routes = []*Route{
		{
			Method: http.MethodGet,
			Path:   "/",
			Handler: httprouter.Handle(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
				w.Write([]byte("API is up and running!"))
				w.WriteHeader(http.StatusOK)
			}),
		},
		{
			Method:  http.MethodPost,
			Path:    "/login",
			Handler: httprouter.Handle(api.Login),
		},
		{
			Method:  http.MethodGet,
			Path:    "/user/:username",
			Handler: httprouter.Handle(api.GetUser),
		},
		{
			Method:  http.MethodPost,
			Path:    "/newUser",
			Handler: httprouter.Handle(api.NewUser),
		},
		{
			Method:  http.MethodGet,
			Path:    "/firstTrip",
			Handler: httprouter.Handle(api.GetFirstTrip),
		},
		{
			Method:  http.MethodGet,
			Path:    "/trip/:tripId",
			Handler: httprouter.Handle(api.GetTrip),
		},
		{
			Method:     http.MethodPost,
			Path:       "/newTrip",
			Handler:    httprouter.Handle(api.NewTrip),
			Middleware: []Middleware{MiddlewareFunc(middleware.Auth)},
		},
		{
			Method:     http.MethodPost,
			Path:       "/trip/:tripId/spot/add",
			Handler:    httprouter.Handle(api.AddSpot),
			Middleware: []Middleware{MiddlewareFunc(middleware.Auth)},
		},
		{
			Method:  http.MethodGet,
			Path:    "/spot/:spotId",
			Handler: httprouter.Handle(api.GetSpot),
		},
	}

	for i := 0; i < len(routes); i++ {
		r := routes[i]

		handlerChain := r.Handler
		for _, m := range r.Middleware {
			handlerChain = m.Handler(handlerChain)
		}

		router.Handle(r.Method, "/api"+r.Path, handlerChain)
	}

	return router
}

func (api *apiServer) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var recievedUser database.User
	err := json.NewDecoder(r.Body).Decode(&recievedUser)
	if err != nil || recievedUser.Username == "" || recievedUser.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	realUser, err := api.db.GetUserByUsername(recievedUser.Username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(realUser.Password), []byte(recievedUser.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	jwt, err := middleware.GenerateJWT(realUser.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jwt))
}

func (api *apiServer) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//claims := r.Context().Value(middleware.UserContextKey).(*utils.CustomClaims)
	//fmt.Printf("Hello %+v!\n", claims.Username)

	userId := p.ByName("username")

	trip, err := api.db.GetUserByUsername(userId)
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

func (api *apiServer) NewUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var user database.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	fullUser, err := api.db.NewUser(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	// TODO: Return the user
	json.NewEncoder(w).Encode(fullUser)
	return
}

func (api *apiServer) GetFirstTrip(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	trip, err := api.db.GetFirstTrip()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err)
		return
	}

	owner, err := api.db.GetUserByUsername(trip.OwnerUsername)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err)
		return
	}

	trip.Owner = owner

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)

	err = encoder.Encode(trip)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func (api *apiServer) GetTrip(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	tripId := p.ByName("tripId")

	trip, err := api.db.GetTrip(tripId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err)
		return
	}

	owner, err := api.db.GetUserByUsername(trip.OwnerUsername)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err)
		return
	}

	trip.Owner = owner

	w.Header().Add("Content-Type", "application/json")
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

func (api *apiServer) NewTrip(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var trip database.Trip
	err := json.NewDecoder(r.Body).Decode(&trip)

	if err != nil || trip.Title == "" || trip.OwnerUsername == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	trip.StartDate = time.Now().Format(time.RFC3339)

	trip.Spots = []database.TripSpot{}

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

func (api *apiServer) AddSpot(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	tripId := p.ByName("tripId")

	var spot database.Spot

	tripObjectId, err := primitive.ObjectIDFromHex(tripId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	spot.RoadtripId = tripObjectId

	err = json.NewDecoder(r.Body).Decode(&spot)
	if err != nil || spot.Title == "" || len(spot.Images) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	//claims := r.Context().Value(middleware.UserContextKey).(*utils.CustomClaims)
	//if(claims.Username != )

	err = api.db.AddSpot(spot)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

func (api *apiServer) GetSpot(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	spotId := p.ByName("spotId")

	spot, err := api.db.GetSpot(spotId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)

	err = encoder.Encode(spot)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
