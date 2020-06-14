# Engage by Efergy/Energyhive Prometheus Exporter

A Prometheus/OpenMetrics exporter for the Engage by Efergy and EnergyHive Energy Meter

## Usage

```console
$ engage-efergy-exporter
Usage of ./engage-efergy-exporter:
  -efergy.endpoint string
    	Endpoint to scrape for data (default "https://engage.efergy.com/proxy/getCurrentValuesSummary")
  -efergy.token string
    	API token, obtained from the Engage portal
  -web.listen-address string
    	Address on which to expose metrics and web interface. (default ":9236")
```

Obtain a token from the Engage Portal at https://engage.efergy.com/, then run the exporter, passing the token in through the `-efergy.token` argument

```console
$ engage-efergy-exporter -efergy.token token_goes_here
...
```


## Output

A `engage_efergy_current_power_watts` metric will appear for each energy monitor in your installation.

```console
# HELP engage_efergy_current_power_watts Current power reading, in watts
# TYPE engage_efergy_current_power_watts gauge
engage_efergy_current_power_watts{cid="PWER",sid="761019"} 2098
engage_efergy_current_power_watts{cid="PWER_SUB",sid="761068"} 369
engage_efergy_current_power_watts{cid="PWER_SUB",sid="766335"} 758
```

### Licence

Copyright 2020, Andrew Newdigate

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
