Welcome to BitMEME (BTM) World!
ISC License GoDoc

bitmemed is the reference full node BitMEME (BTM) implementation written in Go (golang).

What is BitMEME (BTM)?
BTM is an attempt at a proof-of-work cryptocurrency with instant confirmations and sub-second block times. It is based on the PHANTOM protocol, a generalization of Nakamoto consensus, with the fastest POW L1 available. First MEME fork with Staking rewards.

Compilation Requirements
If you want to compile the binaries yourself (if you aren't sure whether you want to, then you probably don't), you would need Go 1.19 or later.

Setup
Mining BitMEME (BTM) requires two components: a node (bitmemed), and a miner. A third component is required to create and maintain a wallet. The node listens for new blocks while the miner is searching for blocks to report to the node. All three components are provided as stand alone files which require no installation.

You need to either download precompiled binaries, or compile the codebase yourself. The first option is recommended for most users.

Note that all bitmemed and the miner must be running in parallel. That is, each should be running from a different console and should not be distured as long as mining takes place.

Download Binaries
The easiest way to use bitmemed is to download the binaries from here. After downloading the binaries that fit your operating system, you should extract them to some folder.

Notice that the rest of the tutorial assumes that you installed from source, so before each command you run you should first run:

$ cd <THE_EXTRACTED_BINARIES_FOLDER>
Linux and Mac users might need to add ./ to any command so it'll run the corresponding binary. For example:

./bitmemed --utxoindex
Build from Source
Install Go according to the installation instructions here: http://golang.org/doc/install

Ensure Go was installed properly and is a supported version:

$ go version
Run the following commands to obtain and install bitmemed including all dependencies:
$ git clone https://github.com/bitmeme-taxi/bitmemed
$ cd bitmemed
$ go install . ./cmd/...
Bitmemed (and utilities) should now be installed in $(go env GOPATH)/bin. If you did not already add the bin directory to your system path during Go installation, you are encouraged to do so now.
Getting Started
Bitmemed has several configuration options available to tweak how it runs, but all of the basic operations work with zero configuration except the --utxoindex flag (you can omit this flag if you don't use the wallet):

$ Bitmemed --utxoindex
You can invoke Bitmemed --help to get a list of more running flags.

The first time you run Bitmemed it will retrieve peer information from BitMEME's DNS server and will start synchronizing with the network. First synchronization may take up to several hours (depending on your CPU strength and bandwidth). It is impossible to mine before the network is synced. Every time you run Bitmemed it will incrementally sync any blocks accumulated while it was offline, this is typically a much shorter process (as long as Bitmemed was not shut down for more than several hours).

Creating a Wallet
To run a miner you need to create a keypair to mine into:

$ bitmemewallet create
You will be asked to choose a password for the wallet (a password must be at least 8 characters long, and it won't be shown on the screen you as you entering it). After that you should run this command in order to start the wallet daemon:

$ bitmemewallet start-daemon
And then run this in order to request an address from the wallet:

$ bitmemewallet new-address
Your screen will show you something like this:

The wallet address is:
bitmeme:0123456789abcdef0123456789abcdef0123456789
Note: Every time you ask bitmemewallet for an address you will get a different address. This is perfectly fine. Every secret key is associated with many different public addresses and there is no reason not to use a fresh one for each transaction.

At this point your can close the wallet daemon, though you should keep it running of you want to be able to check your balance and make transactions

CPU Miner (optional)
Note: Our miner was highly superceded by Elichai's miner (see below), we recommend that you use that miner instead.

After having created a wallet, copy the address and run bitmememiner with it:

$ bitmememiner --miningaddr bitmeme:<YOUR_CREATED_ADDRESS>
Note: The miner is single threaded, so it is best to run several instances of it to utilize more than one CPU core. Note: Mining cannot start before the network is syncrhonized. In order to conserve your CPU, the miner will not start mining before the node is synced. Hence, it is expected to see a mining rate of 0 Hashes/second for a while as bitmemed obtains the current network state.

GPU Miner
A community developed GPU miner with CUDA support is available here: 
The GPU miner now supports openCL and AMD GPUs as well.

Mining on Additional Computers
Not all machines need to run bitmemed. Once you have a running node, any other machine can report their blocks to it by using the -s flag:

$ bitmememiner -s <node IP address> --miningaddr bitmeme:<YOUR_CREATED_ADDRESS>
You can run ifconfig in Linux or Mac or ipconfig in Windows on the machine running bitmemed to find out its IP address.

Opening Ports
By forwarding port 36111 (unless configured otherwise) to the machine running bitmemed, your node becomes a public node which other members of the network can use to sync. Even though private nodes can still mine, it is encouraged that you make your node public for the general health of the network. Like any other decentralized systems, BitMEME works best when there are many public nodes.

bitmemed Hardware Requirements
Minimum:
20 GB disk space
2-core processor or AMD equivalent
4GB memory
1 Mbit internet connection


You may want to join our Social media servers for further questions: 
Discord: https://discord.gg/bitmemebtm
Telegram: https://t.me/bitmemebtm
Twitter: https://twitter.com/bitmemebtm
Website: https://bitmeme.world/

License
bitmemed is licensed under the copyfree ISC License.