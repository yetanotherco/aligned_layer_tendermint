#!/bin/bash

if [[ $# -lt 1 || (($1 == "setup" || $1 == "run" || $1 == "run-validator") && $# -lt 2) || ($1 != "stop" && $1 != "logs" && $1 != "setup" && $1 != "run" && $1 != "down" && $1 != "run-validator") ]]
then
	echo "Usage:"
	echo -e "\t$0 <setup|run|run-validator> <node_name>"
	echo -e "\t$0 <stop|logs>"
	exit 1
fi

NODES_PATH="../prod-sim"
CMD="$1"
NODE_NAME="$2"

if [ "$CMD" == "setup" ]
then
	mkdir -p $NODES_PATH
	mkdir $NODES_PATH/$NODE_NAME 2>/dev/null
	if [ $? != 0 ]
	then
		echo -n "A node with name $NODE_NAME already exists. Do you want to override it? [y/N] "
		read

		if [ "$REPLY" != "y" ]
		then
			echo "Aborting..."
			exit 2
		fi

		rm -r $NODES_PATH/$NODE_NAME
		mkdir $NODES_PATH/$NODE_NAME
	fi

	NODE_NAME=$NODE_NAME docker compose -f compose/validator.docker-compose.yml up node-setup
fi

if [ "$CMD" == "run" ] || [ "$CMD" == "run-validator" ]
then
	ls $NODES_PATH/$NODE_NAME >/dev/null 2>&1
	if [ $? != 0 ]
	then
		echo "No node config was found with that name. Try running \`$0 setup $NODE_NAME\` first"
		exit 3
	fi

	NODE_NAME=$NODE_NAME docker compose -f compose/validator.docker-compose.yml up -d validator-runner

	if [ "$CMD" == "run" ]  # If it's not a validator, not extra actions needed
	then
		exit 0
	fi

	ls $NODES_PATH/$NODE_NAME/config/validator.json >/dev/null 2>&1
	if [ $? != 0 ]  # The validator is not initialized yet
	then
		echo -n "Waiting the node to be fully synced "
		while $(curl -s localhost:26657/status | jq .result.sync_info.catching_up); do
			echo -n '.'
			sleep 1
		done
		NODE_NAME=$NODE_NAME docker compose -f compose/validator.docker-compose.yml up validator-setup
	fi
fi

if [ "$CMD" == "stop" ]
then
	NODE_NAME=$NODE_NAME docker compose -f compose/validator.docker-compose.yml stop validator-runner
fi

if [ "$CMD" == "down" ]
then
	NODE_NAME=$NODE_NAME docker compose -f compose/validator.docker-compose.yml down
fi


if [ "$CMD" == "logs" ]
then
	NODE_NAME=$NODE_NAME docker compose -f compose/validator.docker-compose.yml logs -f
fi
