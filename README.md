# sparkline

A Go implementation of [spark][spark] for Unicode sparklines in your terminal.

```console
$ go get github.com/travis-g/sparkline
```

This implementation works off of floats rather than integers, so percentage values and decimals should work fine.

## Usage

Pass `sparkline` arguments or pipe space- or comma-separated data into it:

```console
$ sparkline $(seq 0 20 | sort -R)
▇▆▅▃▃▂▄▅▆▄▁▁█▇▂▆▂▃▄▁▅
$ seq 0 20 | sort -R | sparkline
▁▃▁▃█▂▇▅▂▆▆▄▅▇▂▅▁▄▄▆▃
```

Here's a sparkline of the character count of each line of this README:

```console
$ while read -r line; do echo "$line" | wc -m; done <README.md | sparkline
▁▁▅▁▁▂▁▁▆▁▁▁▄▁▁▂▂▂▂▁▁▄▁▁▄▂▁▁▅▁▁█▃▁▁▃▃
```

Here's the hourly temperatures forecasted for my area (via [DarkSky][darksky-api]):

```console
$ sparkline $(curl -s https://api.darksky.net/forecast/$API_KEY/38.907,-77.036 | jq -r '.hourly.data|.[].temperature' | paste -sd" " -)
▁▁▁▁▁▁▁▁▁▁▁▂▂▂▃▄▄▅▅▅▄▄▃▃▃▃▃▃▃▃▃▂▂▂▂▃▄▅▆▇▇█▇▇▆▅▄▄▄
```

[darksky-api]: https://darksky.net/poweredby/
[spark]: https://github.com/holman/spark
