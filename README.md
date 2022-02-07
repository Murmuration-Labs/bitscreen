# BitScreen Overview

  * [BitScreen Installation & Usage](https://github.com/Murmuration-Labs/bitscreen/blob/master/README.md#bitscreen-installation--usage)

## What is BitScreen?

[BitScreen](https://bitscreen.co/) is a distributed file filtering system for Filecoin Network Storage Providers. Its aim is to give Storage Providers on the Filecoin Network granular control over files being stored and retrieved through their Lotus node. Using BitScreen, a Storage Provider can choose to prevent storage and retrieval deals related to specified CIDs, in order to mitigate the risk of hosting potentially illegal or otherwise high-risk content. 

BitScreen’s core functionality is that a user can create their own lists of CIDs for which future storage and retrieval deals will be declined. They can also publicly share their CID lists so that other Storage Providers can use them, and import filter lists maintained by third parties, including other Storage Providers or content assessors (parties who have inspected files, assessed their risk under some body of laws or rules, and are making the resulting CID list available). This sharing function will allow Storage Providers to leverage the collective intelligence of the Filecoin Network, and outside stakeholders offering their assessments of files, to mitigate negative effects of hosting potentially illegal or otherwise high-risk content.

  * *See also:* [BitScreen official site](https://bitscreen.co/)

## Why is BitScreen useful?

BitScreen starts with the recognition that in some cases the contents of a file leads to risk for its host, the Storage Provider (and perhaps others too). Sometimes that risk is legal in nature – the contents of a given file may be illegal to host, depending on where a Storage Provider is located. The type of legal risk might be civil (leading a SP to owe monetary damages to some other private party) or criminal (leading to potential criminal penalties from a government actor) or some combination of them. The legal risk might be fairly clear or depend on an ambiguous set of facts or untested legal theory. In addition to legal risk, there are other types of risk, such as brand risk. Hosting files with contents that are extreme, taboo, or controversial may lead to brand damage for a Storage Provider. In any case, a Storage Provider must assess these risks and decide how to act as a result. We expect that Storage Providers, based on diverse factors including location, size, values, and other activities, will arrive at a range of different decisions based on these risk assessments. BitScreen aims to help Storage Providers implement those decisions in the maner Storage Providers see fit.

We believe that making this functionality available to Storage Providers is a necessary part of mitigating their operational risk and allowing the Filecoin Network to mature in compliance with existing law. At the same time, by making this tool available, we don’t want to encourage Storage Providers (who are fundamentally infrastructure providers) to engage in more moderation than is necessary based on their assessment of legal risk and their values, and underlying web3 values. Content moderation was once primarily the concern of social media platforms and other consumer facing applications. However, infrastructure providers - including at the storage level - are increasingly receiving legal and other complaints related to content hosted by third parties using their systems. BitScreen is a practical response to that developing reality.

Whether complaints relate to potential copyright infringement (as under the Digital Millennium Copyright Act) or other types of objectionable content, Filecoin Storage Providers may find it desirable to take preventive action against specific files for a variety of reasons. This may include: compliance with legal obligations; reduction of risk to their business; prevention of harms to persons affected by the contents of malicious files; or removal of files that otherwise violate the content policies of the storage provider.

BitScreen provides a suite of tools for Storage Providers to effectively address these legal and other concerns related to files stored on their systems.

## How does BitScreen work?

BitScreen is composed of several components that storage providers can use in the configuration they prefer. These include the BitScreen Lotus Plugin, a cloud-based List Manager (available as a GUI or CLI), a Local CID List, and the List Updater.

If a storage provider receives an apparently valid complaint or legal request to remove, disable, or block a piece of content stored on their system and associated with a CID (or for some other reason chooses to make a particular file unavailable for retrieval), the storage provider may choose to add the CID to a filter list using the cloud-based List Manager, or to the Local CID List. The BitScreen Plugin then works by comparing a given Payload ContentID (CID) against the active list of known CIDs. Payload CIDs found in a BitScreen filter list will activate the Filecoin DealMaking filter and will be filtered out from deals. 

  * *See also*: [Filecoin Lotus Miner Filter](https://lotus.filecoin.io/docs/storage-providers/config/#dealmaking-section)

On the back end, the cloud-based List Manager communicates with the BitScreen Lotus Plugin via the List Updater to efficiently keep current filter lists up to date. With the List Manager, users of BitScreen can keep their filter lists private, share them directly with another user, or add them to a public directory where they may be imported by any user. Alternatively, storage providers can make manual entries of known CIDs in their Local CID List, without the need to run the List Manager or List Updater.

In this way, BitScreen acts as an external program which provides input to the DealMaking filter function in Filecoin’s Lotus Miner. As a result, clients seeking to store a file with a filtered CID will not be able to form a storage deal with the storage providers. Similarly, clients seeking to retrieve a file already stored by the storage providers, and that has a filtered CID, will not be able to retrieve the file.

## Who is the development team?

[Murmuration Labs](https://murmuration.ai/) is a policy and product development studio consisting of Trust & Safety professionals with over a decade of combined experience in handling content moderation, tech policy, and legal compliance for social media and other platforms. We help develop product solutions for blockchain and other tech companies dealing with issues in this space.

[Keyko](https://www.keyko.io/) is full-service engineering support provider specializing in technical solutions for Web 3.0 digital ecosystems, based in Switzerland.

---

# BitScreen Installation & Usage

  * [BitScreen installation guide](https://github.com/Murmuration-Labs/bitscreen/blob/master/bitscreen_installation_guide.md)
  * [BitScreen CLI guide](https://github.com/Murmuration-Labs/bitscreen-cli/blob/main/README.md)
  * [BitScreen GUI guide](https://github.com/Murmuration-Labs/bitscreen/blob/master/README.md#bitscreen-gui-user-guide) (below)

---

## BitScreen GUI User Guide

### Installation  

To get started, first install the command line interface (CLI) from the terminal, by following the instructions in our installation guide:  

  * [https://github.com/Murmuration-Labs/bitscreen/blob/master/bitscreen\_installation\_guide.md](https://github.com/Murmuration-Labs/bitscreen/blob/master/bitscreen_installation_guide.md)  

Lotus Miner must be restarted once the BitScreen Plugin and other components are installed.

### BitScreen Components:  

BitScreen has the following components, which operators can use in the combination they prefer:  

1. **Lotus Plugin (required):** Prevents deals in Lotus for CIDs contained in Local CID List or List Manager.
2. **Local CID List:** Flat file containing CIDs to be filtered. Can be edited manually. Acts as fallback if List Updater not in use. 
3. **List Updater:** Daemon that periodically checks List Manager for presence of CIDs requested by Lotus Plugin. If installed, overrides Local CID List.
4. **List Manager (optional):** Advanced cloud-based utility to create, share & import filter lists
    * [GUI client](http://app.bitscreen.co/)
    * [CLI](https://github.com/Murmuration-Labs/bitscreen-cli/blob/main/README.md)

### Modes of use:  

To use BitScreen, run the configuration script and install the required components.   

* **Basic:** If you're using the Plugin with the Local CID list only, you can simply edit the CIDs contained on the list manually. You won't be able to share or import lists from other users.   

* **Advanced:** If you're using the Plugin with the List Updater daemon, you will also need to run the List Manager, either from the command line, or using the GUI client.

### Connecting a wallet  

The cloud-based BitScreen List Manager (both the GUI and CLI) uses Metamask as a login provider. Your Ethereum wallet address is stored hashed as a unique identifier within our system to control access, and to determine statistical usage information. The instructions below will walk you through set-up of the GUI. Instructions for [set-up and use of the CLI](https://pypi.org/project/bitscreen-cli/) are here.  If you are only using the Local CID List without the List Manager or List Updater, you can skip these steps.   

1. To begin, navigate to [app.bitscreen.co](https://app.bitscreen.co/) and activate the filtering toggle to turn on BitScreen. 
2. Click Connect with Metamask. You will be asked to sign a message in [Metamask](https://metamask.io/) to verify the wallet. You will not be charged for this or any wallet interaction in BitScreen. It is only used for authenticating you as a user.
3. Once you've connected Metamask, your wallet address will appear on [app.bitscreen.co/settings](https://app.bitscreen.co/settings) and in the navigation bar. You can log out by disconnecting your wallet at any time. 
4. Next, run the configuration script described in the Installation section of this user guide, if you have not already done so. 
5. In Metamask, you will need to export either a [private key](https://metamask.zendesk.com/hc/en-us/articles/360015289632-How-to-Export-an-Account-Private-Key) or a [seed phrase](https://metamask.zendesk.com/hc/en-us/articles/360015290032-How-to-reveal-your-Secret-Recovery-Phrase), which you will then input into the BitScreen updater to connect it to your cloud-based List Manager account.

### Activating importing & sharing  

As a user of the cloud-based List Manager, you may use it to maintain and run your own private filter lists without adding any additional user information. Private lists are never accessible to other users.   

You also have the option to activate importing lists created by other users, and to activate sharing your lists with others. Neither of these sets of functionality are obligatory to use BitScreen. If you choose to activate them, BitScreen requires a few additional points of information to be added for the purposes of transparency to other end users.   

1. From the settings page, click the Activate Importing Lists toggle to turn on importing. This will enable you to import and run lists created by other users. 
2. You will be asked to add your country data for statistical usage purposes to use importing.
3. Click the Activate Sharing Lists toggle to turn on sharing. This will enable you to share your lists with others directly, and to make them available for all users to import via the Directory. 
4. You will be asked to add additional contact information for your business in order to use sharing, and your list provider data will be visible in the Directory for public lists.

### Deleting & exporting your account  

If you wish to cease using the BitScreen List Manager, you may do so at any time following the instructions below.  

1. From the settings page, you can delete the BitScreen account linked to your wallet, or export your filter lists. 
2. Exports consist of the following:
    1. Any list provider data you entered
    2. In cases where you are the owner/author of a filter list, the full list of CIDs included in each list
    3. In cases where you are importing a list owner by another user, the name(s) of any imported list(s) and their URL.   

The above only applies to the cloud-based List Manager, and does not apply to the Local CID List. If you wish to cease using the local list, simply delete or erase the contents of that local file. 

### BitScreen Dashboard  

Once you have activated the cloud-based List Manager, you will be able to view relevant performance data on the BitScreen dashboard at [app.bitscreen.co/dashboard](https://app.bitscreen.co/dashboard). The dashboard displays the following information:  

1. Total number of CIDs currently being filtered by your node
2. Number of subscribers to your shared or public lists
3. Total number of storage and retrieval deals (combined) which have been declined due to filtering of CIDs on your lists
4. Number of:
    1. Active and inactive lists (you can deactivate lists temporarily to stop them from filtering)
    2. Imported lists created by other users
    3. Private lists accessible only to you
    4. Public lists accessible via the Directory to all users
5. A chart with customizable timeframes comparing:
    1. Total of CIDs currently being filtered
    2. Total number of requests blocked involving those CIDs  

### Creating new filter lists  

1. To create a new filter list, navigate to [app.bitscreen.co/filters](https://app.bitscreen.co/filters) and click New filter, or navigate directly to [app.bitscreen.co/filters/new](https://app.bitscreen.co/filters/new). 
2. Add a filter name and description. Bear in mind that, if you choose to make the list shared or public, this information will be visible to other users.
3. Click Add CID, and then choose whether you want to add one or multiple CIDs to the filter list. 
4. You will then be asked to add the CID itself (not the URL where a CID appears). 
5. Optionally, if there is a public record of the complaint for transparency purposes (such as may be posted on [Lumen](https://www.lumendatabase.org/), Github, or elsewhere), you may add it to the field marked Public complaint URL. 
  1. Note: If you bulk input CIDs, and add a public complaint URL, the complaint address will be applied to all the inputted CIDs.
6. If you make an error while entering CIDs or need to change something later, you can edit individual CID entries by clicking the edit icon, or can edit them as a batch using bulk actions on any selected entries. 
7. CIDs, likewise, can be individually or bulk deleted from a list by using the delete icon, or the bulk actions delete button. 
8. CIDs can also be individually moved from one list to another using the move icon. 
9. Under list Type, choose whether to make the list Private, Shared, Public, or an exception list. More information on each can be found in the next section. 
10. Click Save changes to save your list. 

### List types  

There are four types of filter lists in the BitScreen List Manager: private, shared, public, and exception lists. In order to use shared and public lists, you will need to activate list sharing functionality and enter the requested information on the settings page.  

1. Private: Private lists are only visible to you.
2. Shared: Shared lists are only visible to other users if they have the URL.
    1. To share a list, first save it, and then click Copy link to get the shareable list URL
3. Public: Public lists are visible to all users via the directory, along with the list provider information you entered in settings. 
    1. To share a public list, first save it, and then click Copy link to get the shareable URL. 
4. Exception: Exception lists prevent CIDs from imported lists from being filtered. They cannot be shared.
    1. Exception lists are used to override filtering of a given CID(s) which are being blocked by an imported filter list. Rather than disable or discard an imported list entirely, creating an exception list lets you keep using the list, while excluding desired CIDs from filtering. 
    2. Note: Exception lists are intended to be used to override filtering of CIDs from imported lists only. If you enter a CID into an exception list that you are filtering yourself on a list you own, you will be asked to resolve the conflict.

### Importing filter lists  

Once you have activated importing lists from the settings page, and entered the required information, you can then import and run lists shared by other BitScreen users. When you import a list, you are considered to be a subscriber of that list. The number of active subscribers for a given list is shown on the My filters page, the filter list detail page, and the Directory, if applicable.  

1. To import a list directly from its URL, go to app.bitscreen.co/filters and click Import filter.
2. You will be asked to input the URL of the filter, whether it is a direct shared filter or a public filter. You can also enter private notes when you import a filter, which will only be accessible to you. 
3. To import a list from the public Directory, go to app.bitscreen.co/directory, and choose the list you want. You can click the Import button directly from there, or you can view the list details page for more information about the list and provider. From the list details page, click Import filter, add any notes desired (optional), and click Fetch remote filter. Review the provided details to make certain, and then either import the filter, discard it, or cancel. 

### Activating & deactivating filter lists  

For troubleshooting purposes, you may encounter situations where it is necessary to disable one or more filter lists running on your node. Active filters run on your node to prevent storage and retrieval deals with included CIDs. Inactive filters have no effect.   

1. From the My filters main page, click the Active toggle to the on or off position as desired. 
2. Alternatively, from the filter detail page of a given filter, choose Active or Inactive as desired. 
3. Note: If you deactivate a filter that is shared or public and has active subscribers, you will be asked whether you want to only deactivate the list for yourself, or for all subscribers.

### Deleting & discarding filter lists  

1. As the owner/author of a filter list, to delete it, go to the list detail page and click the three dot menu next to the save button, and choose delete. You can also delete a list from the My filters main page. 
2. If you attempt to delete a list that has active subscribers, you will be notified that those users will be impacted, and asked to confirm that you want to proceed. When you delete a list with subscribers, it will appear as orphaned to those subscribers (see next section).
3. When you import a filter list, you do not have the ability to delete the list, since you are not the owner of it. You can, however, discard the list from running on your node altogether. Either go to the filter list details page of the imported list, and click the three dot menu, followed by discard. Or go to the My filters main page, and click the trash can icon to discard it.   

### Orphaned lists  

When the owner of a filter list deletes a list that has active subscribers, an entry for the list will still remain on the My filters page of subscribers, with the status of "Orphaned." The list will not be functional, and cannot be reactivated. It is preserved on the My filters page of subscribers in order to show them that a change has occurred remotely which impacts them.   

If necessary, subscribers may communicate with the list provider via their public contact details in the Directory in order to find out what happened, and receive a new replacement list URL, if there is one.
