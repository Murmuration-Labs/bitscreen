# BitScreen IPFS Port Proposal

## Background

[BitScreen](https://github.com/Murmuration-Labs/bitscreen) is a utility originally built for Filecoin Lotus storage providers (SP) to be able to create, share, and subscribe to custom filter lists in order to filter out known CIDs in a distributed fashion. BitScreen for Filecoin Lotus is currently in public beta. It functions presently with a Lotus plugin that interfaces with the dealmaking filter, and that checks either a local CID list, or interfaces with an updater daemon that checks a CLI/GUI list manager application running in the cloud for CIDS to filter. 

A follow-up suite of tools still under development will also enable SPs and other third party assessors to receive and assess legal and other complaints related to content hosted on the Filecoin network. Upon evaluation of complaints and affected content, SPs and assessors would be able to add reported CIDs to filter lists they manage, and publish their assessments to a transparency hub. Other SPs would be able to opt-in to subscribing to these lists, which would have the effect of filtering these CIDs from storage & retrieval deals on their Lotus node. BitScreen users can also choose to run their own local CID filtering lists if they prefer. 

Murmuration Labs (the developer of BitScreen) is currently also in the process of integrating third party hash databases of known illegal content (e.g., CSAM, terrorist materials) as list providers for BitScreen filtering. This will enable out of the box filtering of such materials for SPs who subscribe to these lists.

This work builds on the multi-assessor model for decentralized content moderation, described in the provisional [Songbird proposal](https://github.com/Murmuration-Labs/songbird-decentralized-moderation). Murmuration's aim is to provide the tools for many diverse assessors to apply their own public criteria to content moderation decisions (even where they may conflict with those of other assessors), and enable the greater community to leverage their expertise. 

## Proposal

Under the IPFS roadmap theme proposal outlined in issue [#64: Tools for node operators to configure what content they want to store/retrieve/provide](https://github.com/ipfs/roadmap/issues/64), Murmuration Labs proposes to port the work we have done for distributed content filtering on Filecoin to the IPFS network as well. 

Like BitScreen for Filecoin, the BitScreen IPFS port would operate under the same principles of voluntary opt-in for node operators, and make use of a multi-assessor framework, with transparency about complaints received, and assessments delivered. We intend to build this as a [plugin for go-ipfs](https://github.com/ipfs/go-ipfs/blob/master/docs/plugins.md).

Upon completion of the port to IPFS, BitScreen users would be able to create, share, and subscribe to third party lists for both Filecoin & IPFS all in one application (whether they choose to use CLI or GUI). Optionally, users seeking a greater degree of privacy could make use of local CID filtering lists if they choose, as appears to be the approach in [Cloudflare's related work on Safemode for IPFS gateways](https://blog.cloudflare.com/cloudflare-ipfs-safe-mode/).

Further, BitScreen for IPFS would be configurable to enable node operators to choose whether to run filtering at the gateway or p2p level (or both), depending on their compliance needs and business goals & values. 

Lastly, Murmuration Labs would also like to eventually enable web visitors to gateways who might encounter filtered CIDs to be able to inspect public records for any filtered CIDs they might encounter. Public records will consist of, at minimum, information about any relevant complaints received, the assessor(s) who reviewed it, and filter lists containing the affected CID(s), as well as any relevant data about assessment frameworks applied. 

## Request for feedback

Murmuration Labs is in the early stages of evaluating how best to proceed with porting BitScreen to work on both Filecoin & IPFS. We are requesting both general feedback on the above, as well as feedback specifically about proper pathways to develop the plugin for go-IPFS. Our technical team will be able to answer questions that might come up around this. 

Thanks in advance!
