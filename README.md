# BitScreen Overview

## What is BitScreen?

BitScreen is a distributed file filtering system for Filecoin Network Storage Providers. Its aim is to give Storage Providers on the Filecoin Network granular control over files being stored and retrieved through their Lotus node. Using BitScreen, a Storage Provider can choose to prevent storage and retrieval deals related to specified CIDs, in order to mitigate the risk of hosting potentially illegal or otherwise high-risk content. 

BitScreen’s core functionality is that a user can create their own lists of CIDs for which future storage and retrieval deals will be declined. They can also publicly share their CID lists so that other Storage Providers can use them, and import filter lists maintained by third parties, including other Storage Providers or content assessors (parties who have inspected files, assessed their risk under some body of laws or rules, and are making the resulting CID list available). This sharing function will allow Storage Providers to leverage the collective intelligence of the Filecoin Network, and outside stakeholders offering their assessments of files, to mitigate negative effects of hosting potentially illegal or otherwise high-risk content.

## Why is BitScreen useful?

BitScreen starts with the recognition that in some cases the contents of a file leads to risk for its host, the Storage Provider (and perhaps others too). Sometimes that risk is legal in nature – the contents of a given file may be illegal to host, depending on where a Storage Provider is located. The type of legal risk might be civil (leading a SP to owe monetary damages to some other private party) or criminal (leading to potential criminal penalties from a government actor) or some combination of them. The legal risk might be fairly clear or depend on an ambiguous set of facts or untested legal theory. In addition to legal risk, there are other types of risk, such as brand risk. Hosting files with contents that are extreme, taboo, or controversial may lead to brand damage for a Storage Provider. In any case, a Storage Provider must assess these risks and decide how to act as a result. We expect that Storage Providers, based on diversity factors including location, size, values, and other activities, will arrive at a range of different decisions based on these risk assessments. BitScreen aims to help Storage Providers implement those decisions.

We believe that making this functionality available to Storage Providers is a necessary part of mitigating their operational risk and allowing the Filecoin Network to mature in compliance with existing law. At the same time, by making this tool available, we don’t want to encourage Storage Providers (who are fundamentally infrastructure providers) to engage in more moderation than is necessary based on their assessment of legal risk and their values, and underlying web3 values. Content moderation was once primarily the concern of social media platforms and other consumer facing applications. However, infrastructure providers - including at the storage level - are increasingly receiving legal and other complaints related to content hosted by third parties using their systems. BitScreen is a practical response to that developing reality.

Whether complaints relate to potential copyright infringement (as under the Digital Millennium Copyright Act) or other types of objectionable content, Filecoin Storage Providers may find it desirable to take preventive action against specific files for a variety of reasons. This may include: compliance with legal obligations; reduction of risk to their business; prevention of harms to persons affected by the contents of malicious files; or removal of files that otherwise violate the content policies of the storage provider.

BitScreen provides a suite of tools for Storage Providers to effectively address these legal and other concerns related to files stored on their systems.

## How does BitScreen work?

BitScreen is composed of several components that storage providers can use in the configuration they prefer. These include the BitScreen Lotus Plugin, a cloud-based List Manager (available as a GUI or CLI), a Local CID List, and the List Updater.

If a storage provider receives an apparently valid complaint or legal request to remove, disable, or block a piece of content stored on their system and associated with a CID (or for some other reason chooses to make a particular file unavailable for retrieval), the storage provider may choose to add the CID to a filter list using the cloud-based List Manager, or to the Local CID List. The BitScreen Plugin then works by comparing a given Payload ContentID (CID) against the active list of known CIDs. Payload CIDs found in a BitScreen filter list will activate the Filecoin DealMaking filter and will be filtered out from deals. 

On the back end, the cloud-based List Manager communicates with the BitScreen Lotus Plugin via the List Updater to efficiently keep current filter lists up to date. With the List Manager, users of BitScreen can keep their filter lists private, share them directly with another user, or add them to a public directory where they may be imported by any user. Alternatively, storage providers can make manual entries of known CIDs in their Local CID List, without the need to run the List Manager or List Updater.

In this way, BitScreen acts as an external program which provides input to the DealMaking filter function in Filecoin’s Lotus Miner. As a result, clients seeking to store a file with a filtered CID will not be able to form a storage deal with the storage providers. Similarly, clients seeking to retrieve a file already stored by the storage providers, and that has a filtered CID, will not be able to retrieve the file.

## Who is the development team?

[Murmuration Labs](https://murmuration.ai/) is a policy and product development studio consisting of Trust & Safety professionals with over a decade of combined experience in handling content moderation, tech policy, and legal compliance for social media and other platforms. We help develop product solutions for blockchain and other tech companies dealing with issues in this space.

[Keyko](https://www.keyko.io/) is full-service engineering support provider specializing in technical solutions for Web 3.0 digital ecosystems, based in Switzerland.

---

# BitScreen Installation & Usage

  * [BitScreen installation guide](https://github.com/Murmuration-Labs/bitscreen/blob/master/bitscreen_installation_guide.md)
  * [BitScreen CLI guide](https://github.com/Murmuration-Labs/bitscreen-cli/blob/main/README.md)
  * BitScreen GUI guide
