# Diffie-Hellman Key Exchange with Caesar Cipher (Go)

This project is a simple Go application that demonstrates the Diffie-Hellman key exchange protocol for generating a shared secret, which is then used as a key for encrypting and decrypting messages using a Caesar cipher.

## Overview

- Implements the Diffie-Hellman key exchange algorithm between two parties (hardcoded private key for one party).
- Uses the shared secret to perform Caesar cipher encryption and decryption.
- Accepts user input to simulate a message-response interaction.
- Decrypts and evaluates the response to return an appropriate encrypted reply.

## Key Concepts Demonstrated

- Modular exponentiation using `math/big` for large number handling
- Basic symmetric key cryptography via Caesar cipher
- User input handling and string manipulation in Go
- Command-line interaction and basic control flow

