<?php declare(strict_types=1);

use Hyperf\Contract\StdoutLoggerInterface;
use Hyperf\HttpMessage\Stream\SwooleStream;
use Hyperf\HttpServer\Contract\RequestInterface;
use Hyperf\HttpServer\Contract\ResponseInterface;
use Hyperf\Nano\Factory\AppFactory;
use Hyperf\Utils\Codec\Json;
use Spiral\Goridge\Relay;
use Spiral\Goridge\RPC\RPC;
use Swoole\Http\Status;

require_once __DIR__ . '/vendor/autoload.php';

$app = AppFactory::createBase();

$app->addExceptionHandler(function (Throwable $err, \Psr\Http\Message\ResponseInterface  $res): \Psr\Http\Message\ResponseInterface {
    $this->get(StdoutLoggerInterface::class)->error(sprintf("%s\n%s:%s", $err->getMessage(), $err->getFile(), $err->getLine()));
    return $res->withBody(new SwooleStream(Json::encode(['message' => $err->getMessage()])))->withStatus(Status::INTERNAL_SERVER_ERROR);
});

$app->getContainer()->define('go', function (): RPC {
    $relay = Relay::create('unix:///tmp/php-go-scylladb.sock');
    defer(static fn () => $relay->close());
    return new RPC($relay);
});

$app->get('/', function (RequestInterface $req, ResponseInterface $res): \Psr\Http\Message\ResponseInterface {
    /** @var RPC $go */
    $go = $this->make('go');
    $todos = $go->call('ScyllaDB.Query', 'SELECT * FROM todos');
    return $res->json($todos);
});

$app->run();
