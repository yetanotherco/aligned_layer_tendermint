import express from 'express';

import { DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { SigningStargateClient } from "@cosmjs/stargate";

import conf from './config/config.js'
import { FrequencyChecker } from './checker.js';

import { Mutex, withTimeout } from 'async-mutex';

// load config
console.log("loaded config: ", conf)

const mutex = new Mutex();

const app = express()

app.set("view engine", "ejs");

const checker = new FrequencyChecker(conf)

app.get('/', (req, res) => {
  res.render('index', conf);
})

app.get('/config.json', async (req, res) => {
  const sample = {}
  for (let i = 0; i < conf.blockchains.length; i++) {
    const chainConf = conf.blockchains[i]
    const wallet = await DirectSecp256k1HdWallet.fromMnemonic(chainConf.sender.mnemonic, chainConf.sender.option);
    const [firstAccount] = await wallet.getAccounts();
    sample[chainConf.name] = firstAccount.address
  }

  const project = conf.project
  project.sample = sample
  project.blockchains = conf.blockchains.map(x => x.name)
  res.send(project);
})

app.get('/balance/:chain', async (req, res) => {
  const { chain } = req.params

  let balance = {}

  try {
    const chainConf = conf.blockchains.find(x => x.name === chain)
    if (chainConf) {
      const rpcEndpoint = chainConf.endpoint.rpc_endpoint;
      const wallet = await DirectSecp256k1HdWallet.fromMnemonic(chainConf.sender.mnemonic, chainConf.sender.option);
      const client = await SigningStargateClient.connectWithSigner(rpcEndpoint, wallet);
      const [firstAccount] = await wallet.getAccounts();
      await client.getBalance(firstAccount.address, chainConf.tx.amount[0].denom).then(x => {
        balance = x
      }).catch(e => console.error(e));
    }
  } catch (err) {
    console.log(err)
  }
  res.send(balance);
})

app.get('/send/:chain/:address', async (req, res) => {
  if (mutex.isLocked()) {
    return res.status(500).send({ result: 'Faucet is busy, Please try again later.' })
  }
  const { chain, address } = req.params;
  const ip = req.headers['cf-connecting-ip'] || req.headers['x-real-ip'] || req.headers['X-Forwarded-For'] || req.ip
  console.log('request tokens from', address, ip)
  if (chain || address) {
    try {
      const chainConf = conf.blockchains.find(x => x.name === chain)
      if (chainConf && address.startsWith(chainConf.sender.option.prefix)) {
        if (await checker.checkAddress(address, chain) && await checker.checkIp(`${chain}${ip}`, chain)) {
          checker.update(`${chain}${ip}`) // get ::1 on localhost
          console.log('send tokens to ', address)
          await mutex.runExclusive(async () => {
            await sendTx(address, chain).then(ret => {
              console.log(ret)
              checker.update(address)
              res.send({ result: { code: ret.code, tx_hash: ret.transactionHash, height: ret.height } })
            }).catch(err => {
              res.status(500).send({ result: `err: ${err}` })
            });
          });
        } else {
          res.status(429).send({ result: "You requested too often" })
        }
      } else {
        res.status(400).send({ result: `Address [${address}] is not supported.` })
      }
    } catch (err) {
      console.error(err);
      res.status(500).send({ result: 'Failed, Please contact to admin.' })
    }

  } else {
    // send result
    res.send({ result: 'address is required' });
  }
})

app.listen(conf.port, () => {
  console.log(`Faucet app listening on port ${conf.port}`)
})

async function sendTx(recipient, chain) {
  const chainConf = conf.blockchains.find(x => x.name === chain)
  if (chainConf) {
    const wallet = await DirectSecp256k1HdWallet.fromMnemonic(chainConf.sender.mnemonic, chainConf.sender.option);
    const [firstAccount] = await wallet.getAccounts();
    console.log("sender", firstAccount);

    const rpcEndpoint = chainConf.endpoint.rpc_endpoint;
    const client = await SigningStargateClient.connectWithSigner(rpcEndpoint, wallet);

    const amount = chainConf.tx.amount;
    const fee = chainConf.tx.fee;
    console.log("recipient", recipient, amount, fee);

    let response = client.sendTokens(firstAccount.address, recipient, amount, fee)
    console.log(response)
    return response;
  }
  throw new Error(`Blockchain Config [${chain}] not found`)
}
