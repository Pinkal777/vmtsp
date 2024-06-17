load:
	go run . problem1.txt

build:
	go build

tidy:
	go mod tidy

run:
	py .\evaluateShared.py --cmd .\vmtsp.exe --problemDir trainingProblems

.PHONY: load,tidy,build,run