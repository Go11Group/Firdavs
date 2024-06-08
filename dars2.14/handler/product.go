package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"n11/Firdavs/dars2.14/model"

)

func (h *Handler) Product(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid product ID: %s", err.Error()), http.StatusBadRequest)
		return
	}

	product, err := h.Productr.Product_Get(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while getting product: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(product); err != nil {
		http.Error(w, fmt.Sprintf("Error while encoding response: %s", err.Error()), http.StatusInternalServerError)
	}
}

func (h *Handler) Productdelete(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/productdelete/")

	err := h.Productr.Product_delete(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error while Decode, err: %s", err.Error())))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode("deleted")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error while Encode, err: %s", err.Error())))
	}
}

func (h *Handler) ProductCreate(w http.ResponseWriter, r *http.Request) {
	var b model.Product
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error while Decode, err: %s", err.Error())))
		return
	}

	fmt.Println(b)
	w.Write([]byte("SUCCESS"))

	err = h.Productr.Product_Create(b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error while Decode, err: %s", err.Error())))
		return
	}
}

func (h *Handler) ProductUpdate(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/productupdate/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("invalid user ID, err: %s", err.Error())))
		return
	}
	fmt.Println(3)


	var updated model.Product
	err = json.NewDecoder(r.Body).Decode(&updated)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error while decoding, err: %s", err.Error())))
		return
	}
	fmt.Println(2)


	err = h.Productr.Product_Update(id, updated)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error while updating user, err: %s", err.Error())))
		return
	}
	fmt.Println(1)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error while encoding response, err: %s", err.Error())))
	}
}