<?php

namespace Twirp;

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
