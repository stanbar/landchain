# Blockchain

To start the HL Fabric test-netwrok with two organizations setup, two peers and ordering service, and 2 certificate authorities.

## Install Fabric and Fabric Samples

Follow the steps from https://hyperledger-fabric.readthedocs.io/en/latest/install.html

## Run the HL Fabric network

Follow the steps from https://hyperledger-fabric.readthedocs.io/en/latest/test_network.html

1. `cd ~/fabric-samples/test-network`
2. Run `./network.sh up createChannel -ca`

## Deploy scripts

From now, you can deploy chaincode (smartcontract):

1. From directory `/fabric-samples/test-network`
2. `./network.sh deployCC -ccn basic -ccp ../asset-transfer-basic/chaincode-javascript/ -ccl javascript`