<?php

namespace Twirp;

class FilePutRequest
{
	public function __construct(
		public ?string $sessionID = "",
		public ?int $folderID = 0, // int64
		public ?string $authorization = "",
		public ?string $name = "",
		public ?string $mimeType = "",
		public ?int $size = 0, // int64
		public ?string $body = "", // byte
		public ?string $source = "",
		public ?string $keywords = "",
	) {}
}
