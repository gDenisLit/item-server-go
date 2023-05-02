
# Item server - Fiber Go

This is a production-ready Go server built with the Fiber framework. It provides RESTful API endpoints for managing items, users, and authentication. This is the Fiber Go version of the Item App - A generic e2e app im building with various technologies, using Vue, React, Express, Go and Pyhton.

## Tech Stack

**Server:** go, fiber, securecookie, crypto, mongodb,


## Run Locally

Clone the project

```bash
  git clone https://github.com/gDenisLit/item-server-ts.git
```

Go to the project directory

```bash
  cd item-server-ts
```

Install dependencies

```bash
  go mod tidy
```

Start the server in dev enviroment

```bash
  make run
```

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`ATLAS_URL`
Connection URI to mongo

`DB_NAME`
Database name

`CRYPTER_KEY`
Base 64 secret key

## API Reference

#### Get all items

```http
  GET /api/item
```

#### Get item

```http
  GET /api/item/${id}
```

#### Add Item

```http
  POST /api/item
```
#### Update Item

```http
  PUT /api/item
```
#### Remove Item
```http
  DELETE /api/item/${id}
```


## Authors

- [@gDenislit](https://www.github.com/gDenislit)
