flash:
	tinygo flash -target pybadge .

flash-gcuk:
	go run cmd/main.go
	tinygo flash -target pybadge .