# Go CRUD API using Gin and Gorm

This is a simple CRUD (Create, Read, Update, Delete) API implemented in Go using the Gin web framework and Gorm as the ORM (Object Relational Mapping) library.

## Requirements

- Go (v1.16 or later)
- SQLite (or any other supported database)
- `github.com/gin-gonic/gin` package
- `github.com/jinzhu/gorm` package

## Installation

1. Make sure you have Go installed. You can download it from [here](https://golang.org/dl/).
2. Clone this repository to your local machine.
3. Install the required packages using the following commands:
   ```bash
   go get -u github.com/gin-gonic/gin
   go get -u github.com/jinzhu/gorm
   go get -u github.com/jinzhu/gorm/dialects/sqlite

## Run Application

go run main.go

## local host
http://localhost:6666


## Endpoints
POST /data: Create a new data entry.
GET /data/:id: Get data by ID.
PUT /data/:id: Update data by ID.
DELETE /data/:id: Delete data by ID.
GET /data: Get all data entries.

## Data Structure
The data structure for the database table is defined as follows:

ID: Primary key (auto-generated)
ENAME: Employee Name (string)
EAGE: Employee Age (integer)
EADDRESS: Employee Address (string)
EPOST: Employee Post (string)
EUNIVERSITY: Employee University (string)
ESTIPEND: Employee Stipend (string)
EDATE: Date of Entry (string)
ESTATUS: Employee Status (string)

## Usage
Use any API testing tool (e.g., Postman) to interact with the API endpoints.