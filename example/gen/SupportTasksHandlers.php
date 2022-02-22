<?php

namespace Twirp;

use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;

abstract class SupportTasksHandlers implements SupportTasks
{

	/**  */
	public function handleIndex(Request $request, Response $response, array $args): Response
	{
		$params = new Empty($request);
		$data = $this->Index($params);
		$response->getBody()->write(json_encode($data));
		return $response->withHeader('Content-Type', 'application/json');
	}
}
