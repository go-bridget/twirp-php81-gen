<?php

namespace Twirp;

use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;

/** UploadService handles file uploads */
interface UploadService
{
	/** Upload a file */
	public function FilePut(FilePutRequest $req): FilePutResponse;

	/** Retrieve a file */
	public function FileGet(FileGetRequest $req): FileGetResponse;

	/** Delete a file */
	public function FileDelete(FileDeleteRequest $req): FileDeleteResponse;
}
