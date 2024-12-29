sow-build:
	go build -o builds/ ./services/sow/

sow-proto:
	protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     proto/sow/v1/sow.proto

sow-start:
	./builds/sow

sow-run:
	go run ./services/sow/

sow: sow-build sow-start


reap-build:
	go build -o builds/ ./services/reap/

reap-start:
	./builds/reap

reap-run:
	go run ./services/reap/

reap: reap-build reap-start

# Bro this is demented silly but what the fk ever

test-reap-curl-get:
	curl -L -X GET localhost:8099/food/18ade0f9-a2de-4d07-8dbd-5ccf8bb27625/18ade0f9-a2de-4d07-8dbd-5ccf8bb27625

test-reap-curl-delete:
	curl -L -X DELETE localhost:8099/food/18ade0f9-a2de-4d07-8dbd-5ccf8bb27625/18ade0f9-a2de-4d07-8dbd-5ccf8bb27625

test-reap-curl-update:
	curl -X PUT -H "Content-Type: application/json" --data-binary @test/json/get-filtered-food.json localhost:8099/food/18ade0f9-a2de-4d07-8dbd-5ccf8bb27625/18ade0f9-a2de-4d07-8dbd-5ccf8bb27625

test-reap-curl-gets:
	curl -X GET -H "Content-Type: application/json" --data-binary @test/json/get-filtered-food.json localhost:8099/food/18ade0f9-a2de-4d07-8dbd-5ccf8bb27625

test-reap-curl-create:
	curl -X POST -H "Content-Type: application/json" --data-binary @test/json/get-filtered-food.json localhost:8099/food/18ade0f9-a2de-4d07-8dbd-5ccf8bb27625
