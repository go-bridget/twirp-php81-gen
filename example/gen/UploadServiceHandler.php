<?php

namespace Twirp;

use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;

class UploadServiceHandler
{
	private UploadService $service;

	public function __construct(UploadService $service)
	{
		$this->service = $service;
	}

	public function handleFilePut(Request $request, Response $response, array $args): Response
	{
		$params = new FilePutRequest($request);
		$data = $this->service->FilePut($params);
		$response->getBody()->write(json_encode($data));
		return $response->withHeader('Content-Type', 'application/json');
	}

	public function handleFileGet(Request $request, Response $response, array $args): Response
	{
		$params = new FileGetRequest($request);
		$data = $this->service->FileGet($params);
		$response->getBody()->write(json_encode($data));
		return $response->withHeader('Content-Type', 'application/json');
	}

	public function handleFileDelete(Request $request, Response $response, array $args): Response
	{
		$params = new FileDeleteRequest($request);
		$data = $this->service->FileDelete($params);
		$response->getBody()->write(json_encode($data));
		return $response->withHeader('Content-Type', 'application/json');
	}
}
