The project is a simple expense tracker that allows users to add, delete, and view their expenses. The project is based on the following stack:

- Go
- PostGres
- Redis

## Getting Started

First, clone the repository:

```bash
git clone https://github.com/khushal87/expense-tracker-backend-go
```

Next, navigate to the project directory:

```bash
cd expense-tracker-backend-go
```

Create a `.env` file in the root of the project and add the following environment variables:

```bash
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
DB_PORT=your_db_port
DB_HOST=your_db_host
REDIS_ADDRESS=your_redis_address
REDIS_PASSWORD=your_redis_password
```

Next, run the following command to start the project:

- Install the dependencies:

```bash
go mod tidy
```

- Run the project:

```bash
go run main.go
```

The project will be running on `http://localhost:8080`.

## API Endpoints

The project has the following API endpoints:

- `GET /sources`: Get all expenses
- `POST /sources/create`: Add a new expense
- `DELETE /sources/delete/{id}`: Delete an expense
- `GET /transactions`: Get all transactions
- `POST /transactions/create`: Add a new transaction
- `DELETE /transactions/delete/{id}`: Delete a transaction

## Future Improvements

- Add authentication
- Add tests
