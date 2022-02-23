<?php

namespace Twirp;

use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;

/** Support tasks service */
interface SupportTasks
{
	/** Index request */
	public function Index(Empty $req): Empty;
}
