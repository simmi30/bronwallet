bronwallet
=========

[![Build Status](https://travis-ci.org/brsuite/bronwallet.png?branch=master)](https://travis-ci.org/brsuite/bronwallet)
[![Build status](https://ci.appveyor.com/api/projects/status/88nxvckdj8upqr36/branch/master?svg=true)](https://ci.appveyor.com/project/jrick/bronwallet/branch/master)

bronwallet is a daemon handling brocoin wallet functionality for a
single user.  It acts as both an RPC client to brond and an RPC server
for wallet clients and legacy RPC applications.

Public and private keys are derived using the hierarchical
deterministic format described by
[BIP0032](https://github.com/brocoin/bips/blob/master/bip-0032.mediawiki).
Unencrypted private keys are not supported and are never written to
disk.  bronwallet uses the
`m/44'/<coin type>'/<account>'/<branch>/<address index>`
HD path for all derived addresses, as described by
[BIP0044](https://github.com/brocoin/bips/blob/master/bip-0044.mediawiki).

Due to the sensitive nature of public data in a BIP0032 wallet,
bronwallet provides the option of encrypting not just private keys, but
public data as well.  This is intended to thwart privacy risks where a
wallet file is compromised without exposing all current and future
addresses (public keys) managed by the wallet. While access to this
information would not allow an attacker to spend or steal coins, it
does mean they could track all transactions involving your addresses
and therefore know your exact balance.  In a future release, public data
encryption will extend to transactions as well.

bronwallet is not an SPV client and requires connecting to a local or
remote brond instance for asynchronous blockchain queries and
notifications over websockets.  Full brond installation instructions
can be found [here](https://github.com/brsuite/brond).  An alternative
SPV mode that is compatible with brond and Brocoin Core is planned for
a future release.

Wallet clients can use one of two RPC servers:

  1. A legacy JSON-RPC server mostly compatible with Brocoin Core

     The JSON-RPC server exists to ease the migration of wallet applications
     from Core, but complete compatibility is not guaranteed.  Some portions of
     the API (and especially accounts) have to work differently due to other
     design decisions (mostly due to BIP0044).  However, if you find a
     compatibility issue and feel that it could be reasonably supported, please
     report an issue.  This server is enabled by default.

  2. An experimental gRPC server

     The gRPC server uses a new API built for bronwallet, but the API is not
     stabilized and the server is feature gated behind a config option
     (`--experimentalrpclisten`).  If you don't mind applications breaking due
     to API changes, don't want to deal with issues of the legacy API, or need
     notifications for changes to the wallet, this is the RPC server to use.
     The gRPC server is documented [here](./rpc/documentation/README.md).

## Requirements

[Go](http://golang.org) 1.12 or newer.

## Installation and updating

### Windows - MSIs Available

Install the latest MSIs available here:

https://github.com/brsuite/brond/releases

https://github.com/brsuite/bronwallet/releases

### Windows/Linux/BSD/POSIX - Build from source

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Ensure Go was installed properly and is a supported version:

```bash
$ go version
$ go env GOROOT GOPATH
```

NOTE: The `GOROOT` and `GOPATH` above must not be the same path.  It is
recommended that `GOPATH` is set to a directory in your home directory such as
`~/goprojects` to avoid write permission issues.  It is also recommended to add
`$GOPATH/bin` to your `PATH` at this point.

- Run the following commands to obtain bronwallet, all dependencies, and install it:

```bash
$ cd $GOPATH/src/github.com/brsuite/bronwallet
$ GO111MODULE=on go install -v . ./cmd/...
```

- bronwallet (and utilities) will now be installed in ```$GOPATH/bin```.  If you did
  not already add the bin directory to your system path during Go installation,
  we recommend you do so now.

## Updating

#### Windows

Install a newer MSI

#### Linux/BSD/MacOSX/POSIX - Build from Source

- Run the following commands to update brond, all dependencies, and install it:

```bash
$ cd $GOPATH/src/github.com/brsuite/bronwallet
$ git pull
$ GO111MODULE=on go install -v . ./cmd/...
```

## Getting Started

The following instructions detail how to get started with bronwallet connecting
to a localhost brond.  Commands should be run in `cmd.exe` or PowerShell on
Windows, or any terminal emulator on *nix.

- Run the following command to start brond:

```
brond -u rpcuser -P rpcpass
```

- Run the following command to create a wallet:

```
bronwallet -u rpcuser -P rpcpass --create
```

- Run the following command to start bronwallet:

```
bronwallet -u rpcuser -P rpcpass
```

If everything appears to be working, it is recommended at this point to
copy the sample brond and bronwallet configurations and update with your
RPC username and password.

PowerShell (Installed from MSI):
```
PS> cp "$env:ProgramFiles\brond Suite\brond\sample-brond.conf" $env:LOCALAPPDATA\brond\brond.conf
PS> cp "$env:ProgramFiles\brond Suite\bronwallet\sample-bronwallet.conf" $env:LOCALAPPDATA\bronwallet\bronwallet.conf
PS> $editor $env:LOCALAPPDATA\brond\brond.conf
PS> $editor $env:LOCALAPPDATA\bronwallet\bronwallet.conf
```

PowerShell (Installed from source):
```
PS> cp $env:GOPATH\src\github.com\brsuite\brond\sample-brond.conf $env:LOCALAPPDATA\brond\brond.conf
PS> cp $env:GOPATH\src\github.com\brsuite\bronwallet\sample-bronwallet.conf $env:LOCALAPPDATA\bronwallet\bronwallet.conf
PS> $editor $env:LOCALAPPDATA\brond\brond.conf
PS> $editor $env:LOCALAPPDATA\bronwallet\bronwallet.conf
```

Linux/BSD/POSIX (Installed from source):
```bash
$ cp $GOPATH/src/github.com/brsuite/brond/sample-brond.conf ~/.brond/brond.conf
$ cp $GOPATH/src/github.com/brsuite/bronwallet/sample-bronwallet.conf ~/.bronwallet/bronwallet.conf
$ $EDITOR ~/.brond/brond.conf
$ $EDITOR ~/.bronwallet/bronwallet.conf
```

## Issue Tracker

The [integrated github issue tracker](https://github.com/brsuite/bronwallet/issues)
is used for this project.

## GPG Verification Key

All official release tags are signed by Conformal so users can ensure the code
has not been tampered with and is coming from the brsuite developers.  To
verify the signature perform the following:

- Download the public key from the Conformal website at
  https://opensource.conformal.com/GIT-GPG-KEY-conformal.txt

- Import the public key into your GPG keyring:
  ```bash
  gpg --import GIT-GPG-KEY-conformal.txt
  ```

- Verify the release tag with the following command where `TAG_NAME` is a
  placeholder for the specific tag:
  ```bash
  git tag -v TAG_NAME
  ```

## License

bronwallet is licensed under the liberal ISC License.
# bronwallet
