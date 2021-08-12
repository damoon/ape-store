<?php

// Test this using following command
// php -S localhost:8080 ./graphql.php &
// curl http://localhost:8080 -d '{"query": "query { echo(message: \"Hello World\") }" }'
// curl http://localhost:8080 -d '{"query": "mutation { sum(x: 2, y: 2) }" }'
require_once 'vendor/autoload.php';

use GraphQL\GraphQL;
use GraphQL\Utils\BuildSchema;

try {
    $originalSchema = BuildSchema::build(file_get_contents('schema.graphqls'));

    $federation = new \PascalDeVink\GraphQLFederation\Federation();
    $schema = $federation->extendSchema($originalSchema);

    $rootValue = include 'rootvalue.php';
    $rootValue = $federation->addResolversToRootValue($rootValue);

    $rawInput = file_get_contents('php://input');
    $input = json_decode($rawInput, true);
    $query = $input['query'];
    $variableValues = isset($input['variables']) ? $input['variables'] : null;

    $result = GraphQL::executeQuery($schema, $query, $rootValue, null, $variableValues);
} catch (\Exception $e) {
    $result = [
        'error' => [
            'message' => $e->getMessage()
        ]
    ];
}

header('Content-Type: application/json; charset=UTF-8');
echo json_encode($result);
