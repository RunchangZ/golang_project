package utils

import (
	"encoding/json"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"time"
	"unicode"
	"github.com/RunchangZ/golang_project/types"
	"github.com/google/uuid"
)


func ReadJSONFile(filepath string) (types.Receipt, error) {
	var receipt types.Receipt

	//Open the JSON file
	file, err := os.Open(filepath) 
	if err != nil {
		return receipt, err
	}
	defer file.Close()

	//Read the file content
	bytesValue, err := io.ReadAll(file) 
	if err != nil {
		return receipt, err
	}

	//Parse the JSON
	err = json.Unmarshal(bytesValue, &receipt) 
	if err != nil {
		return receipt, err
	}

	return receipt, nil
}




const (
	dateLayout = "2006-01-02"
	timeLayout = "15:04"
)


//Help function for Rule1 
func countAlphanumericChars(name string) int64{
	var counter int64 = 0

	//Loop over each char in the strinig and check if it is a letter or a digit 
	for _, char := range name{
		if unicode.IsLetter(char) || unicode.IsDigit(char){
			counter++ 
		}
	}
	return counter 

}


func CalPoints(receipt types.Receipt) int64 {
	//Apply rules, base on the receipt's info, calculate and return points  

	var points int64 
	points = 0

	//Rule 1: # of AlphanumericChars
    points += countAlphanumericChars(receipt.Retailer)


	//Total is a string based on the yaml file, here we need to convert it into a float 
	totalVal, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		log.Println(err)
	}


	//Rule 2: Round dollar
	if math.Mod(totalVal, 1.0) == 0 { // sine totalVal is a float, we prefer to use Mod 
		points += 50
	}

	//Rule 3: multiple of 0.25 
	if math.Mod(totalVal, 0.25) == 0 {
		points += 25
	}

	//Rule 4: number of items  
    points += int64(len(receipt.Items)) * 5 / 2

	//Rule 5: length of ShortDescription
	for _, item := range receipt.Items{

		if countAlphanumericChars(item.ShortDescription) % 3 == 0 {

			//Convert string price to float
			priceVal, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				log.Println(err)
			} 
			points += int64(math.Ceil(priceVal * 0.2))
		}
	}

	//Rule 6: Odd date 
	t, _ := time.Parse(dateLayout, receipt.PurchaseDate)
	if t.Day() % 2 == 1 {
		points += 6
	}

	//Rule 7: after 2:00PM and before 4:00PM 
    t, _ = time.Parse(timeLayout, receipt.PurchaseTime)
    if t.Hour() >= 14 && t.Hour() < 16 {
        points += 10
    }

	return points
}


func GenerateReceiptID() string {
	uuid := uuid.New()
	return uuid.String()
}