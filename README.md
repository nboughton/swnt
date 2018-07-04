# swnt - Command line tools for Stars Without Number GMs

swnt provides a command line interface to the roll tables and generators found in [Stars Without Number : Revised Edition (free version)](https://www.drivethrurpg.com/product/230009/Stars-Without-Number-Revised-Edition-Free-Version).
While the code is MIT licensed, all roll table content (except the name.System table) is the copyright of [Kevin Crawford, Sine Nominee Publishing](https://sinenominepublishing.com/).

## Installation

You will need the [go compiler](https://golang.org/) to install swnt. Once installed just run

    go get github.com/nboughton/swnt
    go install github.com/nboughton/swnt

swnt should build on other platforms but I'm not able to test them so I can't guarantee it'll work as expected on anything other than linux. At some point I'll provide
a binary download and probably an AUR package for Arch Linux.

## Usage

At present all commands are for generating content. Most useful is likely to be the Sector generator. Generating a new sector is done by issuing the following command:

    swnt new sector

You can optionally add the -l flag if you want cell text to be coloured according to Tech Level. When you choose to write a sector a new folder will be created and each system will be provided with its own folder. This is designed to make it easy to correlate notes on a per system basis and keep things reasonably well organised. Each folder will contain a text file with the primary world's rolled information (biosphere, temperature, techlevel... etc.). There is also a 10% possiblity of a secondary world and 30% possibility of a point of interest.

Make sure you use a monospace font in your terminal otherwise the output won't line up properly.

![A generated sector](screenshot.png "A generated sector")

All commands can be queried for their available options with the -h flag. For example:

    swnt new -h
    Generate content

    Usage:
      swnt new [command]

    Available Commands:
      adventure   Generate an Adventure
      alien       Generate an Alien
      conflict    Generate a Conflict
      corporation Generate a Corporation
      culture     Generate a Culture
      encounter   Generate a quick encounter
      heresy      Generate a Heresy
      npc         Generate a NPC
      place       Generate a place
      poi         Generate a Point of Interest
      religion    Generate a Religion
      sector      Create the skeleton of a Sector
      world       Generate a secondary World for a Sector cell

    Flags:
      -h, --help   help for new

    Use "swnt new [command] --help" for more information about a command.

This is currently an alpha as there are a number of additions and changes to be made.

## FAQ

### Why not make a web app?

Because [Sectors Without Number](https://sectorswithoutnumber.com/) already exists and is an awesome tool. Also, I live on the CLI. It's where I work, where I'm comfortable and I suspect I'm not alone in that. swnt lets me just hammer a few keys and instantly get something back without having to faff around with GUIs and that's how I like it. I also like using screenshots of the player maps as my actual maps in Roll20. It's got that awful 80s retro-scifi look that just burns the retinas in all the right ways.
