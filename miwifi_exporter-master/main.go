package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/helloworlde/miwifi-exporter/collector"
	"github.com/helloworlde/miwifi-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	metricsPath      = flag.String("web.telemetry-path", "/metrics", "A path under which to expose metrics.")
	metricsNamespace = flag.String("metric.namespace", "miwifi", "Prometheus metrics namespace, as the prefix of metrics name")
)

func main() {
	log.Println("Bem-vindo ao cliente Prometheus para monitoramento de roteadores Xiaomi, projeto miwifi_exporter, autor: Huck. Sinta-se à vontade para enviar issues e Pull Requests.")
	log.Println("Inicializando o programa")
	flag.Parse()
	config.GetConfig()
	log.Println("Inicialização concluída")

	log.Println("Inicializando métricas de monitoramento")
	metrics := collector.NewMetrics(*metricsNamespace)
	registry := prometheus.NewRegistry()
	registry.MustRegister(metrics)
	log.Println("Métricas de monitoramento registradas com sucesso")

	log.Println("Iniciando o servidor, escutando na porta: " + strconv.Itoa(config.Configs.Port))
	http.Handle(*metricsPath, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`<html>
			<head><title>Um Exportador Prometheus</title></head>
			<body>
			<h1>Um Exportador Prometheus</h1>
			<p><a href='/metrics'>Métricas</a></p>
			</body>
			</html>`))
		if err != nil {
			log.Println("Erro ao executar o exportador", err)
			os.Exit(1)
		}
	})

	log.Printf("Localização das Métricas de Monitoramento: http://localhost:%d%s", config.Configs.Port, *metricsPath)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Configs.Port), nil))
}

