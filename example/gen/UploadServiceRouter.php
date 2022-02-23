<?php

namespace Twirp;

use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;
use Slim\Routing\RouteCollectorProxy;

class UploadServiceRouter
{
	public function Mount(\Slim\App $app, string $serviceClass)
	{
		$app->post("/twirp/upload.UploadService/FilePut", $serviceClass . ":handleFilePut");
		$app->post("/twirp/upload.UploadService/FileGet", $serviceClass . ":handleFileGet");
		$app->post("/twirp/upload.UploadService/FileDelete", $serviceClass . ":handleFileDelete");
		$app->post("/upload/v1/file", $serviceClass . ":handleFilePut");
		$app->get("/upload/v1/file/{fileID}", $serviceClass . ":handleFileGet");
		$app->delete("/upload/v1/file/{fileID}", $serviceClass . ":handleFileDelete");
	}
}
