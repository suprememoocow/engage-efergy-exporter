package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type reading struct {
	CID    string               `json:"cid"`
	SID    string               `json:"sid"`
	Values []map[string]float64 `json:"data"`
	Age    float64              `json:"age"`
}

func engageEfergyQuery(url string) ([]reading, error) {
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("user-agent", "engage-efergy-exporter")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var readings []reading
	err = json.Unmarshal(body, &readings)
	if err != nil {
		return nil, err
	}

	return readings, nil
}

type engageEfergyCollector struct {
	endpoint     string
	currentPower *prometheus.Desc
}

func (c *engageEfergyCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.currentPower
}

func (c *engageEfergyCollector) Collect(ch chan<- prometheus.Metric) {
	readings, err := engageEfergyQuery(c.endpoint)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(c.currentPower, err)
		return
	}

	for _, r := range readings {
		// This should only contain a single value
		// so break after the first value
	ReadingsLoop:
		for _, d := range r.Values {
			for _, v := range d {
				ch <- prometheus.MustNewConstMetric(
					c.currentPower,
					prometheus.GaugeValue,
					v,
					r.CID,
					r.SID,
				)

				break ReadingsLoop
			}
		}
	}

}

func main() {
	var (
		listenAddress = flag.String("web.listen-address", ":9236",
			"Address on which to expose metrics and web interface.")
		endpoint = flag.String("efergy.endpoint", "https://engage.efergy.com/proxy/getCurrentValuesSummary",
			"Endpoint to scrape for data")
		token = flag.String("efergy.token", "", "API token, obtained from the Engage portal")
	)
	flag.Parse()

	if *token == "" {
		flag.Usage()
		os.Exit(1)
	}

	endpointURL, error := url.Parse(*endpoint)
	if error != nil {
		log.Fatalf("Invalid endpoint URL %s", *endpoint)
		return
	}

	values, _ := url.ParseQuery(endpointURL.RawQuery)
	values.Set("token", *token)

	endpointURL.RawQuery = values.Encode()

	prometheus.MustRegister(&engageEfergyCollector{
		endpoint: endpointURL.String(),
		currentPower: prometheus.NewDesc(
			"engage_efergy_current_power_watts",
			"Current power reading, in watts",
			[]string{"cid", "sid"},
			nil,
		),
	})

	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(*listenAddress, nil)

	select {}
}
