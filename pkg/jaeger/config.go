package jaeger

type Config struct {
	Host        string `env:"JAEGER_HOST"`
	ServiceName string `env:"JAEGER_SERVICE_NAME"`
	LogSpans    bool   `env:"JAEGER_LOG_SPANS"`
}
