<?php

namespace Twirp;

class UploadServiceRouter
{
	public function Mount(\Slim\App $app)
	{
		$app->map('post', '/upload/v1/file', '\Twirp\UploadServiceHandler:FilePut')->setName('FilePut');
		$app->map('get', '/upload/v1/file/{fileID}', '\Twirp\UploadServiceHandler:FileGet')->setName('FileGet');
		$app->map('delete', '/upload/v1/file/{fileID}', '\Twirp\UploadServiceHandler:FileDelete')->setName('FileDelete');
	}
}
