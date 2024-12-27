sow-build:
	go build -o builds/ ./services/sow/
	
sow-start:
	./builds/sow

sow-run:
	go run ./services/sow/

sow: sow-build sow-start

