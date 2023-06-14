package api


import(
	"net/http"
	"github.com/RunchangZ/golang_project/types"
	"github.com/RunchangZ/golang_project/utils"
	"encoding/json"
	"github.com/gorilla/mux"
)




// not a requirement, but I want to see the reseipt itself on the server. 
func (s *Server) HandleHomePage(w http.ResponseWriter, r *http.Request, filepath string) {
	var receipt types.Receipt

	//Read it locally from example folders. 
	receipt, err := utils.ReadJSONFile(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(receipt)
}



//in-memory stoarage, for the future, can be replace with database. 
var idPointsMap = map[string]int64{}

func (s *Server) HandlePostProcessReceipts(w http.ResponseWriter, r *http.Request, filepath string) {

	var receipt types.Receipt

	//Read it locally from example folders. 
	receipt, err := utils.ReadJSONFile(filepath)
	if err != nil {
		http.Error(w, "The receipt is invalid", http.StatusBadRequest)
		return
	}
	
	// Generate random ID
	receiptID := utils.GenerateReceiptID()

	//Calcluate points based on the given rules
	points := utils.CalPoints(receipt)

	idPointsMap[receiptID] = points 

	// Prepare the response JSON
	response := types.ProcessReceiptResponse{
		ID: receiptID,
	}

	// log.Printf("points : %d", points)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) HandleGetPoints(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	receiptID := vars["id"]

	// Look up the receipt's points by the ID
	points, ok := idPointsMap[receiptID]
	if !ok {
		http.Error(w, "No receipt found for that id", http.StatusNotFound)
		return
	}

	// Prepare the response JSON
	response := types.GetPointsResponse{
		Points: points,
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}