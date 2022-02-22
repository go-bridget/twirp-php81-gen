<?php

namespace Twirp;

use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;
use Slim\Routing\RouteCollectorProxy;

class UploadServiceRouter
{
	public function Mount(\Slim\App $app, string $serviceClass)
	{
		$app->group("/upload/v1/file", function (RouteCollectorProxy $group) use ($serviceClass)
		{
			$group->map(["POST"], "", function (Request $request, Response $response, array $args) use ($serviceClass) {
				$service = new $serviceClass;
				$handler = new UploadServiceHandler($service);
				return $handler->FilePut($request, $response, $args);
			})->setName("FilePut");
			$group->map(["GET"], "/{fileID}", function (Request $request, Response $response, array $args) use ($serviceClass) {
				$service = new $serviceClass;
				$handler = new UploadServiceHandler($service);
				return $handler->FileGet($request, $response, $args);
			})->setName("FileGet");
			$group->map(["DELETE"], "/{fileID}", function (Request $request, Response $response, array $args) use ($serviceClass) {
				$service = new $serviceClass;
				$handler = new UploadServiceHandler($service);
				return $handler->FileDelete($request, $response, $args);
			})->setName("FileDelete");
		});
	}
}
