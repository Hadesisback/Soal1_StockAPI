
[![StockAPI](https://github.com/Hadesisback/Soal1_StockAPI/actions/workflows/go.yml/badge.svg)](https://github.com/Hadesisback/Soal1_StockAPI/actions/workflows/go.yml)

# Soal1 - CRUD 



## Usage


### Create a stock
#### POST /stocks



Windows (CMD)
```
curl -X POST http://febrian.me:8089/stocks -H "Content-Type: application/json" -d "{\"name\": \"Apple Inc.\", \"code\": \"AAPL\", \"price\": 150.00, \"frequency\": 100, \"volume\": 2000}"

```
Linux
```
curl -X POST http://febrian.me:8089/stocks -H "Content-Type: application/json" -d '{"name": "Apple Inc.", "code": "AAPL", "price": 150.00, "frequency": 100, "volume": 2000}'
```

###  Get a stock by ID
#### GET /stocks
```
curl http://febrian.me:8089/stocks/1
```


### Get all stocks
#### GET /stocks
```
curl http://febrian.me:8089/stocks
```


###  Update a stock by ID
#### PUT /stocks/<id>
```
curl -X PUT http://febrian.me:8089/stocks/1 -H "Content-Type: application/json" -d '{"name": "Google", "code": "GOOG", "price": 250.00, "frequency": 200, "volume": 5000}'

```


### Delete a stock by ID
#### DELETE /stocks/<id>
```
curl -X DELETE http://febrian.me:8089/stocks/1
```
