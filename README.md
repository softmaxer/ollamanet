# OllamaNet
A load balanced network of ollama instances

## Usage
```sh
git clone https://github.com/softmaxer/ollamanet.git

cd ollamanet

go build .

./ollamanet
```

### Command line arguments
```md
- **-config**: A JSON file containing the adresses of ollama instances started by docker. Default: `configuration.json` (See example)
- **-log**: File to log all the network outputs to. Default: `network_log`
```

## Upcoming features:
- The ability to create sequential pipeliens of Ollama instances, therefore combining models, and potentially running pipelines in parallel (if enough resources are available).
