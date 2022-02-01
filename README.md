# BitScreen

## Introduction

BitScreen is a tool for Filecoin storage providers. It’s purpose is to allow storage providers to prevent a particular file (or files) from being stored or retrieved (or both). Storage providers may want to do this if a particular file contains material that is illegal, or violates the node operators' content policy -- that is, if the file contains content that the storage providers does not want to host or does not want to serve. By adding a CID to BitScreen, a storage providers can decline to host or allow retrieval of files.

BitScreen works by comparing a Payload ContentID (CID) against a selected list of known CIDs. Payload CIDs found in the BitScreen list will activate the Filecoin [DealMaking filter](https://lotus.filecoin.io/docs/storage-providers/config/#dealmaking-section) and will be filtered out from deals. Storage providerss may use this functionality to enforce their content policy, where appropriate, as well as satisfy other legal or compliance requirements.

*See also*: [Content Policy Guide for Filecoin Node Operators](https://github.com/Murmuration-Labs/filecoin-node-operator-kit/blob/main/Content-Policy-Guide.md)

## How It Works

If a storage providers receives an apparently valid complaint or legal request to remove, disable, or block a piece of content stored on their system and associated with a CID (or for some other reason chooses to make a particular file unavailable for retrieval), the storage providers may choose to add the CID to their local BitScreen list.

Storage providers may add a CID manually to their local BitScreen list by adding the CID to the file located at `lotus/.murmuration/bitscreen`

BitScreen then acts as an external program which provides input to the DealMaking filter function in Filecoin’s Lotus Miner. Node operators can use this input to block storage and retrieval deals for known CIDs. As a result, clients seeking to store a file with a filtered CID will not be able to form a storage deal with the storage providers. Similarly, clients seeking to retrieve a file already stored by the storage providers, and that has a filtered CID, will not be able to retrieve the file.

*See also*: [Filecoin Lotus Miner Filter](https://lotus.filecoin.io/docs/storage-providers/config/#dealmaking-section)

## Upcoming Features

Future updates to the BitScreen tool will enable sharing lists between node operators, and the use of 
lists based on third party databases.
