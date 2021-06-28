from dadata import Dadata
import pandas as pd
import os


token = os.environ['DADATA_TOKEN']
secret = os.environ['DADATA_SECRET']

res_dict = {'address': [], 'fixed_address': [], 'fias': []}

req_df = pd.read_csv('req_krasn_100.csv')


dadata = Dadata(token, secret)

try:
    for address in req_df['address_string']:
        full_address = 'Красноярск ' + address

        cleaned_data = dadata.clean(name="address", source=full_address)
        result = cleaned_data['result']
        fias_id = cleaned_data['fias_id']

        res_dict['address'].append(full_address)
        res_dict['fixed_address'].append(result)
        res_dict['fias'].append(fias_id)
except:
    pass

dadata.close()

res_df = pd.DataFrame(res_dict, columns = ['address', 'fixed_address', 'fias'])
res_df.to_csv('corrected_krasn_100.csv')
