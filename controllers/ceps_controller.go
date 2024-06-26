package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type CepStorage interface {
	GetAllCEPS() ([]CEP, error)
	GetCEPByID(id string) (CEP, error)
	CreateNewCEP(cep CEP) error
}

type Controller struct {
	cepStorage CepStorage
}

type CEP struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
	DDD        string `json:"ddd"`
}

func NewController(storage CepStorage) *Controller {
	return &Controller{cepStorage: storage}
}

func (c *Controller) GetCEPS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ceps, err := c.cepStorage.GetAllCEPS()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(ceps)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Controller) GetCEPByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cep, err := c.cepStorage.GetCEPByID(params["id"])
	if err != nil {
		http.Error(w, "CEP não encontrado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(cep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Controller) CreateCEP(w http.ResponseWriter, r *http.Request) {
	var newCEP CEP
	err := json.NewDecoder(r.Body).Decode(&newCEP)
	if err != nil {
		http.Error(w, "Falha ao decodificar o corpo da requisição", http.StatusBadRequest)
		return
	}

	err = c.cepStorage.CreateNewCEP(newCEP)
	if err != nil {
		http.Error(w, "Falha ao criar novo CEP", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newCEP)
	fmt.Printf("CEP adicionado ao banco de dados: %+v\n", newCEP)
}
