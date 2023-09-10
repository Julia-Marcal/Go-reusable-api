# Go-reusable-api
An api created with golang, gin and gorm. The main function of this api is to be fast, reusable and easy to understand.

## running docker
To run the docker file you should use these scripts:
docker build -t go-reusable-api .
docker run -p 8080:8080 go-reusable-api

## env.go 
Env.go should have 1 function:
func SetSalt that returns the salt in byte type


