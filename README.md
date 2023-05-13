# yaku-cli

Yet another small tool to translate texts from your command line using DeepL API (and other services such as OpenAI Chat Completion)

Yaku means "translation" in Japanese (or it could mean a yak of course).

## Prerequisites

This is not a scraping tool.
You need a proper DeepL API auth key or an OpenAI API key to use this tool.

You can sign up and get a DeepL API auth key [here](https://www.deepl.com/pro#developer).
(I'm not involved in DeepL in any way)

DISCLAIMER: YOU MIGHT BE BILLED THAN YOU ARE EXPECTED. READ THE LICENSE AND UNDERSTAND THAT THISE SOFTWARE COMES WITHOUT WARRANTY OF ANY KIND BEFORE YOU USE IT.

## Installation

```bash
go install github.com/goura/yaku-cli/cmd/yaku@latest
```
`yaku` will be installed under your GOPATH/bin (`go env GOPATH` if you want the path).

## Usage

DeepL
```bash
% export YAKU_DEEPL_AUTH_KEY="<YOUR_DEEPL_API_KEY>"

% echo "おはよう！" | yaku -s ja -t id
Pagi!
```

OpenAI (much experimental)
```bash
% export YAKU_OPENAI_API_KEY="<YOUR_OPENAI_API_KEY>"

% echo "Saya tidak mau bekerja, tapi saya mau uang" | go run cmd/yaku/main.go -e openai -s id -t ja
私は働きたくないですが、お金が欲しいです。
```

For source and target languages we internally use [golang.org/x/text/language](golang.org/x/text/language) so that we can take BCP 47 language tags such as `en-US` or just a language tag such as `en` or `ja`. We convert those to DeepL's language codes internally.

For OpenAI's Chat Completion, we don't know which language they are capable of, so we just pass the language tag to OpenAI's Chat Completion API (`gpt-3.5-turbo`).

## License & disclaimer
See [LICENSE](LICENSE)
