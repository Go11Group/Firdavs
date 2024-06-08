package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Go11Group/at_lesson/lesson35/model"
	"github.com/Go11Group/at_lesson/lesson35/storage/postgres"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	Player postgres.PlayerRepo
}

func NewHandler(player postgres.PlayerRepo) *http.Server {

	handler := Handler{Player: player}

	//mux := http.NewServeMux()
	m := mux.NewRouter()

	//http.HandleFunc("/player/", player.GetPlayer)
	m.HandleFunc("/player", handler.CreatePlayer).Methods("POST")
	//mux.HandleFunc("/player/", handler.UpdatePlayer)
	//mux.HandleFunc("/player/", handler.DeletePlayer)
	return &http.Server{Addr: ":8080", Handler: m}
}

//func (p *Handler) GetPlayer(w http.ResponseWriter, r *http.Request) {
//	id := strings.TrimPrefix(r.URL.Path, "/player/")
//
//	if id == p.Id {
//		json.NewEncoder(w).Encode(p)
//	} else {
//		w.Write([]byte("NOT FOUND!"))
//	}
//}

func (p *Handler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	player := model.Player{}

	err := json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("ERROR!"))
		return
	}
	err = p.Player.Create(player)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("ERROR!"))
		return
	}
	fmt.Println(player)

	w.Write([]byte("201 OK OK OK!"))
}

//func (p *Handler) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
//	id := strings.TrimPrefix(r.URL.Path, "/player/")
//	NewPlayer := Player{}
//	json.NewDecoder(r.Body).Decode(&NewPlayer)
//	if id == p.Id {
//		p.Name = NewPlayer.Name
//		p.Number = NewPlayer.Number
//		p.Birthday = NewPlayer.Birthday
//	} else {
//		w.Write([]byte("ERROR"))
//		return
//	}
//}

//func (p *Handler) DeletePlayer(w http.ResponseWriter, r *http.Request) {
//	id := strings.TrimPrefix(r.URL.Path, "/player/")
//
//	if id == p.Id {
//		p = &Player{}
//	}
//	w.Write([]byte("DELETED OK!"))
//}
