sow-build:
	go build -o builds/ ./services/sow/
	
sow-start:
	./builds/sow

sow-run:
	go run ./services/sow/

sow: sow-build sow-start

# Bro this is demented silly but what the fk ever

test-sow-curl-get:
	curl -L -X GET localhost:8099/food/18ade0f9-a2de-4d07-8dbd-5ccf8bb27625/18ade0f9-a2de-4d07-8dbd-5ccf8bb27625

test-sow-curl-delete:
	curl -L -X DELETE localhost:8099/food/18ade0f9-a2de-4d07-8dbd-5ccf8bb27625/18ade0f9-a2de-4d07-8dbd-5ccf8bb27625

test-sow-curl-update:
	curl -X PUT -H "Content-Type: application/json" --data-binary @test/json/get-filtered-food.json localhost:8099/food/18ade0f9-a2de-4d07-8dbd-5ccf8bb27625/18ade0f9-a2de-4d07-8dbd-5ccf8bb27625

test-sow-curl-gets:
	curl -X GET -H "Content-Type: application/json" --data-binary @test/json/get-filtered-food.json localhost:8099/food/18ade0f9-a2de-4d07-8dbd-5ccf8bb27625

test-sow-curl-create:
	curl -X POST -H "Content-Type: application/json" --data-binary @test/json/get-filtered-food.json localhost:8099/food/18ade0f9-a2de-4d07-8dbd-5ccf8bb27625


