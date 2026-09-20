const response = await fetch("<BASE_URL>/rerank", {
  method: "POST",
  headers: {
    Authorization: "Bearer <LIARA_API_KEY>",
    "Content-Type": "application/json",
  },
  body: JSON.stringify({
    model: "<RERANK_MODEL_NAME>",
    query: "And who is God?",

    documents: [

        "God means love, purity, intimacy, friendship",
        "God is kind",
        "God loves us and we should love him too",
        "AI means Artificial Intelligence",

    ],
    top_n: 3,
  }),
});

const data = await response.json();
console.log(data);