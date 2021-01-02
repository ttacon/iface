package statsd

type ClientInterface interface {
	Close()
	Dec(stat string, value int64, rate float32) error
	Gauge(stat string, value int64, rate float32) error
	GaugeDelta(stat string, value int64, rate float32) error
	Inc(stat string, value int64, rate float32) error
	Raw(stat string, value string, rate float32) error
	SetPrefix(prefix string)
	Timing(stat string, delta int64, rate float32) error
}

type Client struct {
	ClientInterface
}
