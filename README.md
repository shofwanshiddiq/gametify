# Gametify
This is a RESTful API built using Golang with Gin frameworks, GORM as ORM for MySQL. 

### Apps Business Process 
* User registration, login, and update with email and password authorization
* Upload and update user's profile photo
* Find rooms, places, package type based on user's needs
* See available rooms based on time constraint
* Perform booking, finish booking, and post rating
* Check booking history

### Technical Features
* User registration with SHA-256 password hashing
* User login with JWT-based authentication
* Middleware for protected routes
* User role authorization using Middleware
* MySQL database connection using .env

# Technologies
![Golang](https://img.shields.io/badge/golang-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)  ![REST API](https://img.shields.io/badge/restapi-%23000000.svg?style=for-the-badge&logo=swagger&logoColor=white)   ![MySQL](https://img.shields.io/badge/mysql-%234479A1.svg?style=for-the-badge&logo=mysql&logoColor=white)

# Models
![image](https://github.com/user-attachments/assets/75154ddd-9c9e-4a8f-b1ac-5432c9114db3)

# API Endpoints Documentation

This document provides an overview of the API endpoints, their methods, and functionality.
| Method     | API Endpoint                        | Description                                      | Pages              |
|------------|--------------------------------------|--------------------------------------------------|--------------------|
| **POST**   | `/api/auth/register`                 | Register a new user                              | Authentication     |
| **POST**   | `/api/auth/login`                    | Authenticate user and return JWT token           | Authentication     |
| **GET**    | `/api/users`                         | Get all users                                    | User Management    |
| **GET**    | `/api/users/:id`                     | Get user by ID                                   | User Management    |
| **GET**    | `/api/users/profile`                 | Get logged-in user profile                       | User Profile       |
| **PUT**    | `/api/users/:id`                     | Update user by ID                                | User Management    |
| **DELETE** | `/api/users/:id`                     | Delete user by ID                                | User Management    |
| **POST**   | `/api/users/profile-picture`         | Upload user profile picture                      | User Profile       |
| **GET**    | `/api/users/profile-picture`         | Get user profile picture                         | User Profile       |
| **GET**    | `/api/places`                        | Get all places                                   | Place Management   |
| **GET**    | `/api/places/:id`                    | Get place by ID                                  | Place Management   |
| **GET**    | `/api/rooms`                         | Get all rooms                                    | Room Management    |
| **GET**    | `/api/rooms/:id`                     | Get room by ID                                   | Room Management    |
| **GET**    | `/api/rooms/place/:place_id`          | Get rooms by place ID                            | Room Management    |
| **GET**    | `/api/rooms/console/:console_type`   | Get rooms by console type                        | Room Management    |
| **GET**    | `/api/rooms/console`                 | Get available console types                      | Room Management    |
| **GET**    | `/api/bookings`                      | Get all bookings                                 | Booking Management |
| **GET**    | `/api/bookings/:id`                  | Get booking by ID                                | Booking Management |
| **POST**   | `/api/bookings`                      | Create a new booking                             | Booking Management |
| **PATCH**  | `/api/bookings/:id`                  | Update booking status                            | Booking Management |
| **POST**   | `/api/bookings/:id/rate`              | Post rating for a booking                        | Booking Management |
| **GET**    | `/api/bookings/room/:room_id`         | Get average room rating                          | Booking Management |
| **GET**    | `/api/bookings/place/:place_id`       | Get average place rating        | Booking Management |

# Front End Mockup
<img src="https://github.com/user-attachments/assets/4d9600c8-021d-46d3-b28d-db405af8ddd1" alt="Capture9" width="150" style="border-radius: 20px;">
<img src="https://github.com/user-attachments/assets/98448b2d-5fd8-4960-81ba-2cf7c5c141d1" alt="Capture9" width="150" style="border-radius: 20px;">
<img src="https://github.com/user-attachments/assets/855eff84-e7bd-4042-be1e-de609e4af423" alt="Capture9" width="150" style="border-radius: 20px;">
<img src="https://github.com/user-attachments/assets/04c6092e-d92a-43e2-bd58-24bc264e6ca2" alt="Capture9" width="150" style="border-radius: 20px;">
<img src="https://github.com/user-attachments/assets/751f1043-b252-45c5-97cc-5e68e0de8117" alt="Capture9" width="150" style="border-radius: 20px;">

## Initialization

Follow these steps to set up the project:

### 1. Initialize the Go Module
Run the following command in the project directory:
```sh
git clone https://github.com/shofwanshiddiq/gametify
go mod init gametify
```

### 2. Install Dependencies
Install the required packages:

```sh
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u github.com/joho/godotenv
go get github.com/golang-jwt/jwt/v5
```

### 3. Configure Database
Create a .env file in the root directory and add your database credentials:

```env
DB_USER=root
DB_PASSWORD=password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=dbname

JWT_SECRET_KEY=super_secret_key
JWT_EXPIRATION_IN=24h
```

### 4. Run API
```sh
go run main.go
```

# API Testing
Use this Postman documentation for endpoint testing
[Here](https://.postman.co/workspace/Dibimbing-Golang~b5255e78-e541-48aa-a48f-df1842830c9c/collection/31117152-ce3df710-928d-4b76-935f-b9035e058c2f?action=share&creator=31117152
)


