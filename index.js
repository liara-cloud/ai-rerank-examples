// npm install dotenv
require("dotenv").config();

async function main() {
  try {
    const response = await fetch(`${process.env.BASE_URL}/rerank`, {
      method: "POST",

      headers: {
        Authorization: `Bearer ${process.env.LIARA_API_KEY}`,
        "Content-Type": "application/json",
      },

      body: JSON.stringify({
        model: process.env.RERANK_MODEL_NAME,

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

    if (!response.ok) {
      const error = await response.text();
      throw new Error(`Request failed: ${response.status} - ${error}`);
    }

    const data = await response.json();

    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

main();