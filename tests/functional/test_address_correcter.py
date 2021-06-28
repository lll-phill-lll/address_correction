import pandas as pd
import json
import requests

# TODO refactor

addresses = pd.read_csv('corrected_krasn_100.csv')

for i in range(3):
    row = addresses.loc[[i]]
    address = row['address']
    fixed_address = row['fixed_address']
    fias = row['fias']

    krasn = "Красноярск"
    address = str(address).split('\n')[0]
    ind = address.find(krasn)
    address = address[ind + len(krasn):]
    print(address)
    data = json.dumps({'initial_address': address, 'city': 'Красноярск'})
    headers = {'Content-type': 'application/json', 'Accept': 'text/plain'}

    url = 'http://localhost:8080/correct'
    response = requests.post(url, data.encode('utf-8'), headers=headers)
    print(response)
    print(response.text)
