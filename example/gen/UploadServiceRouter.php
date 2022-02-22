<?php

namespace Twirp;

use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;
use Slim\Routing\RouteCollectorProxy;

class UploadServiceRouter
{
	public function Mount(\Slim\App $app, string $serviceClass)
	{
		$app->group("/upload/v1/file", function (RouteCollectorProxy $group) use ($serviceClass) {
			$group->map(["POST"], "", $serviceClass . ":FilePut")->setName("FilePut");
			$group->map(["GET"], "/{fileID}", $serviceClass . ":FileGet")->setName("FileGet");
			$group->map(["DELETE"], "/{fileID}", $serviceClass . ":FileDelete")->setName("FileDelete");
		});
	}
}
