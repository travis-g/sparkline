mkbar
=====

mkbar is for printing configurable, colored, text-based progress bars to a console. Originally built for use with [update-motd][update-motd] and tmux, but should work great for anything that renders numeric output from a shell command, like [Übersicht][uebersicht] or [BitBar][bitbar].

Colored output is from [mitchellh/colorstring][colorstring], and environment variable parsing from [caarlos0/env][env].

```console
$ go get github.com/travis-g/utils/mkbar
```

## Usage

Either provide a 0-100 percentage value as `$1` or pipe a value in:

```console
$ mkbar 30
######--------------
$ pmset -g batt | grep -Eo "\d+%" | cut -d% -f1 | mkbar # my battery at time of writing
##############------
```

Configure it inline or with shell variables:

```console
$ export SIZE=15 START="[" END="]"
$ CHAR1="━" SEP="╋" CHAR2="━" mkbar 50
[━━━━━━━╋━━━━━━━]
```

Note that `SEP` is the first character of the _inactive_ portion of the bar.

If you're looking for conditional color formatting, try a wrapper script:

```console
$ cat mkbar.sh
#!/usr/bin/env bash
[[ "$1" -le "20" ]] && export START="[red]⚠️ "
mkbar $1

$ ./mkbar.sh 50
##########----------
$ ./mkbar.sh 20
⚠️ ####----------------
```

[bitbar]: https://getbitbar.com/
[colorstring]: https://github.com/mitchellh/colorstring
[env]: https://github.com/caarlos0/env
[uebersicht]: http://tracesof.net/uebersicht/
[update-motd]: http://manpages.ubuntu.com/manpages/trusty/man5/update-motd.5.html
