<?php

namespace Twirp;

class FileGetResponse
{
	public function __construct(
		public mixed $file, // File
	) {}
}
