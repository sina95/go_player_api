# Go Player Balance API

## Overview

This is a basic Go application that provides a simple API for getting a player's balance. It demonstrates how to set up a RESTful API using Go.

## Features

- **Get player balance:** An endpoint to retrieve the balance of a player by their username.
- **Authentication:** Header-based authentication to secure the API endpoints.

## Installation and Setup
First, you need to have Go installed on your local machine. If you haven't installed Go yet, please follow [LINK](https://go.dev/doc/install). After that:
1. **Clone the Repository:**
   - Clone the repository to your local environment.
   ```bash
   git clone https://github.com/sina95/go_player_api.git
   ```
2. **Navigate to the Project Directory:**
   ```bash
   cd go_player_api
   ```
3. **Run the Application:**
   - Build and start the Go application.
  ```bash
  go run cmd/api/main.go
  ```
## Usage
**Authentication:**
- Include the Authorization header in your requests.
- Header Format: `Authorization: <credentials>`
- Example Header: `Authorization: dXNlcm5hbWU6cGFzc3dvcmQ=`
  
**Get Player balance:**
- Endpoint: GET /account/balance/?username={username}
- Example Request:
```bash
curl -H "Authorization: your-secret-token" "http://localhost:8000/account/balance/?username=your-username"
```
- Working Request:
```bash
curl -H "Authorization: test" "http://localhost:8000/account/balance/?username=test"
```
- Example Response:
```json
{
  "username": "test",
  "balance": 2528
}
```
## Contact
**For questions or support, please contact [sinisa.madzar95@gmail.com](mailto:sinisa.madzar95@gmail.com?subject=[GitHub]%20Source%20GO%20PLAYER%20API).**

Happy coding!
