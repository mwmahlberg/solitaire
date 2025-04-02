solitaire
=========

This is a Go library and a command line utility for Solitaire, an [encryption
algorithm based on a deck of playing cards invented by Bruce
Schneier][schneier:solitaire].

It is highly recommended to read Bruce Schneier's original post before using the
algorithm and by extension this library or the acompanying command line utility
**BEFORE** you put this to any serious application.

Also, please read and understand the [License](./LICENSE).

Table of Contents
-----------------

- [Install](#install)
  - [Library](#library)
  - [Command line utility](#command-line-utility)
- [Usage](#usage)
- [Operational Notes](#operational-notes)

Install
-------

### Library

```shell
go get github.com/mwmahlberg/solitaire
```

### Command line utility


Usage
-----

Operational Notes
------------------

Quoted from Bruce Schneier's original post:

> The first rule of an output-feedback mode stream cipher, any of them, is that
> you should never use the same key to encrypt two different messages. Repeat
> after me: NEVER USE THE SAME KEY TO ENCRYPT TWO DIFFERENT MESSAGES. If you
> do, you completely break the security of the system. Here’s why: if you have
> two ciphertext streams, A+K and B+K, and you subtract one from the other, you
> get (A+K)-(B+K) = A+K-B-K = A-B. That’s two plaintext streams combined with
> each other with no key involved, and is very easy to break. Trust me on this
> one: you might not be able to recover A and B from A-B, but a professional
> cryptanalyst can. This is vitally important: never use the same key to
> encrypt two different messages. Keep your messages short. This algorithm is
> designed to be used with small messages: a couple of thousand characters at
> most. Use shorthand, abbreviations, and slang in your messages. Don’t be
> chatty. If you have to encrypt a 100,000-word novel, use a computer algorithm.
>
> Like all output-feedback stream ciphers, this system has the unfortunate
> feature of never recovering from a mistake. If you’re encrypting a message,
> and you make a mistake in one of the operations, every letter afterwards will
> be encrypted wrong. You won’t be able to decrypt it, even with the key. And
> you’ll never know. So if you’re encrypting a message, go through the
> encryption process twice to make sure they agree. If you’re decrypting, check
> to make sure the message makes sense as you decrypt it. And if you’re keying
> from a random deck, keep a spare copy of the ordered deck for this reason.
>
> Solitaire is reversible. This means that if you leave the deck lying around
> after you’ve encrypted your message, the secret police can find it and work
> the algorithm backwards using the deck. This process can recover all of the
> output cards and decrypt a message. It is important that you shuffle the deck
> completely, six times, after you finish encrypting a message.
>
> For maximum security, try to do everything in your hands and head. If the
> secret police starts breaking down your door, just calmly shuffle the deck.
> (Don’t throw it up in the air; you’d be surprised how much of the deck
> ordering is maintained during a game of 52-Pickup.) Remember to shuffle the
> backup deck, if you have one.
>
> Be careful about worksheets, if you have to write things down. They will have
> sensitive information on them.

[schneier:solitaire]: https://www.schneier.com/academic/solitaire/ "Original blog post"
