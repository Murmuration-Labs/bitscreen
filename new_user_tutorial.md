# Bitscreen Installation & Usage Tutorial

This guide presents the steps a Filecoin storage provider must take to install and use the BitScreen tool. Through it the storage provider can prevent certain files (decided by the storage provider) from being stored and/or retrieved.

By following the steps provided in this guide it means you are already an 'up-and-running' Filecoin storage provider, therefore you already have an active and working Filecoin Lotus Node & Miner. If this is not the case, you can follow the steps provided by Protocol Labs here: https://lotus.filecoin.io/docs/set-up/install .


## 1. BitScreen Prerequisites:

a) GO version 1.17.1. Can be found here: https://go.dev/dl/#go1.17.1
    
  * Debian based: `sudo apt update && sudo apt upgrade && sudo apt install golang-go`
 
b) Python 3.9. Can be found here: https://www.python.org/downloads/release/python-390/
    
  * Debian based: `sudo apt update && sudo apt install software-properties-common && sudo add-apt-repository ppa:deadsnakes/ppa && sudo apt install python3.9`
 
  * *Note:* Any Python 3 version 3.7+ will work.
 
c) pip for Python 3 (in some cases it does not come with Python out of the box). Can be found here: https://pip.pypa.io/en/stable/installation/ (the versions of that 100% working in the installation process are pip 20.2.3 & pip 20.3.4).
    
  * Debian based: `sudo apt install python3-pip`

d) ZeroMQ library - used for the communication between the Bitscreen Plugin and the Bitscreen Updater

  * Debian based: `sudo apt install libzmq3-dev`


## 2. Installing - BitScreen-CLI

When installing the BitScreen-CLI please make sure you do not use Python 3.6 when using pip (the 3.6 version comes pre-installed on different operating systems) because the installation might not succeed. Make sure you are using a version higher or equal to 3.7. Usually stating the exact version of Python3 should work when using the pip module:

python3.9 -m pip install bitscreen-cli --upgrade

This should install the latest version of the BitScreen-CLI which at the moment of writting this guide is 0.1.16.

NOTE: If the "bitscreen-cli" command is not found, you may be missing the python binary path from your $PATH . You can add it like this: export PATH=$PATH:/home/$USER/.local/bin


## 3. Registration Process - BitScreen-CLI / BitScreen GUI Client

To obtain a BitScreen account you can use two options:

a) Directly through the BitScreen-CLI, using the command bitscreen-cli auth register WALLET-ADDRESS (replace WALLET-ADDRESS with the appropriate value for the wallet account you want to use on BitScreen)

b) Through the BitScreen GUI Client which can be found at: https://bxn.mml-client.keyko.rocks/. You will need a browser that has the MetaMask extension installed in order to do so. If you have multiple accounts associated with the same wallet, please make sure you take notice of which one is used to register as it will matter in the following step.


## 4. Authentication Process - BitScreen-CLI

a) Obtain the information of the account created at step 3
This can be done through the usage of either a Private Key of the account or using the Recovery Phrase in case the account is associated directly to the master key. NOTE! If the MetaMask account used to register on BitScreen is a child key derived from the master key associated with the Recovery Phrase (Seed Phrase) then the Private Key of the said account must be used. In this case using the Recovery Phrase will not work.

Use the links bellow to learn how to obtain either of these:
- Account Private Key from Metamask -> https://metamask.zendesk.com/hc/en-us/articles/360015289632-How-to-Export-an-Account-Private-Key
- Account Recovery Phrase from Metamask -> https://metamask.zendesk.com/hc/en-us/articles/360015290032-How-to-reveal-your-Secret-Recovery-Phrase

b) Authenticate in the BitScreen-CLI
With Private Key:
i. bitscreen-cli auth login
ii. A prompt will appear which will ask for the wallet address associated with the account
iii. A prompt will appear which will ask for the account Private Key

OR

With Recovery Phrase (Seed Phrase):
i. bitscreen-cli auth login --fromseed
ii. A prompt will appear which will ask for the account Recovery Phrase (Seed Phrase)

In both cases, after being successfully authenticated you will be able to opt for saving the credentials for future logins.


## 5. Run BitScreen-CLI setup installation

Before running the below command make sure the environment variable LOTUS_MINER_PATH is set. This variable should have been set during the Lotus Node & Miner installation process.
Command: bitscreen-cli setup install

This process is installing and configuring two of the main components of the filtering process: the BitScreen Updater and the BitScreen Plugin. The BitScreen Updater tool is used to fetch and store the lists of filters you want to use from the server and to provide them to the BitScreen Plugin which in turn communicates with the Lotus Node.

When this prompt appears during the setup process: Would you like to authenticate the BitScreen Updater with your CLI credentials? -> You can opt to use the credentials used to login into the BitScreen-CLI for the Updater as well. We recommend you doing so!

To confirm that the setup was done properly please check that the Lotus Miner was configured properly to use the BitScreen Suite. In order to do so you must check that the variables 'Filter' and 'RetrievalFilter' (which should be found in ~/.lotusminer/config.toml if not previously configured otherwise) point towards the BitScreen go plugin. Debian based should be: /home/USERNAME/go/bin/bitscreen.


## 6. Run the BitScreen Updater

Command: bitscreen-updater start

Check if the tool was properly started by using the command: bitscreen-updater status

The logs which will show the BitScreen Updater's activity should be found at the following path: /tmp/bitscreen_updater.log


7. Filtering Retrieval and Storage Deals

Create Filter Lists which your miner will use to block storage or retrieve deals from happening.

This can be done directly in the CLI. More information here:
- https://github.com/Murmuration-Labs/bitscreen-cli#bitscreen-cli-filter or
- using the command: bitscreen-cli --help

Filter Lists can also be created and managed through the GUI Web Interface (https://bxn.mml-client.keyko.rocks/). The Settings page has a quick start guide which you can check out to help you get started).
