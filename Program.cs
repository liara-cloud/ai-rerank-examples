using System.Net.Http.Headers;
using System.Text;
using System.Text.Json;
using DotNetEnv;

Env.Load();

var baseUrl = Environment.GetEnvironmentVariable("BASE_URL");
var apiKey = Environment.GetEnvironmentVariable("LIARA_API_KEY");
var modelName = Environment.GetEnvironmentVariable("RERANK_MODEL_NAME");

var url = $"{baseUrl}/rerank";

var payload = new
{
    model = modelName,

    query = "And who is God?",

    documents = new[]
    {
        "God means love, purity, intimacy, friendship",
        "God is kind",
        "God loves us and we should love him too",
        "AI means Artificial Intelligence"
    },

    top_n = 3
};

using var client = new HttpClient();

client.DefaultRequestHeaders.Authorization =
    new AuthenticationHeaderValue("Bearer", apiKey);

var json = JsonSerializer.Serialize(payload);

var content = new StringContent(
    json,
    Encoding.UTF8,
    "application/json"
);

var response = await client.PostAsync(url, content);

var result = await response.Content.ReadAsStringAsync();

Console.WriteLine(result);