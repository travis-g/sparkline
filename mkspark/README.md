mkspark
=======

mkspark is a Go implementation of [spark][spark] for Unicode sparklines in your terminal.

```console
$ go get github.com/travis-g/utils/mkspark
```

This implementation works off of floats rather than integers, so decimals and percentage values (converted to decimals) should work fine.

## Usage

Pass `mkspark` arguments or pipe space- or comma-separated data into it:

```console
$ mkspark $(seq 0 20 | sort -R)
▇▆▅▃▃▂▄▅▆▄▁▁█▇▂▆▂▃▄▁▅
$ seq 0 20 | sort -R | mkspark
▁▃▁▃█▂▇▅▂▆▆▄▅▇▂▅▁▄▄▆▃
```

Here's a sparkline of the character count of each line of this README:

```console
$ while read -r line; do echo "$line" | wc -m; done <README.md | mkspark
▁▁▄▁▁▃▁▁▇▁▁▁▄▁▁▂▂▂▂▁▁▄▁▁▄▂▁▁▅▁▁▄▃▃▁▁▃▃
```

Here's the hourly temperatures forecasted for my area (via [DarkSky][darksky-api]):

```console
$ mkspark $(curl -s https://api.darksky.net/forecast/$API_KEY/38.907,-77.036 \
    | jq -r '.hourly.data|.[].temperature' | paste -sd" " -)
▁▁▁▁▁▁▁▁▁▁▁▂▂▂▃▄▄▅▅▅▄▄▃▃▃▃▃▃▃▃▃▂▂▂▂▃▄▅▆▇▇█▇▇▆▅▄▄▄
```

[darksky-api]: https://darksky.net/poweredby/
[spark]: https://github.com/holman/spark
