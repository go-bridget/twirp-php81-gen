<?php

namespace Twirp;

class FilePutResponse
{
	public function __construct(
		public ?File $file = null,
	) {}
}
