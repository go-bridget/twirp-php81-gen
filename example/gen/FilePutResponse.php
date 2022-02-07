<?php

namespace Twirp;

class FilePutResponse
{
	public function __construct(
		public mixed $file, // File
	) {}
}
