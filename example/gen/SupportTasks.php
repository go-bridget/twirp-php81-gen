<?php

namespace Twirp;

use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;

/**  */
interface SupportTasks
{
	/**  */
	public function Index(Empty $req): Empty;
}
