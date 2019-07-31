## Dummy Envoy ALS Sink

TL;DR

```
$ make example-up
....
sink_1   | 2019/07/31 06:57:58 Started stream
sink_1   | 2019/07/31 06:57:58 Received value
sink_1   | 2019/07/31 06:57:58 {"identifier":{"node":{"cluster":"envoy-proxy","buildVersion":"e95ef6bc43daeda16451ad4ef20979d8e07a5299/1.10.0/Clean/RELEASE/BoringSSL"}},"envoyMetrics":[{"name":"listener_manager.listener_create_success","metric":[{"counter":{"value":1},"timestampMs":"1564556278869"}]},{"name":"listener_manager.listener_added","metric":[{"counter":{"value":1},"timestampMs":"1564556278869"}]},{"name":"runtime.deprecated_feature_use","metric":[{"counter":{"value":18},"timestampMs":"1564556278869"}]},{"name":"cluster_manager.cluster_added","metric":[{"counter":{"value":2},"timestampMs":"1564556278869"}]},{"name":"cluster.service_stats.update_attempt","metric":[{"counter":{"value":2},"timestampMs":"1564556278869"}]},{"name":"cluster.service_stats.membership_change","metric":[{"counter":{"value":1},"timestampMs":"1564556278869"}]},{"name":"cluster.service_stats.update_success","metric":[{"counter":{"value":1},"timestampMs":"1564556278869"}]},{"name":"cluster.service_google.update_attempt","metric":[{"counter":{"value":1},"timestampMs":"1564556278869"}]},{"name":"cluster.service_google.update_success","metric":[{"counter":{"value":1},"timestampMs":"1564556278869"}]},{"name":"cluster.service_google.membership_change","metric":[{"counter":{"value":1},"timestampMs":"1564556278869"}]},{"name":"listener_manager.total_listeners_active","type":"GAUGE","metric":[{"gauge":{"value":1},"timestampMs":"1564556278869"}]},{"name":"server.parent_connections","type":"GAUGE","metric":[{"gauge":{},"timestampMs":"1564556278869"}]},{"name":"cluster_manager.warming_clusters","type":"GAUGE","metric":[{"gauge":{},"timestampMs":"1564556278869"}]},{"name":"server.total_connections","type":"GAUGE","metric":[{"gauge":{},"timestampMs":"1564556278869"}]},{"name":"server.concurrency","type":"GAUGE","metric":[{"gauge":{"value":1},"timestampMs":"1564556278869"}]},{"name":"server.memory_allocated","type":"GAUGE","metric":[{"gauge":{"value":2003920},"timestampMs":"1564556278869"}]},{"name":"server.hot_restart_epoch","type":"GAUGE","metric":[{"gauge":{},"timestampMs":"1564556278869"}]},{"name":"server.uptime","type":"GAUGE","metric":[{"gauge":{"value":5},"timestampMs":"1564556278869"}]},{"name":"server.memory_heap_size","type":"GAUGE","metric":[{"gauge":{"value":3145728},"timestampMs":"1564556278869"}]},{"name":"server.days_until_first_cert_expiring","type":"GAUGE","metric":[{"gauge":{"value":2147483647},"timestampMs":"1564556278869"}]},{"name":"listener_manager.total_listeners_warming","type":"GAUGE","metric":[{"gauge":{},"timestampMs":"1564556278869"}]},{"name":"runtime.num_keys","type":"GAUGE","metric":[{"gauge":{},"timestampMs":"1564556278869"}]},{"name":"cluster_manager.active_clusters","type":"GAUGE","metric":[{"gauge":{"value":2},"timestampMs":"1564556278869"}]},{"name":"server.version","type":"GAUGE","metric":[{"gauge":{"value":15294198},"timestampMs":"1564556278869"}]},{"name":"server.live","type":"GAUGE","metric":[{"gauge":{"value":1},"timestampMs":"1564556278869"}]},{"name":"cluster.service_stats.membership_degraded","type":"GAUGE","metric":[{"gauge":{},"timestampMs":"1564556278869"}]},{"name":"cluster.service_stats.membership_total","type":"GAUGE","metric":[{"gauge":{"value":1},"timestampMs":"1564556278869"}]},{"name":"cluster.service_stats.membership_healthy","type":"GAUGE","metric":[{"gauge":{"value":1},"timestampMs":"1564556278869"}]},{"name":"cluster.service_google.membership_total","type":"GAUGE","metric":[{"gauge":{"value":1},"timestampMs":"1564556278869"}]},{"name":"cluster.service_google.membership_healthy","type":"GAUGE","metric":[{"gauge":{"value":1},"timestampMs":"1564556278869"}]},{"name":"cluster.service_google.membership_degraded","type":"GAUGE","metric":[{"gauge":{},"timestampMs":"1564556278869"}]}]}
sink_1   | 2019/07/31 06:58:03 Received value
```

As you can see from above, we received `envoyMetrics` and `identifier` from Envoy through gRPC connection.

From another tab, you can call:

```
// Use $(docker-machine ip) or localhost for non docker-machine platforms.
$ curl $(docker-machine ip):10000
```

And to tear-down:

```
$ make example-down
```

### License

MIT
