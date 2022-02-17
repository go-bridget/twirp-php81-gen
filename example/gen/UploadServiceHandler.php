<?php

namespace Twirp;

use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;

class UploadServiceHandler
{
	public function __construct(
		public $service = new UploadService;
	) {}

	public function FilePut(Request $request, Response $response, array $args): Response
	{
		$serviceRequest = new FilePutRequest($request);
		$response->writeJSON($this->service->FilePut($serviceRequest));
	}

	public function FileGet(Request $request, Response $response, array $args): Response
	{
		$serviceRequest = new FileGetRequest($request);
		$response->writeJSON($this->service->FileGet($serviceRequest));
	}

	public function FileDelete(Request $request, Response $response, array $args): Response
	{
		$serviceRequest = new FileDeleteRequest($request);
		$response->writeJSON($this->service->FileDelete($serviceRequest));
	}
}
