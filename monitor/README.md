# Installation 

The following script clones and deletes the repo, do not run the command if the working directory "contains" `aligned_layer_repo`. The script is meant to be used to deploy a new monitor instance.

Step1:

```sh
export SLACK_URL=your_url
```

Step2:
```sh
curl https://raw.githubusercontent.com/yetanotherco/aligned_layer_tendermint/block_monitor/monitor/install_monitor.sh | bash
```
