import os
import requests
from dotenv import load_dotenv

load_dotenv()

url = f"{os.getenv('BASE_URL')}/rerank"

headers = {
    "Authorization": f"Bearer {os.getenv('LIARA_API_KEY')}",
    "Content-Type": "application/json",
}

payload = {
    "model": os.getenv("RERANK_MODEL_NAME"),

    "query": "And who is God?",

    "documents": [
        "God means love, purity, intimacy, friendship",
        "God is kind",
        "God loves us and we should love him too",
        "AI means Artificial Intelligence",
    ],

    "top_n": 3,
}

response = requests.post(
    url,
    json=payload,
    headers=headers,
)

print(response.json())