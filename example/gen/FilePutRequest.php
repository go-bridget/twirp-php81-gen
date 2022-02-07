<?php

namespace Twirp;

class FilePutRequest
{
	public function __construct(
		public string $sessionID,
		public string $folderID, // int64
		public string $authorization,
		public string $name,
		public string $mimeType,
		public string $size, // int64
		public mixed $body, // bytes
		public string $source,
		public string $keywords,
	) {}
}
