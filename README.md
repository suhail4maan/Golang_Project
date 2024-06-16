# Golang_Project

> Overview
Developed a RESTful API to manage a movie watchlist, leveraging Golang for backend development and MongoDB as the database.

> Technologies Used
Golang: For creating a robust and efficient backend.
MongoDB: As the NoSQL database for flexible and scalable data storage.
Gorilla/Mux: For routing and handling HTTP requests.
Project Structure
Model:

> Defined a struct to represent movie entities with the following fields:
ID (string): Unique identifier for each movie.
MovieName (string): Name of the movie.
Watched (bool): Boolean flag indicating whether the movie has been watched.
Router:

> Utilized Gorilla/Mux to set up API endpoints.
Configured routes to handle HTTP methods (GET, POST, PUT, DELETE) for interacting with the movie watchlist.
Controller:

> Implemented controller functions to handle the business logic:
Insert: Add a new movie to the watchlist.
Update: Modify existing movie details.
Delete: Remove a movie from the watchlist.
Get: Retrieve movies from the watchlist.
Each function interacts with the MongoDB database to perform the required operations.
API Endpoints
GET /movies: Retrieve the list of all movies in the watchlist.
POST /movies: Add a new movie to the watchlist.
PUT /movies/{id}: Update the details of an existing movie.
DELETE /movies/{id}: Remove a movie from the watchlist.
Key Features
Efficiently managed CRUD operations with MongoDB using Go.
Implemented RESTful principles for scalable and maintainable API design.
Used Gorilla/Mux for robust routing and middleware support.
Designed a modular architecture with clear separation of concerns (model, router, controller).
