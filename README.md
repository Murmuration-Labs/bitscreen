# BitScreen

## Introduction

BitScreen is a tool for Filecoin node operators. It’s purpose is to allow node operators to prevent a particular file (or files) from being stored or retrieved (or both). Node operators may want to do this if a particular file contains material that is illegal, or violates the node operators content policy -- that is, if the file contains content that the node operator does not want to host or does not want to serve. By adding a CID to Bitscreen, a node operator can decline to host or allow retrieval of files.

Bitscreen works by comparing a Payload ContentID (CID) against a selected list of known CIDs. Payload CIDs found in the BitScreen list will activate the Filecoin [DealMaking filter](https://github.com/filecoin-project/filecoin-docs/blob/master/docs/mine/lotus/miner-configuration.md#dealmaking-section) and will be filtered out from deals. Node operators may use this functionality to enforce their content policy, where appropriate, as well as satisfy other legal or compliance requirements.

*See also*: [Content Policy Guide for Filecoin Node Operators](https://github.com/Murmuration-Labs/filecoin-node-operator-kit/blob/main/Content-Policy-Guide.md)

## How It Works

If a node operator receives an apparently valid complaint or legal request to remove, disable, or block a piece of content stored on their system and associated with a CID (or for some other reason chooses to make a particular file unavailable for retrieval), the node operator may choose to add the CID to their local BitScreen list.

Node operators may add a CID manually to their local BitScreen list by adding the CID to the file located at `lotus/.murmuration/bitscreen`

BitScreen then acts as an external program which provides input to the DealMaking filter function in Filecoin’s Lotus Miner. Node operators can use this input to block storage and retrieval deals for known CIDs. As a result, clients seeking to store a file with a filtered CID will not be able to form a storage deal with the Node Operator. Similarly, clients seeking to retrieve a file already stored by the node operator, and that has a filtered CID, will not be able to retrieve the file.

*See also*: [Filecoin Lotus Miner Filter](https://github.com/filecoin-project/filecoin-docs/blob/master/docs/mine/lotus/miner-configuration.md#dealmaking-section)

## Upcoming Features

Future updates to the BitScreen tool will enable sharing lists between node operators, and the use of lists based on third party databases.
