# Api em golang de Books

## Implementação do serviço Book

- Adicionado método service.CreateBook
- Adicionado método service.GetBooks
- Adicionado método service.GetBookByID
- Adicionado método service.UpdateBook
- Adicionao método service.DeleteBook

## Implementação do HandlersBook

- Adicionado método bookHandlers.CreateBook
- Adicionado método bookHandlers.GetBooks
- Adicionado método bookHandlers.GetBookByID
- Adicionado método bookHandlers.UpdateBook
- Adicionado método bookHandlers.DeleteBook

## Implementação de rotas

- ("GET /books", bookHandlers.GetBooks)
- ("POST /books", bookHandlers.CreateBook)
- ("GET /books/{id}", bookHandlers.GetBookByID)
- ("PUT /books/{id}", bookHandlers.UpdateBook)
- ("DELETE /books/{id}", bookHandlers.DeleteBook)

## Instalando banco de dados sqlite3 books.db

- _ "github.com/mattn/go-sqlite3"

## Testando os endpoints

### Create a new book
```
POST http://localhost:8080/books
Content-Type: application/json

{
  "title": "The Lord of the Rings",
  "author": "J.R.R. Tolkien",
  "genre": "Fantasy"
}
```

### Get all books
```
GET http://localhost:8080/books
Accept: application/json
```

### Get a book by id
```
GET http://localhost:8080/books/1
Accept: application/json
```

### Update a book
```
PUT http://localhost:8080/books/1
Content-Type: application/json

{
  "title": "The Lord of the Rings",
  "author": "J.R.R. Tolkien",
  "genre": "Fantasy OK"
} 
```

### Delete a book
```
DELETE http://localhost:8080/books/2
Accept: application/json
```

### Subindo servidor

```
go run cmd/gobook/main.go
```


Font: #fullcycle