# address_correction

# To run:
1. `cd cmd/corrector`
2. `go build`
3. `./corrector`

# Example:
curl --header "Content-Type: application/json" \\n  --request POST \\n  --data '{"username":"xyz","password":"xyz"}' \\n  http://localhost:8080/correct
