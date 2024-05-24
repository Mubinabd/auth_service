mig-run:
	migrate create -ext sql -dir migrations -seq create_table

mig-up:
	migrate -database 'postgres://postgres:pass@localhost:5432/go10?sslmode=disable' -path migrations up

mig-down:
	migrate -database 'postgres://postgres:pass@localhost:5432/go10?sslmode=disable' -path migrations down
