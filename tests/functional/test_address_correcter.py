import pandas as pd
import json
import requests
from termcolor import colored

# TODO refactor

addresses = pd.read_csv('corrected_krasn_100.csv')

found = 0
found_correct = 0

for i in range(len(addresses)):
    print('-------')
    row = addresses.loc[[i]]
    address = row['address']
    fixed_address = row['fixed_address']
    fias = row['fias']
    fias = str(fias).split('\n')[0].split()[1]

    krasn = "Красноярск"
    address = str(address).split('\n')[0]
    ind = address.find(krasn)
    address = address[ind + len(krasn):]
    print(i, colored(address, 'yellow'))
    data = json.dumps({'initial_address': address, 'city': 'Красноярск'})
    headers = {'Content-type': 'application/json', 'Accept': 'text/plain'}

    url = 'http://localhost:8080/correct'
    response = requests.post(url, data.encode('utf-8'), headers=headers)
    print(json.loads(response.text)['corrected_address'])
    if json.loads(response.text)['fias'] != 'Not found':
        if fias == json.loads(response.text)['fias']:
            found_correct += 1
            print(colored('Correct', 'green'))
        else:
            print(colored('Incorrect', 'red'))
        found += 1
        print(colored('Found', 'green'))
    else:
        print(colored('Not found', 'red'))

print('Total found:', found, 'of:', len(addresses))
print('Total found correct:', found_correct, 'of:', len(addresses))
