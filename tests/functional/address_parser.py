import pandas as pd
import json
import requests
from termcolor import colored

# TODO refactor

addresses = pd.read_csv('req_novosib.csv')

found = 0
found_correct = 0

for i in range(len(addresses)):
    print('-------')

    address = addresses['C_Address'][i]
    print(i, colored(address, 'yellow'))
    data = json.dumps({'initial_address': address, 'city': ''})
    headers = {'Content-type': 'application/json', 'Accept': 'text/plain'}

    # url = 'http://35.222.114.24/correct'
    url = 'http://localhost:8080/correct'
    response = requests.post(url, data.encode('utf-8'), headers=headers)
    print(json.loads(response.text)['corrected_address'])
    if json.loads(response.text)['fias'] != 'Not Found':
        print("Fias:", json.loads(response.text)['fias'])
        found += 1
        print(colored('Found', 'green'))
    else:
        print(colored('Not found', 'red'))

print('Total found:', found, 'of:', len(addresses))
# print('Total found correct:', found_correct, 'of:', len(addresses))
