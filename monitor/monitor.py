import os
import requests
import time
from slack_sdk import WebhookClient

SLACK_URL = os.environ["SLACK_URL"]

urls = ["http://91.107.239.79:26657/",
        "http://116.203.81.174:26657/",
        "http://88.99.174.203:26657/",
        "http://128.140.3.188:26657/"]

NUMBER_OF_NODES = 4

#url = "http://0.0.0.0:26657/"
#url = "http:/100.76.93.84:26657/"

def get_block_of(url):
    while True:
            try: 
                height = requests.get(url+"abci_info?").json()["result"]["response"]["last_block_height"]
                timestamp =  requests.get(url+"block?", params={"height": height}).json()["result"]["block"]["header"]["time"]
                return (height,timestamp)
            except:
                print("Waiting to check again...")
                time.sleep(5)
                continue
        

def send_alert(node_url, height, timestamp):
    webhook = WebhookClient(SLACK_URL)
    webhook.send(text="Node with ip:" + node_url + "is not advancing it's state. The last block height is "+ height + " at "+ timestamp)

def send_blockchain_halted_alert():
    webhook = WebhookClient(SLACK_URL)
    webhook.send(text="The chain is halted. There aren't enough nodes validating blocks for consensus")

if __name__ == "__main__":
    last_height = [""] * 4
    current_height = [""] * 4
    alive = [True for i in range(NUMBER_OF_NODES)]
    for i in range(NUMBER_OF_NODES):
        print(i)
        last_height[i], _ = get_block_of(urls[i])
    
    while True:
        time.sleep(60)
        amount_of_failures = 0
        for i in range(NUMBER_OF_NODES):
            current_height[i], last_timestamp = get_block_of(urls[i])
            if current_height[i]==last_height[i] and alive:
                amount_of_failures = amount_of_failures + 1
                send_alert(urls[i],current_height[i],last_timestamp[i])
                alive[i] = False
            else:
                alive[i] = True

            if amount_of_failures > 1: 
                send_blockchain_halted_alert()

            last_height[i] = current_height[i]
