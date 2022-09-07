flash: flash-tinygo

flash-gceu:
	go run cmd/main.go -conf=gceu22
	tinygo flash -size short -target pybadge .

flash-gcuk:
	go run cmd/main.go -conf=gcuk22
	tinygo flash -size short -target pybadge .

flash-gcus:
	go run cmd/main.go -conf=gcus22
	tinygo flash -size short -target pybadge .

flash-tinygo:
	go run cmd/main.go -conf=tinygo
	tinygo flash -size short -target pybadge .

