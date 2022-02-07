<?php

namespace Twirp;

class File
{
	public function __construct(
		public string $ID, // int64
		public string $userID, // int64
		public string $folderID, // int64
		public mixed $folder, // Folder
		public string $link,
		public string $name,
		public string $mimeType,
		public string $size, // int64
		public string $source,
		public string $keywords,
		public string $stamp,
	) {}
}
