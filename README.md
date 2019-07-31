## Dummy Envoy ALS Sink

TL;DR

```
$ make example-up
make -C example up
docker-compose up --build
Creating network "example_default" with the default driver
Building sink
Step 1/3 : FROM gcr.io/tetratelabs/tetrate-base:v0.1
 ---> f453f095a19f
Step 2/3 : ADD build/sink /usr/local/bin/sink
 ---> Using cache
 ---> 93d99cb3a0b0
Step 3/3 : ENTRYPOINT ["/usr/local/bin/sink"]
 ---> Using cache
 ---> 2679d8223bed
Successfully built 2679d8223bed
Successfully tagged example_sink:latest
Creating example_sink_1  ... done
Creating example_envoy_1 ... done
Attaching to example_sink_1, example_envoy_1
envoy_1  | [2019-07-31 06:57:53.811][1][info][main] [source/server/server.cc:205] initializing epoch 0 (hot restart version=10.200.16384.127.options=capacity=16384, num_slots=8209 hash=228984379728933363 size=2654312)
envoy_1  | [2019-07-31 06:57:53.839][1][info][main] [source/server/server.cc:207] statically linked extensions:
envoy_1  | [2019-07-31 06:57:53.839][1][info][main] [source/server/server.cc:209]   access_loggers: envoy.file_access_log,envoy.http_grpc_access_log
envoy_1  | [2019-07-31 06:57:53.839][1][info][main] [source/server/server.cc:212]   filters.http: envoy.buffer,envoy.cors,envoy.ext_authz,envoy.fault,envoy.filters.http.grpc_http1_reverse_bridge,envoy.filters.http.header_to_metadata,envoy.filters.http.jwt_authn,envoy.filters.http.rbac,envoy.filters.http.tap,envoy.grpc_http1_bridge,envoy.grpc_json_transcoder,envoy.grpc_web,envoy.gzip,envoy.health_check,envoy.http_dynamo_filter,envoy.ip_tagging,envoy.lua,envoy.rate_limit,envoy.router,envoy.squash
envoy_1  | [2019-07-31 06:57:53.839][1][info][main] [source/server/server.cc:215]   filters.listener: envoy.listener.original_dst,envoy.listener.original_src,envoy.listener.proxy_protocol,envoy.listener.tls_inspector
envoy_1  | [2019-07-31 06:57:53.839][1][info][main] [source/server/server.cc:218]   filters.network: envoy.client_ssl_auth,envoy.echo,envoy.ext_authz,envoy.filters.network.dubbo_proxy,envoy.filters.network.mysql_proxy,envoy.filters.network.rbac,envoy.filters.network.sni_cluster,envoy.filters.network.thrift_proxy,envoy.filters.network.zookeeper_proxy,envoy.http_connection_manager,envoy.mongo_proxy,envoy.ratelimit,envoy.redis_proxy,envoy.tcp_proxy
envoy_1  | [2019-07-31 06:57:53.839][1][info][main] [source/server/server.cc:220]   stat_sinks: envoy.dog_statsd,envoy.metrics_service,envoy.stat_sinks.hystrix,envoy.statsd
envoy_1  | [2019-07-31 06:57:53.839][1][info][main] [source/server/server.cc:222]   tracers: envoy.dynamic.ot,envoy.lightstep,envoy.tracers.datadog,envoy.zipkin
envoy_1  | [2019-07-31 06:57:53.839][1][info][main] [source/server/server.cc:225]   transport_sockets.downstream: envoy.transport_sockets.alts,envoy.transport_sockets.tap,raw_buffer,tls
envoy_1  | [2019-07-31 06:57:53.839][1][info][main] [source/server/server.cc:228]   transport_sockets.upstream: envoy.transport_sockets.alts,envoy.transport_sockets.tap,raw_buffer,tls
envoy_1  | [2019-07-31 06:57:53.839][1][info][main] [source/server/server.cc:234] buffer implementation: old (libevent)
envoy_1  | [2019-07-31 06:57:53.848][1][info][main] [source/server/server.cc:281] admin address: 127.0.0.1:9901
envoy_1  | [2019-07-31 06:57:53.850][1][info][config] [source/server/configuration_impl.cc:50] loading 0 static secret(s)
envoy_1  | [2019-07-31 06:57:53.851][1][info][config] [source/server/configuration_impl.cc:56] loading 2 cluster(s)
envoy_1  | [2019-07-31 06:57:53.855][1][info][config] [source/server/configuration_impl.cc:60] loading 1 listener(s)
envoy_1  | [2019-07-31 06:57:53.857][1][info][config] [source/server/configuration_impl.cc:85] loading tracing configuration
envoy_1  | [2019-07-31 06:57:53.858][1][info][config] [source/server/configuration_impl.cc:105] loading stats sink configuration
envoy_1  | [2019-07-31 06:57:53.859][1][info][main] [source/server/server.cc:478] starting main dispatch loop
envoy_1  | [2019-07-31 06:57:53.882][1][info][upstream] [source/common/upstream/cluster_manager_impl.cc:137] cm init: all clusters initialized
envoy_1  | [2019-07-31 06:57:53.883][1][info][main] [source/server/server.cc:462] all clusters initialized. initializing init manager
sink_1   | 2019/07/31 06:57:53 Listening on tcp://localhost:9001
sink_1   | 2019/07/31 06:57:58 Started stream
sink_1   | 2019/07/31 06:57:58 Received value
sink_1   | 2019/07/31 06:57:58 {"identifier":{"node":{"cluster":"envoy-proxy","buildVersion":"e95ef6bc43daeda16451ad4ef20979d8e07a5299/1.10.0/Clean/RELEASE/BoringSSL"}},"envoyMetrics":[{"name":"listener_manager.listener_create_success","metric":[{"counter":{"value":1},"timestampMs":"1564556278869"}]},{"name":"listener_manager.listener_added","metric":[{"counter":{"value":1},"timestampMs":"1564556278869"}]},{"name":"runtime.deprecated_feature_use","metric":[{"counter":{"value":18},"timestampMs":"1564556278869"}]},{"name":"cluster_manager.cluster_added","metric":[{"counter":{"value":2},"timestampMs":"1564556278869"}]},{"name":"cluster.service_stats.update_attempt","metric":[{"counter":{"value":2},"timestampMs":"1564556278869"}]},{"name":"cluster.service_stats.membership_change","metric":[{"counter":{"value":1},"timestampMs":"1564556278869"}]},{"name":"cluster.service_stats.update_success","metric":[{"counter":{"value":1},"timestampMs":"1564556278869"}]},{"name":"cluster.service_google.update_attempt","metric":[{"counter":{"value":1},"timestampMs":"1564556278869"}]},{"name":"cluster.service_google.update_success","metric":[{"counter":{"value":1},"timestampMs":"1564556278869"}]},{"name":"cluster.service_google.membership_change","metric":[{"counter":{"value":1},"timestampMs":"1564556278869"}]},{"name":"listener_manager.total_listeners_active","type":"GAUGE","metric":[{"gauge":{"value":1},"timestampMs":"1564556278869"}]},{"name":"server.parent_connections","type":"GAUGE","metric":[{"gauge":{},"timestampMs":"1564556278869"}]},{"name":"cluster_manager.warming_clusters","type":"GAUGE","metric":[{"gauge":{},"timestampMs":"1564556278869"}]},{"name":"server.total_connections","type":"GAUGE","metric":[{"gauge":{},"timestampMs":"1564556278869"}]},{"name":"server.concurrency","type":"GAUGE","metric":[{"gauge":{"value":1},"timestampMs":"1564556278869"}]},{"name":"server.memory_allocated","type":"GAUGE","metric":[{"gauge":{"value":2003920},"timestampMs":"1564556278869"}]},{"name":"server.hot_restart_epoch","type":"GAUGE","metric":[{"gauge":{},"timestampMs":"1564556278869"}]},{"name":"server.uptime","type":"GAUGE","metric":[{"gauge":{"value":5},"timestampMs":"1564556278869"}]},{"name":"server.memory_heap_size","type":"GAUGE","metric":[{"gauge":{"value":3145728},"timestampMs":"1564556278869"}]},{"name":"server.days_until_first_cert_expiring","type":"GAUGE","metric":[{"gauge":{"value":2147483647},"timestampMs":"1564556278869"}]},{"name":"listener_manager.total_listeners_warming","type":"GAUGE","metric":[{"gauge":{},"timestampMs":"1564556278869"}]},{"name":"runtime.num_keys","type":"GAUGE","metric":[{"gauge":{},"timestampMs":"1564556278869"}]},{"name":"cluster_manager.active_clusters","type":"GAUGE","metric":[{"gauge":{"value":2},"timestampMs":"1564556278869"}]},{"name":"server.version","type":"GAUGE","metric":[{"gauge":{"value":15294198},"timestampMs":"1564556278869"}]},{"name":"server.live","type":"GAUGE","metric":[{"gauge":{"value":1},"timestampMs":"1564556278869"}]},{"name":"cluster.service_stats.membership_degraded","type":"GAUGE","metric":[{"gauge":{},"timestampMs":"1564556278869"}]},{"name":"cluster.service_stats.membership_total","type":"GAUGE","metric":[{"gauge":{"value":1},"timestampMs":"1564556278869"}]},{"name":"cluster.service_stats.membership_healthy","type":"GAUGE","metric":[{"gauge":{"value":1},"timestampMs":"1564556278869"}]},{"name":"cluster.service_google.membership_total","type":"GAUGE","metric":[{"gauge":{"value":1},"timestampMs":"1564556278869"}]},{"name":"cluster.service_google.membership_healthy","type":"GAUGE","metric":[{"gauge":{"value":1},"timestampMs":"1564556278869"}]},{"name":"cluster.service_google.membership_degraded","type":"GAUGE","metric":[{"gauge":{},"timestampMs":"1564556278869"}]}]}
sink_1   | 2019/07/31 06:58:03 Received value
```

From another tab, you can call:

```
// Use $(docker-machine ip) or localhost for non docker-machine platforms.
$ curl $(docker-machine ip):10000
```

### License

MIT
