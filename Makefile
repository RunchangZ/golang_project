# Define the image name
IMAGE_NAME := fetch


.PHONY: help 
help:
	@echo ""
	@echo "-- Help Menu"
	@echo ""
	@echo "   1. make build - build the image"
	@echo "   2. make run - run docker with default json file: M&M.json"
	@echo "   3. make run file=[filename] - destroy docker-cleanup container"
	@echo "   4. make cleanall - forcely remove all containers"
	@echo ""
	@echo "List of JSON files"
	@echo "M&M.json  (default)"
	@echo "morning_receipt.json"
	@echo "single_receipt.json"
	@echo "Target_receipt.json"


.PHONY: build
build:
	go build 
	docker build -t $(IMAGE_NAME) .

.PHONY: run
run:
	docker run -d -p 8000:8000 -e INPUT_FILENAME=$(file) $(IMAGE_NAME)

.PHONY: cleanall
cleanall:
	docker rm $$(docker ps -aq) -f


