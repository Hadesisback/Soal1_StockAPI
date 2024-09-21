
[![StockAPI](https://github.com/Hadesisback/Soal1_StockAPI/actions/workflows/go.yml/badge.svg)](https://github.com/Hadesisback/Soal1_StockAPI/actions/workflows/go.yml)

# Soal1 - CRUD 



## Usage


### Create a stock
#### POST /stocks



Windows (CMD)
```
curl -X POST http://localhost:8000/stocks -H "Content-Type: application/json" -d "{\"name\": \"Apple Inc.\", \"code\": \"AAPL\", \"price\": 150.00, \"frequency\": 100, \"volume\": 2000}"

```
Linux
```
curl -X POST http://localhost:8000/stocks -H "Content-Type: application/json" -d '{"name": "Apple Inc.", "code": "AAPL", "price": 150.00, "frequency": 100, "volume": 2000}'
```

### Get all stocks
#### GET /stocks
```
curl http://localhost:8000/stocks
```


###  Update a stock by ID
#### PUT /stocks/<id>
```
curl -X PUT http://localhost:8000/stocks/1 -H "Content-Type: application/json" -d '{"name": "Google", "code": "GOOG", "price": 250.00, "frequency": 200, "volume": 5000}'

```


### Delete a stock by ID
#### DELETE /stocks/<id>
```
curl -X DELETE http://localhost:8000/stocks/1
```
