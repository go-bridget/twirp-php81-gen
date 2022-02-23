<?php

namespace Twirp;

use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;

abstract class UploadServiceHandlers implements UploadService
{
	/** Upload a file */
	public function handleFilePut(Request $request, Response $response, array $args): Response
	{
		$params = new FilePutRequest($request);
		$data = $this->FilePut($params);
		$response->getBody()->write(json_encode($data));
		return $response->withHeader('Content-Type', 'application/json');
	}
	/** Retrieve a file */
	public function handleFileGet(Request $request, Response $response, array $args): Response
	{
		$params = new FileGetRequest($request);
		$data = $this->FileGet($params);
		$response->getBody()->write(json_encode($data));
		return $response->withHeader('Content-Type', 'application/json');
	}
	/** Delete a file */
	public function handleFileDelete(Request $request, Response $response, array $args): Response
	{
		$params = new FileDeleteRequest($request);
		$data = $this->FileDelete($params);
		$response->getBody()->write(json_encode($data));
		return $response->withHeader('Content-Type', 'application/json');
	}
}
