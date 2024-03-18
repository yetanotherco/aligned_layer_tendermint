import os
import requests
import time
from slack_sdk import WebhookClient

SLACK_URL = os.environ["SLACK_URL"]

url = "http://91.107.239.79:26657/"
#url = "http:/100.76.93.84:26657/"

def get_block():
    while True:
        try: 
            height = requests.get(url+"abci_info?").json()["result"]["response"]["last_block_height"]
            timestamp =  requests.get(url+"block?", params={"height": height}).json()["result"]["block"]["header"]["time"]
            return (height,timestamp)
        except:
            print("Waiting to check again...")
            time.sleep(30)
            continue
        

def send_alert(height, timestamp):
    webhook = WebhookClient(SLACK_URL)
    webhook.send(text="The chain is not producing blocks currently. The last block height is "+ height + " at "+ timestamp)

if __name__ == "__main__":
    alive = True
    last_height, _ = get_block()
    while True:
        time.sleep(120)
        current_height, last_timestamp = get_block()

        if current_height==last_height and alive:
            send_alert(current_height,last_timestamp)
            alive = False
        else:
            alive = True
        last_height = current_height
        


        
