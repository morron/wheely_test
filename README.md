# Wheely Test

## Run

* ```cp .env.example .env```
* add [Google API key](https://developers.google.com/maps/documentation/directions/get-api-key) in .env file to APIKEY
* ```docker-compose build```
* ```docker-compose up```

```
curl -X GET \
  http://localhost:2001/calc \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'postman-token: 9a8873e1-cfd9-b82c-388c-522d045f39f4' \
  -d '{
        "origin": {
                "lat": 55.8041983,
                "lng": 37.5831677
        },
        "destination": {
                "lat": 55.9663444,
                "lng": 37.4159007
        }
}'
```
