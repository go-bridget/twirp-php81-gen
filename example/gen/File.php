<?php

namespace Twirp;

class File
{
	public function __construct(
		public ?int $ID = 0, // int64
		public ?int $userID = 0, // int64
		public ?int $folderID = 0, // int64
		public ?Folder $folder = null,
		public ?string $link = "",
		public ?string $name = "",
		public ?string $mimeType = "",
		public ?int $size = 0, // int64
		public ?string $source = "",
		public ?string $keywords = "",
		public ?string $stamp = "",
	) {}
}
