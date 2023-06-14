
package api

import(
	"net/http"
	"github.com/gorilla/mux"
)
//This file is about server, and calls functions th  

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server{
	return &Server{
		listenAddr: listenAddr,
	}
}


func (s *Server) Start(filepath string) error{

	router := mux.NewRouter()

	//This one is optional. 
	router.HandleFunc("/receipts", func(w http.ResponseWriter, r *http.Request) {
		s.HandleHomePage(w, r, filepath)
	}).Methods("GET") 

	//Returns the ID assigned to the receipt
	router.HandleFunc("/receipts/process", func(w http.ResponseWriter, r *http.Request) {
		s.HandlePostProcessReceipts(w, r, filepath)
	}).Methods("POST")

	//Returns the points awarded for the receipt
	router.HandleFunc("/receipts/{id}/points", s.HandleGetPoints).Methods("GET")
	err := http.ListenAndServe(s.listenAddr, router)
	return err 
}









