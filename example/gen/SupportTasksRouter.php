<?php

namespace Twirp;

use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;
use Slim\Routing\RouteCollectorProxy;

class SupportTasksRouter
{
	public function Mount(\Slim\App $app, string $serviceClass)
	{
		$requestIndex = function(Request $request, Response $response, array $args) {
			$service = new $serviceClass;
			$service->Index($request, $response, $args);
		};
		$app->map(["GET"], "/", $requestIndex)->setName("Index");
	}
}
