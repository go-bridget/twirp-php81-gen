<?php

namespace Twirp;

class FileDeleteRequest
{
	public function __construct(
		public ?string $sessionID = "",
		public ?int $fileID = 0, // int64
		public ?string $authorization = "",
	) {}
}
