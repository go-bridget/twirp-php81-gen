<?php

namespace Twirp;

class UploadServiceRouter
{
	public function Mount(\Slim\App $app, UploadService $service)
	{
		$app->map('post', '/upload/v1/file', [$service, 'FilePut'])->setName('FilePut');
		$app->map('get', '/upload/v1/file/{fileID}', [$service, 'FileGet'])->setName('FileGet');
		$app->map('delete', '/upload/v1/file/{fileID}', [$service, 'FileDelete'])->setName('FileDelete');
	}
}
