module apple

go 1.23

toolchain go1.23.3

require (
	github.com/BurntSushi/toml v1.4.0
	github.com/Unknwon/log v0.0.0-20200308114134-929b1006e34a
	github.com/allegro/bigcache v1.2.1
	github.com/arl/statsviz v0.6.0
	github.com/armon/go-socks5 v0.0.0-20160902184237-e75332964ef5
	github.com/bradfitz/gomemcache v0.0.0-20230905024940-24af94b03874
	github.com/bsm/redislock v0.9.4
	github.com/btcsuite/btcd/btcec/v2 v2.3.4
	github.com/btcsuite/btcd/chaincfg/chainhash v1.1.0
	github.com/cespare/xxhash v1.1.0
	github.com/cockroachdb/errors v1.11.3
	github.com/coocood/freecache v1.2.4
	github.com/dgraph-io/ristretto v0.2.0
	github.com/dgrijalva/lfu-go v0.0.0-20141010002404-f174e76c5138
	github.com/didip/tollbooth v4.0.2+incompatible
	github.com/felixge/tcpkeepalive v0.0.0-20220224101934-f56176a53a1b
	github.com/gin-contrib/cache v1.3.0
	github.com/gin-contrib/sessions v1.0.1
	github.com/gin-gonic/gin v1.10.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-redis/redis_rate v6.5.0+incompatible
	github.com/go-session/session v2.4.0+incompatible
	github.com/golang/glog v1.2.3
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da
	github.com/golang/protobuf v1.5.4
	github.com/google/btree v1.1.3
	github.com/gorilla/csrf v1.7.2
	github.com/gorilla/mux v1.8.1
	github.com/gorilla/sessions v1.4.0
	github.com/gorilla/websocket v1.5.3
	github.com/harveywangdao/ants v0.0.0-20211009021109-9d057542a967
	github.com/harveywangdao/earth v0.0.0-20240202030638-d7b2d65645d8
	github.com/hashicorp/go-hclog v1.6.3
	github.com/hashicorp/golang-lru v1.0.2
	github.com/hashicorp/raft v1.7.1
	github.com/hashicorp/raft-boltdb v0.0.0-20231211162105-6c830fa4535e
	github.com/ipfs/go-ipfs-api v0.7.0
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible
	github.com/juju/ratelimit v1.0.2
	github.com/justinas/alice v1.2.0
	github.com/justinas/nosurf v1.1.1
	github.com/minio/sha256-simd v1.0.1
	github.com/mitchellh/hashstructure v1.1.0
	github.com/mitchellh/mapstructure v1.5.0
	github.com/mr-tron/base58 v1.2.0
	github.com/oklog/run v1.1.0
	github.com/panjf2000/ants v1.2.1
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/petar/GoLLRB v0.0.0-20210522233825-ae3b015fd3e9
	github.com/pkg/errors v0.9.1
	github.com/quic-go/quic-go v0.48.1
	github.com/rs/zerolog v1.33.0
	github.com/sasha-s/go-deadlock v0.3.5
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.9.3
	github.com/smallfish/ftp v0.0.0-20160801035311-6d094f003ac5
	github.com/streadway/amqp v1.1.0
	github.com/tealeg/xlsx v1.0.5
	github.com/throttled/throttled v1.0.0
	github.com/throttled/throttled/v2 v2.12.0
	github.com/timtadh/fs2 v0.1.0
	github.com/ulule/limiter v2.2.2+incompatible
	github.com/xtaci/kcptun v0.0.0-20241114025705-1989e91e90f2
	go.etcd.io/etcd/client/v3 v3.5.17
	go.uber.org/ratelimit v0.3.1
	go.uber.org/zap v1.27.0
	golang.org/x/crypto v0.29.0
	golang.org/x/net v0.31.0
	google.golang.org/grpc v1.68.0
	gopkg.in/gemnasium/logrus-airbrake-hook.v4 v4.1.0
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
	gopkg.in/yaml.v2 v2.4.0
	h12.io/socks v1.0.3
)

require (
	github.com/PuerkitoBio/boom v0.0.0-20140219125548-fecdef1c97ca // indirect
	github.com/airbrake/gobrake/v4 v4.0.3 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/benbjohnson/clock v1.3.0 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/bytedance/sonic v1.11.6 // indirect
	github.com/bytedance/sonic/loader v0.1.1 // indirect
	github.com/caio/go-tdigest v2.3.0+incompatible // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/redact v1.1.5 // indirect
	github.com/coreos/go-iptables v0.7.0 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/crackcomm/go-gitignore v0.0.0-20170627025303-887ab5e44cc3 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.1.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/getsentry/sentry-go v0.27.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.20.0 // indirect
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible // indirect
	github.com/google/gopacket v1.1.19 // indirect
	github.com/google/pprof v0.0.0-20221203041831-ce31453925ec // indirect
	github.com/gorilla/context v1.1.2 // indirect
	github.com/gorilla/securecookie v1.1.2 // indirect
	github.com/hashicorp/go-immutable-radix v1.0.0 // indirect
	github.com/hashicorp/go-msgpack v0.5.5 // indirect
	github.com/hashicorp/go-msgpack/v2 v2.1.2 // indirect
	github.com/ipfs/boxo v0.12.0 // indirect
	github.com/ipfs/go-cid v0.4.1 // indirect
	github.com/jonboulle/clockwork v0.1.1-0.20190114141812-62fb9bc030d1 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.8 // indirect
	github.com/klauspost/reedsolomon v1.12.4 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/libp2p/go-buffer-pool v0.1.0 // indirect
	github.com/libp2p/go-flow-metrics v0.1.0 // indirect
	github.com/libp2p/go-libp2p v0.26.3 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/memcachier/mc/v3 v3.0.3 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/multiformats/go-base32 v0.1.0 // indirect
	github.com/multiformats/go-base36 v0.2.0 // indirect
	github.com/multiformats/go-multiaddr v0.8.0 // indirect
	github.com/multiformats/go-multibase v0.2.0 // indirect
	github.com/multiformats/go-multicodec v0.9.0 // indirect
	github.com/multiformats/go-multihash v0.2.3 // indirect
	github.com/multiformats/go-multistream v0.4.1 // indirect
	github.com/multiformats/go-varint v0.0.7 // indirect
	github.com/olekukonko/ts v0.0.0-20171002115256-78ecb04241c0 // indirect
	github.com/onsi/ginkgo/v2 v2.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/petermattis/goid v0.0.0-20240813172612-4fcff4a6cae7 // indirect
	github.com/quic-go/qpack v0.5.1 // indirect
	github.com/rakyll/pb v0.0.0-20160123035540-8d46b8b097ef // indirect
	github.com/redis/go-redis/v9 v9.0.5 // indirect
	github.com/robfig/go-cache v0.0.0-20130306151617-9fc39e0dbf62 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/rs/xid v1.5.0 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/templexxx/cpu v0.1.1 // indirect
	github.com/templexxx/xorsimd v0.4.3 // indirect
	github.com/tidwall/btree v1.4.2 // indirect
	github.com/tidwall/buntdb v1.3.2 // indirect
	github.com/tidwall/gjson v1.14.3 // indirect
	github.com/tidwall/grect v0.1.4 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/tidwall/rtred v0.1.2 // indirect
	github.com/tidwall/tinyqueue v0.1.1 // indirect
	github.com/tjfoc/gmsm v1.4.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	github.com/xtaci/kcp-go/v5 v5.6.18 // indirect
	github.com/xtaci/qpp v1.1.17 // indirect
	github.com/xtaci/tcpraw v1.2.31 // indirect
	go.etcd.io/etcd/api/v3 v3.5.17 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.17 // indirect
	go.uber.org/mock v0.4.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/arch v0.8.0 // indirect
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/sys v0.27.0 // indirect
	golang.org/x/text v0.20.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.21.1-0.20240508182429-e35e4ccd0d2d // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240903143218-8af14fe29dc1 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240903143218-8af14fe29dc1 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	lukechampine.com/blake3 v1.1.7 // indirect
)
