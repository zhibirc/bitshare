# bitshare

<h3 align="center">Architecture Overview</h3>
<p align="center">
    <img src="/assets/images/schema.jpg">
</p>

## Purpose/motivation



## Usage

### Development

Run Redis server on port `6379` using Docker. This also exposes _RedisInsight_ on port `8001` for data visualization.
For more information please check [Run Redis Stack on Docker](https://redis.io/docs/stack/get-started/install/docker/).

```shell
docker run -d --name redis-stack -p 6379:6379 -p 8001:8001 redis/redis-stack:6.2.6-v7
```

