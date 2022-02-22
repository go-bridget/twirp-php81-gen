<?php

namespace Twirp;

use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;

class SupportTasksHandler
{
	private SupportTasks $service;

	public function __construct(SupportTasks $service)
	{
		$this->service = $service;
	}

	public function Index(Request $request, Response $response, array $args): Response
	{
		$params = new Empty($request);
		$data = $this->service->Index($params);
		$response->getBody()->write(json_encode($data));
		return $response->withHeader('Content-Type', 'application/json');
	}
}
