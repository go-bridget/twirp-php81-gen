<?php

namespace Twirp;

class Folder
{
	public function __construct(
		public ?int $ID = 0, // int64
		public ?int $selfID = 0, // int64
		public ?string $title = "",
		public ?bool $HasRules = false,
	) {}
}
