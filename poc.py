import requests
import json

LINK = 'https://cdn.jsdelivr.net/gh/fawazahmed0/currency-api@1/latest/'

# get all currecies and their shorthands
all_currencies_response = requests.get(LINK + 'currencies.json')
currencies = json.loads(all_currencies_response.text)

# now compare 2 currencies, the eur and usd just to check that the equality holds
usd_conversions = requests.get(LINK + 'currencies/usd.json')
eur_conversions = requests.get(LINK + 'currencies/eur.json')

usd_conversions = json.loads(usd_conversions.text)
eur_conversions = json.loads(eur_conversions.text)

print(usd_conversions['usd']['eur'], 1/ float(eur_conversions['eur']['usd']))
