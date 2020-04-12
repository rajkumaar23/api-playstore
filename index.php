<?php
/**
 * Copyright (c) 2020 | RAJKUMAR S (http://rajkumaar.co.in)
 */


use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;
use Slim\Exception\HttpNotFoundException;
use Slim\Factory\AppFactory;

require __DIR__ . '/vendor/autoload.php';
require 'DB.php';
require 'Utils.php';
require 'api.php';
define("CACHE_INTERVAL_MINUTES", 1440);

$app = AppFactory::create();

$app->get('/json', function (Request $request, Response $response) {
    $api = new API($request->getQueryParams()['id']);
    $response->getBody()->write(json_encode([
        'package' => $api->getPackageID(),
        'appVersion' => $api->getVersion(),
        'appSize' => $api->getSize(),
        'noOfInstalls' => $api->getInstalls(),
        'lastUpdated' => $api->getLastUpdated(),
        'rating' => $api->getRating(),
        'noOfUsersRated' => $api->getNoOfUsersRated(),
        'developer' => $api->getDeveloper()
    ]));
    return $response->withAddedHeader('Content-Type', 'application/json');
});

$app->get('/{type}', function (Request $request, Response $response, $args) {
    $api = new API($request->getQueryParams()['id']);
    switch ($args['type']) {
        case 'downloads' :
            $label = "Downloads";
            $message = $api->getInstalls();
            break;
        case 'package' :
            $label = "Package ID";
            $message = $api->getPackageID();
            break;
        case 'version' :
            $label = "Version";
            $message = $api->getVersion();
            break;
        case 'size' :
            $label = "App Size";
            $message = $api->getSize();
            break;
        case 'lastUpdated' :
            $label = "Last Updated On";
            $message = $api->getLastUpdated();
            break;
        case 'rating' :
            $label = "Rating";
            $message = $api->getRating();
            break;
        case 'developer' :
            $label = "Developer";
            $message = $api->getDeveloper();
            break;
        case 'noOfUsersRated' :
            $label = "No of users rated";
            $message = $api->getNoOfUsersRated();
            break;
        default :
            $label = "Badge Type";
            $message = "Invalid";
            break;
    }
    $response->getBody()->write(json_encode(['schemaVersion' => 1, 'label' => $label, 'message' => $message]));
    return $response->withAddedHeader('Content-Type', 'application/json');
});

$customErrorHandler = function (
    Request $request,
    Throwable $exception,
    bool $displayErrorDetails,
    bool $logErrors,
    bool $logErrorDetails
) use ($app) {
    $payload = ['error' => $exception instanceof HttpNotFoundException ? $exception->getDescription() : $exception->getMessage()];

    $response = $app->getResponseFactory()->createResponse();
    $response->getBody()->write(
        json_encode($payload, JSON_UNESCAPED_UNICODE)
    );

    return $response->withStatus(400)->withAddedHeader('Content-Type', 'application/json');
};

$errorMiddleware = $app->addErrorMiddleware(false, true, true);
$errorMiddleware->setDefaultErrorHandler($customErrorHandler);
$app->run();


