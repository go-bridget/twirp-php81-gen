<?php

namespace Twirp;

class FileGetRequest
{
	public function __construct(
		public string $sessionID,
		public string $fileID, // int64
	) {}
}
