#!/bin/bash

if [[ $# -lt 1 || (($1 == "setup" || $1 == "run") && $# -lt 2) || ($1 != "stop" && $1 != "logs" && $1 == "setup" && $1 == "run") ]]
then
	echo "Usage:"
	echo -e "\t$0 <setup|run> <node_name>"
	echo -e "\t$0 <stop|logs>"
	exit 1
fi

CMD="$1"
NODE_NAME="$2"

if [ "$CMD" == "setup" ]
then
	mkdir -p prod-sim
	mkdir prod-sim/$NODE_NAME 2>/dev/null
	if [ $? != 0 ]
	then
		echo -n "A validator with name $NODE_NAME already exists. Do you want to override it? [y/N] "
		read

		if [ "$REPLY" != "y" ]
		then
			echo "Aborting..."
			exit 2
		fi

		rm -r prod-sim/$NODE_NAME
		mkdir prod-sim/$NODE_NAME
	fi

	NODE_NAME=$NODE_NAME docker compose -f docker/compose/validator.docker-compose.yml up node-setup
fi

if [ "$CMD" == "run" ]
then
	ls prod-sim/$NODE_NAME >/dev/null 2>&1
	if [ $? != 0 ]
	then
		echo "No validator config was found with that name. Try running \`$0 setup $NODE_NAME\` first"
		exit 3
	fi

	NODE_NAME=$NODE_NAME docker compose -f docker/compose/validator.docker-compose.yml up -d validator-runner

	ls prod-sim/$NODE_NAME/config/validator.json >/dev/null 2>&1
	if [ $? != 0 ]  # The validator is not initialized yet
	then
		while $(curl -s localhost:26657/status | jq .result.sync_info.catching_up); do
			printf '.'
			sleep 1
		done
		NODE_NAME=$NODE_NAME docker compose -f docker/compose/validator.docker-compose.yml up validator-setup
	fi
fi

if [ "$CMD" == "stop" ]
then
	NODE_NAME=$NODE_NAME docker compose -f docker/compose/validator.docker-compose.yml stop validator-runner
fi

if [ "$CMD" == "logs" ]
then
	NODE_NAME=$NODE_NAME docker compose -f docker/compose/validator.docker-compose.yml logs -f
fi
