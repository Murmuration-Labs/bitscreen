# Bitscreen Installation & Usage Tutorial

This guide presents the steps a Filecoin storage provider must take to install and use the BitScreen tools. Using Bitscreen, the storage provider can prevent certain files (decided by the storage provider) from being stored and/or retrieved.

This guide assumes that you are already up and running as a Filecoin storage provider, and already have an active and working Filecoin Lotus node. If this is not the case, please refer to the [steps provided by Protocol Labs here](https://lotus.filecoin.io/docs/set-up/install).


## BitScreen Prerequisites

1. [GO version 1.17.1](https://go.dev/dl/#go1.17.1)
    
    * Debian based: `sudo apt update && sudo apt upgrade && sudo apt install golang-go`
 
2. [Python 3.9](https://www.python.org/downloads/release/python-390/)
    
    * Debian based: `sudo apt update && sudo apt install software-properties-common && sudo add-apt-repository ppa:deadsnakes/ppa && sudo apt install python3.9`
 
    * **Note:** Any Python 3 version 3.7+ will work.
 
3. [pip for Python 3](https://pip.pypa.io/en/stable/installation/) (In some cases it may not come with Python out of the box).  

    * **Note:** The versions of hat are currently fully functional in the installation process are pip 20.2.3 & pip 20.3.4.
    
    * Debian based: `sudo apt install python3-pip`

4. ZeroMQ library - used for the communication between the Bitscreen Plugin and the Bitscreen Updater

    * Debian based: `sudo apt install libzmq3-dev`


## Installing BitScreen-CLI

When installing the BitScreen-CLI please make sure you do not use Python 3.6 while using pip (version 3.6 comes pre-installed on different operating systems), because the installation might not succeed. Please make sure you are using a version equal or higher to 3.7. Usually stating the exact version of Python3 should work when using the pip module:

  * `python3.9 -m pip install bitscreen-cli --upgrade`

This should install the latest version of the BitScreen-CLI which at the moment of writting this guide is v. 0.1.16.

  * **Note:** If the "bitscreen-cli" command is not found, you may be missing the python binary path from your `$PATH`. You can add it using: `export PATH=$PATH:/home/$USER/.local/bin`


## Registration via BitScreen-CLI & BitScreen GUI Client

To obtain a BitScreen account you can use either of the two following options:

1. Directly through the BitScreen-CLI, using the command `bitscreen-cli auth register WALLET-ADDRESS` (replace WALLET-ADDRESS with the appropriate value for the Ethereum wallet address you want to use to log into and manage BitScreen)

2. Through the BitScreen GUI Client which can be found at: https://app.bitscreen.co. 

    * You will need a browser that has the MetaMask extension installed in order to do so, along with an Ethereum wallet address. 
    * If you have multiple addresses associated with the same wallet, please make sure you take notice of which one is used to register as it will matter in the following step.


## Authentication of BitScreen-CLI

1. Obtain the information of the account created in the previous step.

    * This can be done through the usage of either a Private Key of the account or using the Recovery Phrase in case the account is associated directly to the master key. 

    * **Note:** If the MetaMask account used to register on BitScreen is a child key derived from the master key associated with the recovery phrase (seed phrase) then the Private Key of the said account must be used. In this case using the Recovery Phrase will not work.

    * Use the links bellow to learn how to obtain either of these:

      - [Account Private Key from Metamask](https://metamask.zendesk.com/hc/en-us/articles/360015289632-How-to-Export-an-Account-Private-Key) 
      - [Account Recovery Phrase from Metamask](https://metamask.zendesk.com/hc/en-us/articles/360015290032-How-to-reveal-your-Secret-Recovery-Phrase)

2. Authenticate in the BitScreen-CLI

    * With Private Key:

      1. `bitscreen-cli auth login`
      2. A prompt will appear which will ask for the wallet address associated with the account
      3. A prompt will appear which will ask for the account Private Key

    * With Recovery Phrase (Seed Phrase):

      1. `bitscreen-cli auth login --fromseed`
      2. A prompt will appear which will ask for the account Recovery Phrase (Seed Phrase)

In both cases, after being successfully authenticated you can optionally save the credentials for future logins.


## Run BitScreen-CLI Setup Installation

Before running the below command make sure the environment variable `LOTUS_MINER_PATH` is set. This variable should have been set during the Lotus Node & Miner installation process.

  * Then run the command: `bitscreen-cli setup install`

This process will install and configure two of the main components of the filtering process: the BitScreen Updater and the BitScreen Plugin. 

  * The BitScreen Updater tool is used to fetch and store the lists of filters you want to use from the server, and to provide them to the BitScreen Plugin, which in turn communicates with the Lotus Node.
  * **Note:** During setup, you will have the option to use the BitScreen-CLI login credentials for the Updater as well. We recommend doing so!

To confirm that the BitScreen setup was done properly in Lotus, please check the following: 

  * Verify that the variables 'Filter' and 'RetrievalFilter' (which should be found in `~/.lotusminer/config.toml`, if not previously configured otherwise) point towards the BitScreen GO plugin. 
  * Debian based should be: `/home/USERNAME/go/bin/bitscreen`.


## Run the BitScreen Updater

  * Command: `bitscreen-updater start`

  * Check if the tool was properly started by using the command: `bitscreen-updater status`

  * The logs which will show the BitScreen Updater's activity should be found at the following path: `/tmp/bitscreen_updater.log`


## 7. Filtering Retrieval and Storage Deals

You can now create filter lists which your Lotus node will use to block storage or retrieve deals from happening for affected CIDs.

This can be done directly in the CLI:
  * More information: https://github.com/Murmuration-Labs/bitscreen-cli#bitscreen-cli-filter
  * Command: `bitscreen-cli --help`

Filter Lists can also be created and managed through the GUI web client (https://app.bitscreen.co). 

  * More information: https://github.com/Murmuration-Labs/bitscreen/blob/master/README.md 
  * The Settings page on the GUI client has a quick start guide to help you get started.
