<?php

namespace Twirp;

class Folder
{
	public function __construct(
		public string $ID, // int64
		public string $selfID, // int64
		public string $title,
		public bool $HasRules,
	) {}
}
