# address_correction

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
