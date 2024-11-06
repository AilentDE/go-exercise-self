import os
import requests
from dotenv import load_dotenv

load_dotenv()

base_url = os.getenv("BASE_URL")

query = {"name": "test", "address": "no place"}

r = requests.get(f"{base_url}/query", params=query)
# r.raise_for_status()
print(r.status_code, "-", r.json() if r.status_code not in (404, 500) else r.text)
