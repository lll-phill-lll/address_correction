# Here you can place file with addresses and fias numbers.

## Format:
csv file with following columns:
HOUSEGUID,CITY,STREETTYPE,FORMALNAME,HOUSENUM,KORPUS,HOUSETYPE,STROENIE,STATUS
Where:
HOUSEGUID - fias id
CITY - city name. In case of long string only last word will be taken. (split by space)
STREETTYPE - type of street: street, boulevard. Possible values are: ['ул', 'пер', 'пл', 'мкр', 'проезд', 'ост-в', 'тракт', 'пр-кт', 'б-р', 'тер', 'ш', 'наб', 'снт']
FORMALNAME - name of the street
HOUSENUM - number of house
KORPUS - korpus of house
HOUSETYPE - type of house (currently unused)
STROENIE - the same as korpus
STATUS - type of building (currently unused)

All the fields are strings
