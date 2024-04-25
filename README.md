# Go Bank API

This project is a simple banking system implemented in Go. It uses the Gin framework for handling HTTP requests and PostgreSQL for data persistence.

The system supports user registration and login, and authenticated users can perform various operations related to their bank accounts and transactions.

Here are some of the main features:

- User registration and login: Users can register with a username and password, and log in to their account.

- Account management: Users can create new bank accounts, view details of their existing accounts, and delete accounts.

- Transaction management: Users can deposit money into an account, withdraw money from an account, and transfer money between accounts.

## Getting Started

To get started with this project, clone the repository to your local machine.

```console
git clone https://github.com/denny-kurniawan/go-bank.git
```

## Prerequisites

You need to have Go installed on your machine. You can download it from the [official Go website](https://go.dev/dl/).

## Installation

After cloning the repository, navigate to the project directory and download the required dependencies:

```console
go mod download
```

## Running the Application

To run the application, use the `go run` command:

```console
go run main.go
```

## Project Struture

The project structure is organized into several packages, each responsible for a specific aspect of the system:

- `controllers`: Contains the HTTP handlers for the various endpoints.
- `database`: Contains code for database connection and migrations.
- `helpers`: Contains helper functions used across the project.
- `middleware`: Contains middleware functions for handling authentication and cross-cutting concerns.
- `repositories`: Contains code for interacting with the database.
- `routers`: Contains the routing configuration.
- `structs`: Contains the data structures used in the project.

## API Endpoints

The application provides the following endpoints:

## User Management `/users`

### Register User

This endpoint allows a new user to register.

Endpoint: `/register`

Method: `POST`

Request Body

```json
{
  "username": "string",
  "password": "string"
}
```

Parameters

- `username`: The desired username for the new user. It must be unique.
- `password`: The password for the new user. It will be hashed before being stored in the database.

Response

```json
{
  "code": 200,
  "data": {
    "id": "uint64",
    "username": "string",
    "created_at": "string",
    "updated_at": "string"
  },
  "message": "User created successfully"
}
```

### Login User

This endpoint allows a user to log in.

#### Endpoint: `/login`

Method: `POST`

Request Body

```json
{
  "username": "string",
  "password": "string"
}
```

Parameters

- `username`: The username of the user trying to log in.
- `password`: The password of the user trying to log in.

Response

```json
{
  "code": 200,
  "data": {
    "id": "uint64",
    "username": "string",
    "token": "string",
    "expires_at": "string"
  },
  "message": "Login successful"
}
```

### Change Password User

This endpoint allows a user to change the password.

#### Endpoint: `/change-password`

Method: `POST`

Request Body

```json
{
  "username": "string",
  "old_password": "string",
  "new_password": "string",
  "confirm_new_password": "string"
}
```

Parameters

- `username`: The username of the user trying to change their password.
- `old_password`: The current password of the user.
- `new_password`: The new password that the user wants to set.
- `confirm_new_password`: The confirmation of new password that the user wants to set.

Response

```json
{
  "code": 200,
  "data": null,
  "message": "Your password has been successfully updated"
}
```

### Delete User

This endpoint allows a user to delete the user.

#### Endpoint: `/`

Method: `DELETE`

Request Body

```json
{
  "username": "string",
  "password": "string"
}
```

Parameters

- `username`: The username of the user trying to delete their account.
- `password`: The current password of the user.

Response

```json
{
  "code": 200,
  "data": null,
  "message": "User deleted successfully"
}
```

## Account Management `/accounts`

### Create Account

This endpoint allows a new user to register an account.

Endpoint: `/`

Method: `POST`

Request Body

```json
{
  "user_id": "uint64"
}
```

Parameters

- `user_id`: The ID of the user for whom the account is being created.

Response

```json
{
  "code": 200,
  "data": {
    "id": "uint64",
    "user_id": "uint64",
    "account_no": "string",
    "balance": "float64",
    "created_at": "string",
    "updated_at": "string"
  },
  "message": "Account created successfully"
}
```

### Get Account By User ID

This endpoint allows a user to retrieve all their accounts.

Endpoint: `/:userID`

Method: `GET`

Parameters

- `user_id`: The ID of the user for whom the account is being created.

Response

```json
{
  "code": 200,
  "data": [
    {
      "id": "uint64",
      "user_id": "uint64",
      "account_no": "string",
      "balance": "float64",
      "created_at": "string",
      "updated_at": "string"
    },
    // more accounts
  ],
  "message": "Accounts retrieved successfully"
}
```

### Get Account Details

This endpoint allows a user to retrieve their account details along with the transaction history.

Endpoint: `/:userID/:accountNo`

Method: `GET`

Parameters

- `userID`: The ID of the user for whom the account is being retrieved.
- `accountNo`: The account number for which the details are being retrieved.

Response

```json
{
  "code": 200,
  "data": {
    "account_id": "uint64",
    "user_id": "uint64",
    "account_no": "string",
    "balance": "float64",
    "created_at": "string",
    "updated_at": "string",
    "transactions": [
      {
        "id": "uint64",
        "transaction_type": "string",
        "amount": "float64",
        "description": "string",
        "created_at": "string"
      },
      // more transactions
    ]
  },
  "message": "Account retrieved successfully"
}
```

### Delete Account

This endpoint allows a user to delete their account.

Endpoint: `/:userID/:accountNo`

Method: `DELETE`

Parameters

- `userID`: The ID of the user for whom the account is being deleted.
- `accountNo`: The account number which is being deleted.

Response

```json
{
  "code": 200,
  "data": null,
  "message": "Account deleted successfully"
}
```

### Delete Account By User ID

This endpoint allows a user to delete all their accounts.

Endpoint: `/:userID`

Method: `DELETE`

Parameters

- `userID`: The ID of the user for whom the accounts are being deleted.

Response

```json
{
  "code": 200,
  "data": null,
  "message": "Account deleted successfully"
}
```

## Transaction Management `/transactions`

### Deposit

This endpoint allows a user to deposit money into their account.

Endpoint: `/deposit`

Method: `POST`

Request Body

```json
{
  "account_no": "string",
  "amount": "float64",
  "description": "string"
}
```

Parameters

- `account_no`: The account number into which the deposit is being made.
- `amount`: The amount of money to be deposited into the account.
- `description`: Additional information about the deposit. (Optional)

Response

```json
{
  "code": 200,
  "data": {
    "id": "uint64",
    "account_no": "string",
    "transaction_type": "string",
    "amount": "float64",
    "description": "string",
    "created_at": "string"
  },
  "message": "Deposit successful"
}
```

### Withdraw

This endpoint allows a user to withdraw money from their account.

Endpoint: `/withdraw`

Method: `POST`

Request Body

```json
{
  "account_no": "string",
  "amount": "float64",
  "description": "string"
}
```

Parameters

- `account_no`: The account number into which the withdraw is being made.
- `amount`: The amount of money to be withdrawn from the account.
- `description`: Additional information about the withdrawal. (Optional)

Response

```json
{
  "code": 200,
  "data": {
    "id": "uint64",
    "account_no": "string",
    "transaction_type": "string",
    "amount": "float64",
    "description": "string",
    "created_at": "string"
  },
  "message": "Withdraw successful"
}
```

### Transfer

This endpoint allows a user to transfer money from their account to another account.

Endpoint: `/transfer`

Method: `POST`

Request Body

```json
{
  "from_account_no": "string",
  "to_account_no": "string",
  "amount": "float64",
  "description": "string"
}
```

Parameters

- `from_account_no`: The account number from which the transfer is being made.
- `to_account_no`: The account number to which the transfer is being made.
- `amount`: The amount of money to be transferred.
- `description`: Additional information about the transfer. (Optional)

Response

```json
{
  "code": 200,
  "data": {
    "from_account_no": "string",
    "to_account_no": "string",
    "amount": "float64",
    "description": "string"
  },
  "message": "Transfer successful"
}
```

## Built With

- [Go](https://go.dev/) - The programming language used.
- [Gin](https://gin-gonic.com/) - The web framework used.
- [PostgreSQL](https://www.postgresql.org/) - The SQL database used.
