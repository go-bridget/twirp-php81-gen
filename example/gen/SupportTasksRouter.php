<?php

namespace Twirp;

use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;
use Slim\Routing\RouteCollectorProxy;

class SupportTasksRouter
{
	public function Mount(\Slim\App $app, string $serviceClass)
	{
		$app->map(["GET"], "/", $serviceClass . ":handleIndex")->setName("Index");
	}
}
