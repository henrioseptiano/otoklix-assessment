# otoklix-assessment
Otoklix BackEnd Interview Assessment

## Introduction
This repository are used for otoklix backend interview assessment. the program are created using golang and gofiber framework. for database using SQLite with GORM inside. it also have testify for unit testing. the code structure are
using clean architecture

## Specifications
- Golang 1.17
- Gofiber 2.33 -> Golang Framework
- GORM 1.23.5 -> Golang ORM
- Godotenv 1.4 -> for getting ENV Variable
- Testify 1.7.1 -> for unit testing 

## Usage

to run the program:
```
    git pull https://github.com/henrioseptiano/otoklix-assessment.git
    go mod tidy
    create .env file and copy paste variable from .env.example
    go run main.go
```


for using testing:
```
    go mod tidy
    cd app/posts/services 
    go test -v
```