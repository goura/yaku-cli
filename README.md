# yaku-cli

Yet another small tool to translate texts from your command line using DeepL API (and possibly other services in the future).

Yaku means "translation" in Japanese (or it could mean a yak of course).

## Prerequisites

This is not a scraping tool.
You need a proper DeepL API auth key to use this tool.
You can sign up and get it [here](https://www.deepl.com/pro#developer).
(I'm not involved in DeepL in any way)

## Installation

```bash
go install github.com/goura/yaku-cli/cmd/yaku@latest
```
`yaku` will be installed under your GOPATH/bin (`go env GOPATH` if you want the path).

## Usage
```bash
% export YAKU_DEEPL_AUTH_KEY="<YOUR_DEEPL_API_KEY>"

% echo "おはよう！" | yaku -s ja -t id
Pagi!
```

For source and target languages we internally use [golang.org/x/text/language](golang.org/x/text/language) so that we can take BCP 47 language tags such as `en-US` or just a language tag such as `en` or `ja`. We convert those to DeepL's language codes internally.


## License
See [LICENSE](LICENSE).
