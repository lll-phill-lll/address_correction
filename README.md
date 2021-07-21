# address_correction

# About:
Recieves Russian address string in arbitrary format, standardizes and returnes
standardized address with FIAS ФИАС id of this address. This solution can recognize errors
or misspellings in address (see in examples). Also it can parse non-trivial addresses


# Examples

## Non trivial:

Request: `обл Новосибирская г Новосибирск спуск Владимировский 12/1`

Response:
```
{
  "corrected_address": {
    "city": "новосибирск",
    "street_type": "спуск",
    "street_name": "владимировский",
    "house_number": "12/1",
    "korpus": ""
  },
  "fias": "bacb0cdf-826b-4bfe-994f-fbae4c5df299"
}
```


Request: `г Новосибирск ул 40 лет Комсомола 10`

Response:
```
{
  "corrected_address": {
    "city": "новосибирск",
    "street_type": "ул",
    "street_name": "40 лет комсомола",
    "house_number": "10",
    "korpus": ""
  },
  "fias": "9dd8a05f-37d5-43ba-87ab-e1ca7789aec5"
}
```

Request: `ул. Северо-Енисейская, 42`

Response:
```
{
  "corrected_address": {
    "city": "красноярск",
    "street_type": "ул",
    "street_name": "северо-енисейская",
    "house_number": "42",
    "korpus": ""
  },
  "fias": "3f0c2961-0e0e-4f8f-9458-9f17c4dcc764"
}
```

## Errors correction:

Request: `г НовАсибирск ул Беловежская 6/1`

Reaponse:
```
{
  "corrected_address": {
    "city": "новосибирск",
    "street_type": "ул",
    "street_name": "беловежская",
    "house_number": "6/1",
    "korpus": "NA"
  },
  "fias": "92fd061a-e03a-47ac-b363-319ec897ee90"
}
```


## Shorten address:

Request: `ул. С. Лазо, 20`

Response:
```
{
  "corrected_address": {
    "city": "красноярск",
    "street_type": "ул",
    "street_name": "сергея лазо",
    "house_number": "20",
    "korpus": ""
  },
  "fias": "366f41d0-0d74-4777-b594-b4d4be71fc2f"
}
```

Request: `Красноярский рабочий, 160/1`

Response:
```
{
  "corrected_address": {
    "city": "красноярск",
    "street_type": "пр-кт",
    "street_name": "им.газеты \"красноярский рабочий\"",
    "house_number": "160/1",
    "korpus": ""
  },
  "fias": "cee6338e-98ae-4c1c-9ec1-6d9fb2a6a2f2"
}
```

# To run on linux:
1. `cd binary`
2. `./linux`
This command will run the server with config file from `binary/config.yaml`.
In this file you can change port and fias data file. To see the format go to
`data/fias`

# Docker:
1. `docker build -t address_corrector:1.0 .`
2. `docker run -p <your_port>:8080 address_corrector:1.0`


# Example requests:
## Request:
```
curl --header "Content-Type: application/json" \
     --request POST \
     --data '{"initial_address":"Ярыгинская набережная, 7 ","city":"Красноярск"}' \
     http://127.0.0.1/correct
```

## Response:
```
{
  "corrected_address": {
    "city": "красноярск",
    "street_type": "наб",
    "street_name": "ярыгинская",
    "house_number": "7",
    "korpus": "ANY"
  },
  "fias": "af4c1833-feb7-41e8-aded-4ef01ab13182"
}

```

For more navigate to `examples` folder

# To test the solution
1. Run server and navigate to `tests/functional`. Follow the README there
