# HTTP2 Rest API Server
Built in Go (1.22.2) using the native HTTP package with h2c support to disable TLS, removing HTTPS enforcement.

## Setup
Please ensure that you have golang installed and setup within your enviroment.

First, initialize the dependancies and packages:
```bash
cd go-http2-server
go mod init go-http2-server
```
Docker is required to containerize the project
Building the image:
```bash
docker build -t go-http2-server .
```

Starting the service:
```bash
docker run -d -p 8080:8080 go-http2-server
```

The server will be live on http://localhost:8080/api and is now ready to accept requests.

## Commands
Primary URL for server: http://localhost:8080/api/v1

- To add a pet, please append a valid json object into the body of the request. Send the request to: http://localhost:808/api/v1/pets
- To get a pet record, send request to http://localhost:8080/api/v1/pets/:id
- To delete a pet record, send request to http://localhost:8080/api/v1/pets/:id

Methods available:
**POST**
```bash
curl -X 'POST' \
  'http://localhost:8080/api/v1/pets' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "Doggie",
  "status": "available",
  "category": "dog",
  "tags": [
    "friendly",
    "energetic"
  ],
  "age": 3,
  "color": "Golden",
  "gender": "male",
  "weight": 25.5,
  "dateAdded": "2023-01-20T08:30:00Z"
}'
```
**GET**
```bash
curl -X 'GET' \
  'http://localhost:8080/api/v1/pets/10' \
  -H 'accept: application/json'
```
**DELETE**
```bash
curl -X 'DELETE' \
  'http://localhost:8080/api/v1/pets/10' \
  -H 'accept: */*'
```

To run the approriate unit tests:
```bash
go test ./tests
# or
go test -v ./tests
```

## Design and assumptions
The system is built on an MVC framework, with handlers managing the logic for REST controllers and the store handling data model modifications.

Several Go frameworks were evaluated for the project, but they were ultimately deemed unsuitable. The frameworks considered included:
- Echo v4
- Fiber
- Gin

Fiber initially appeared ideal due to its lightweight nature and fasthttp integration, which offers superior HTTP performance. However, Fiber currently lacks HTTP/2 support because the fasthttp package hasn't been extended for HTTP/2 enough for integration.

Gin and Echo are both lightweight, feature-rich, and high-performance frameworks. They provide pre-existing boilerplate code, facilitating quick API server creation and extension with minimal effort. Nonetheless, they were not used for this project due to their extensive features, which were unnecessary and added unwanted complexity. Additionally, their pre-configured settings made fine-tuning difficult.

Utimately, the default Go HTTP package proved sufficient for creating a highly functional, fast-performing server, allowing for easy fine-tuning and control.

Assumption made -> dateAdded field is already populated with a time-stamp. The server does not need to assign the current date-time to the model.





