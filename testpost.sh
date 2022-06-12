#!/bin/sh

curl -X POST --data '{"id": 7, "title": "test title", "author": "test author", "publicationYear": 2000, "genre": "horror"}' http://localhost:5000/books
