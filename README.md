# Ristek Open Recruitment Assignment - Web Development
[Task Assignment](https://docs.google.com/document/d/1NO3Qd4p4QVDM-eYQ0QzYsS9gPthA7LazR-BimuhE3tE/edit?usp=sharing)
## Features
- [x] Backend Level 1
- [x] Backend Level 2
- [x] Backend Level 3
- [x] frontend Level 1
- [x] frontend Level 2

## Prerequisites
Before starting, make sure you have installed:
- [Go](https://go.dev/doc/install) (latest version recommended)
- [Git](https://git-scm.com/)
- [Node.js](https://nodejs.org/) and [npm](https://www.npmjs.com/) (for frontend)

## Installation

1. Clone this repository:
   ```sh
   git clone https://github.com/HeraldoArman/oprec-ristek.git
   cd oprec-ristek
   ```

2. Install dependencies for the backend:
   ```sh
   cd backend
   go mod tidy
   ```

3. Install dependencies for the frontend:
   ```sh
   cd ../frontend
   npm install
   ```

4. Configure the `.env` file for the backend based on the provided example (`.env.example`):
   ```sh
   cd ../backend
   cp .env.example .env
   ```
   Then edit it as needed.



## Running the Project

1. Start the backend:
   ```sh
   cd backend
   go run cmd/main.go
   ```
   The backend will run at `http://localhost:3000`.

2. Start the frontend in a new terminal:
   ```sh
   cd frontend
   npm run dev
   ```
   The frontend will run at `http://localhost:5173`.

## API Documentation
Access the documentation at:
```
https://documenter.getpostman.com/view/41269141/2sAYk7RPNS
```




