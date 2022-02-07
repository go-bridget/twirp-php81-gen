<?php

namespace Twirp;

class FileDeleteRequest
{
	public function __construct(
		public string $sessionID,
		public string $fileID, // int64
		public string $authorization,
	) {}
}
