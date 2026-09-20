<?php

// composer require vlucas/phpdotenv
require __DIR__ . '/vendor/autoload.php';

$dotenv = Dotenv\Dotenv::createImmutable(__DIR__);
$dotenv->load();

$url = $_ENV["BASE_URL"] . "/rerank";

$payload = [
    "model" => $_ENV["RERANK_MODEL_NAME"],

    "query" => "And who is God?",

    "documents" => [
        "God means love, purity, intimacy, friendship",
        "God is kind",
        "God loves us and we should love him too",
        "AI means Artificial Intelligence",
    ],

    "top_n" => 3,
];

$ch = curl_init($url);

curl_setopt_array($ch, [
    CURLOPT_POST => true,
    CURLOPT_RETURNTRANSFER => true,

    CURLOPT_HTTPHEADER => [
        "Authorization: Bearer " . $_ENV["LIARA_API_KEY"],
        "Content-Type: application/json",
    ],

    CURLOPT_POSTFIELDS => json_encode(
        $payload,
        JSON_UNESCAPED_UNICODE
    ),
]);

$response = curl_exec($ch);

if ($response === false) {
    echo "cURL Error: " . curl_error($ch);
} else {
    echo $response;
}

curl_close($ch);