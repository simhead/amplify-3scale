// CHANGE_HERE - Change the module path below to reference packages correctly
module github.com/someorg/samples/apic_discovery_agent

go 1.13

require github.com/Axway/agent-sdk v1.0.20210617-0.20210701184357-df69af9d9be1

replace (
	github.com/Shopify/sarama => github.com/elastic/sarama v0.0.0-20191122160421-355d120d0970
	github.com/dop251/goja => github.com/andrewkroh/goja v0.0.0-20190128172624-dd2ac4456e20
	github.com/fsnotify/fsevents => github.com/fsnotify/fsevents v0.1.1
)
