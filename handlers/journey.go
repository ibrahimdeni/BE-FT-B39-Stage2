package handlers

import (
	journeydto "thejourney/dto/journey"
	dto "thejourney/dto/result"

	"encoding/json"
	"net/http"
	"strconv"
	"thejourney/models"
	"thejourney/repositories"

	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"

	// "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerJourney struct {
	JourneyRepository repositories.JourneyRepository
}

func HandlerJourney(JourneyRepository repositories.JourneyRepository) *handlerJourney {
	return &handlerJourney{JourneyRepository}
}

func (h *handlerJourney) FindJourneys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	journeys, err := h.JourneyRepository.FindJourneys()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	for i, p := range journeys {
	journeys[i].Image = os.Getenv("PATH_FILE") + p.Image
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: journeys}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerJourney) GetJourney(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var journeys models.Journey

	journeys, err := h.JourneyRepository.GetJourney(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	journeys.Image = os.Getenv("PATH_FILE") + journeys.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: journeys}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerJourney) CreateJourney(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	//get datafile ari middleware dan nyetor filename var disini,,
	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string) // add this code
	//get data user dari token ges yak
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))
	//CHECKDOANG
	// fmt.Println(userId)


	request := journeydto.JourneyRequest{
		Title		: r.FormValue("title"),
		UserID		: userId,
		Image		: filename,
		Description	: r.FormValue("description"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	CreatedAt := time.Now()

	journey := models.Journey{
		Title		: request.Title,
		UserID		: userId,
		Image		: filename,
		CreatedAt: CreatedAt,
		Description	: request.Description,
		Message	: "success",
	}

	data, err := h.JourneyRepository.CreateJourney(journey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	journey, _= h.JourneyRepository.GetJourney(data.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseJourney(journey)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerJourney) UpdateJourney(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	//id journey??
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	// fmt.Println('test 1')
	//get datafile ari middleware dan nyetor filename var disini,,
	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string) // add this code
	//get data user dari token ges yak
	// userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	// userId := int(userInfo["id"].(float64))
	
	request := journeydto.UpdateJourneyRequest{
		Title		: r.FormValue("title"),
		Description	: r.FormValue("description"),
	}
	
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	
	UpdatedAt := time.Now()

	journey, _ := h.JourneyRepository.GetJourney(int(id))

	// journey.Title = request.Title
	// journey.UserID = request.UserID
	// journey.Image = request.Image
	// journey.Description = request.Description
	journey.UpdatedAt = UpdatedAt

	if request.Title != "" {
		journey.Title = request.Title
	}
	if filename != "false" {
		journey.Image = filename
	}
	if request.Description != "" {
		journey.Description = request.Description
	}
	
	data, err := h.JourneyRepository.UpdateJourney(journey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseJourney(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerJourney) DeleteJourney(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	journey, err := h.JourneyRepository.GetJourney(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.JourneyRepository.DeleteJourney(journey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseJourney(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseJourney(u models.Journey) models.Journey { 
	return models.Journey{
		ID			: u.ID,
		Title		: u.Title,
		UserID		: u.UserID,
		User		: u.User,
		Image		: u.Image,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Description	: u.Description,
		Message: u.Message,
	}
}