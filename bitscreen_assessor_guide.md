# BitScreen Assessor Quickstart Guide

Welcome to the BitScreen suite of tools for evaluating and taking action regarding content complaints received for files hosted on the Filecoin distributed storage network.

* [More detailed technical information about BitScreen](https://github.com/Murmuration-Labs/bitscreen)

## User typess

There are two types of users of BitScreen:

1. Storage providers who are operating a node on the Filecoin network
2. Assessors who evaluate and act on complaints received regarding Filecoin file contents

## Assessors

As an assessor, you will be responsible for evaluating content complaints received regarding files stored on the Filecoin network. Using your domain expertise, you will make decisions about the apparent facial validity of these complaints and decide whether to add the CIDs to filter lists, to which storage providers can subscribe. (Some storage providers may also share their own assessments with peers in the network as well in the same manner.)

* [How are complaints assessed?](https://github.com/Murmuration-Labs/looking-glass-docs/blob/main/how-are-complaints-assessed.md)
* [How to become an assessor](https://github.com/Murmuration-Labs/looking-glass-docs/blob/main/apply-to-become-an-assessor.md)

## Tools

To review complaints and associated file CIDs, you will use a combination of cloud-based tools, including Rodeo and Bitscreen. 

* Rodeo is the tool where complaints are received and assessed in a queue.
* Bitscreen is a filter list management tool that allows storage providers and node operators to subscribe to lists and filter CIDs out of storage and retrieval deals.

(*Note:* Since you will likely not be running your own node - unless you are also a storage provider - you will not need to worry about running [Filecoin Lotus](https://lotus.filecoin.io/), but may wish to familiarize yourself with it at a high level. Storage providers will also run the BitScreen Filecoin Lotus plugin, but assessors will generally not need to.)

Your assessments in Rodeo will then be made public through a transparency portal called Looking Glass. This portal contains a public record of the complaint, the CIDs, and the actions taken by assessors, such as adding items to a filter list.

## Getting started

When evaluating a complaint in Rodeo, you will:

* Check the complaint and file CID(s) to ensure that the complaint is relevant to the reported CID(s)
* Review the file contents and assess whether the complaint appears to be valid on its face, and if applicable, whether it is in violation of a relevant content policy
* Make a determination about whether or not to add the CID(s) to the filter list
* Add the CID(s) to the filter list if it appears facially valid and violates the content policy
* Complete your assessment in Rodeo, and verify that it has been made public in Looking Glass

It is important to note that the filter list is a tool that helps to improve the overall security and stability of the Filecoin network. Storage providers and node operators have the choice to subscribe to the filter list, but it is ultimately their responsibility to ensure that the files they store and retrieve are in compliance with applicable content policies.

* [More detailed information about BitScreen filter list management](https://github.com/Murmuration-Labs/bitscreen)
* [More information about Rodeo assessment tool]()
* [More info about Looking Glass transparency hub](https://github.com/Murmuration-Labs/looking-glass-docs)

Thank you for your contribution to maintaining a safe and secure Filecoin network!
