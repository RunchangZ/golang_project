# Receipt Processor API



## How to run it:
You can run `make help` to see the instructions 

 1. `make build` - build the image
 2. `make run` - run docker with default json file: M&M.json
 3. `make run file=[filename]` - destroy docker-cleanup container
 4. `make cleanall` - FORCELY remove all containers
 
 Here are some JSON files in the `\example` folder 
 
  List of JSON files other than the default M&M.json:
  
	morning_receipt.json
  
	single_receipt.json
  
	Target_receipt.json

 

If you want to run with single_receipt.json file: 
`make run file=single_receipt.json`

## Test URL: `http://localhost:8000`
Test the working API with Postman

## Endpoints

1. `/receipts` 

**URL**: `http://localhost:8000/receipts` 

**Mathod**: `GET` 

Submit a receipt (For this project, we read the local Json file and show on the server page.) 
![GETreceipt](image.png)

2. `/receipts/process`

**URL**: `http://localhost:8000/receipts/process`

**Method**: `POST` 

Submit a receipt for processing and receive an ID for the receipt.
![POSTprocess](image-1.png)

3. `/receipts/{id}/points`

**URL**: `http://localhost:8000/{id}/points`

Based on the previous screenshot, the url is `http://localhost:8000/receipts/36432e5d-6fb4-4bea-9082-83b08204c1bf/points`

**Method**: `GET`

Retrieve the number of points awarded for a specific receipt.
![GETpoints](image-2.png)

## Error Responses: 
1. 405 Method: please check the chosen method. 
2. 404 PageNote Found: please check the endpoints and port. Id might be invaild. 
3. 400 Bad Request: possible happend on `http://localhost:8000/receipts/process`--invaild receipt

