# GO CRUD

A HTTP server built using Go for managing albums.

[Preview](./preview.gif)

## CREATE

```sh
curl -X POST http://localhost:8080/albums \
  -H "Content-Type: application/json" \
  -d '{
        "id": "4",
        "title": "Future Nostalgia",
        "artist": "Dua Lipa",
        "price": 22.50
      }'
```

## READ

```sh
curl http://localhost:8080/albums
```

## UPDATE

```sh
curl -X PUT http://localhost:8080/albums/2 \
  -H "Content-Type: application/json" \
  -d '{
        "title": "Nothing Was the Same (Deluxe)",
        "artist": "Drake",
        "price": 19.99
      }'
```

## DESTROY

```sh
curl -X DELETE http://localhost:8080/albums/2
```

## Resources

- https://go.dev/doc/tutorial/web-service-gin
