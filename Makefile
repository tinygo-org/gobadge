flash: flash-tinygo

prepare-gceu:
	go run cmd/main.go -conf=gceu22

flash-gceu: prepare-gceu perform-flash

prepare-gcuk:
	go run cmd/main.go -conf=gcuk22

flash-gcuk: prepare-gcuk perform-flash

prepare-gcus:
	go run cmd/main.go -conf=gcus22

flash-gcus: prepare-gcus perform-flash

prepare-tinygo:
	go run cmd/main.go -conf=tinygo

flash-tinygo: prepare-tinygo perform-flash

perform-flash:
	tinygo flash -size short -target pybadge -ldflags="-X main.YourName='$(NAME)' -X main.YourTitle1='$(TITLE1)' -X main.YourTitle2='$(TITLE2)'" .
