<?php

namespace Twirp;

class UploadServiceRouter
{
	public function Mount(\Slim\App $app, string $serviceClass)
	{
		$app->group("/upload/v1/file", function (RouteCollectorProxy $group)
		{
			$service = new $serviceClass;
			$app->map(["POST"], "", [$service, "FilePut"])->setName("FilePut");
			$app->map(["GET"], "/{fileID}", [$service, "FileGet"])->setName("FileGet");
			$app->map(["DELETE"], "/{fileID}", [$service, "FileDelete"])->setName("FileDelete");
		}
	}
}
