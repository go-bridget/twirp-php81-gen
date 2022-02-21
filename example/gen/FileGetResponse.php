<?php

namespace Twirp;

class FileGetResponse
{
	public function __construct(
		public ?File $file = null,
	) {}
}
