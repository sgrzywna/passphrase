# Passphrase

POC of random password generator implemented in go.

The main reason to create the tool was a need for small handy tool that generates easy to remember, but hard to crack passwords.

It has much common with [Diceware](https://en.wikipedia.org/wiki/Diceware) method.

To make things a little bit more interesting, pseudo random number generator uses `The Simple Discard Method` that I found in `NIST SP800-90` publication.

## Usage

The tool supports few command line switches for customization:

```bash
./passphrase -h

Usage of ./passphrase:
  -d string
        path to dictionary file (default "./dicts/polish.txt")
  -p uint
        number of passwords to generate (default 10)
  -w uint
        number of words in each password (default 3)
```

But it works perfectly with defaults:

```bash
./passphrase

morus zlepek nocny
budrys=dwugaz=knel
rojny-kreol-piorun
akces_elki_wino
kwant=glejt=granat
byczek_prym_paniny
wojsko%szwab%skaut
kluch*apetyt*bufet
rumak dubler ploter
ringo:osada:dira
```

There are Polish (default) and English dictionaries available, to use English one:

```bash
./passphrase -d ./dicts/english.txt

golem%brya%gheg
pareto_tuareg_cachi
quarto&myrtle&crest
gidar.binder.storm
planet+estate+lake
punk!woof!nature
newton%brill%beef
moon-quincy-judah
imide:vogue:hmong
louse%trust%filer
```

## Build

To build the tool:

```bash
make
```

For the Windows executable cross compile:

```bash
make build-windows
```
