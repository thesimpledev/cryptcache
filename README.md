# CryptCache

I have used solutions such as AWS Secret Manager in the past. However, for small projects, I often prefer a simple solution—a local file that can securely store secrets and be safely checked into version control and easily run locally. I looked at [SOPS](https://blog.gitguardian.com/a-comprehensive-guide-to-sops/) and while it does a lot of what I wanted, it doesn't do everything. I wanted the ability to have various profiles with unique keys. I wanted CI/CD and SOPS is not a secrets manager. Usually even if it is a personal project I want a production, beta, and local environment available. So I decided to write my own.

[![Go Report Card](https://goreportcard.com/badge/github.com/thesimpledev/cryptcache)](https://goreportcard.com/report/github.com/thesimpledev/cryptcache)

[![Coverage Status](https://coveralls.io/repos/github/thesimpledev/cryptcache/badge.svg?branch=master)](https://coveralls.io/github/thesimpledev/cryptcache?branch=master)

## Security

CryptCache is designed to be reasonably secure. Some may argue that providing attackers with access to encrypted secrets inherently compromises security. However, my counterargument is that if modern encryption standards cannot safeguard shared secrets, the problem extends far beyond this secret manager and reflects fundamental vulnerabilities in current asymmetric encryption standards.

## Core Concept

A vault to manage sensitive information such as API keys and credentials in a way that will allow for local deployment, version control integration, and secure usage during deployment.

## Design Decisions

You may notice a disconnect between how users enter commands and code organization and this is intentional.

It is easier for users to remember a `verb noun` command structure so `create profile` instead of `profile create`, however, when organizing code it makes more sense to organize it as `noun verb` so nouns are at the top level with supporting verbs. While this is a bit more cognitive load in developing I feel it makes the system more user friendly.

## Planned work

- [ ] Store secrets in a plain text TOML file with encrypted values
- [ ] Associate a Private Key with each profile – one key only per profile
- [ ] Allow users to branch off the main profile to create sub-profiles
  - [ ] Sub-profiles will always contain all main profile keys
  - [ ] Sub-profiles can have unique keys that can be added to the main profile
- [ ] Allow Editing individual secrets with public keys
- [ ] Allow Key Rotation updating all secrets to work with the new key
- [ ] CI/CD Pipeline integration
- [ ] Option to transfer encrypted data or unencrypted to deployment
- [ ] HMAC Authentication for secrets and files (with metadata version/timestamp)
- [ ] Integrate with Databases and Key management services

## Planned Commands

- [ ] `cryptcache --init -n "project Name" -p "initial profile name"`
- [ ] `cryptcache --new -p "new profile name"`
- [ ] `cryptcache checkout -p "profile name"`
- [ ] `cryptcache add field "value" --public-key=pub.key`
- [ ] `cryptcache update field "value" --public-key=pub.key`
- [ ] `cryptcache list`
- [ ] `cryptcache get field --private-key=priv.key`
- [ ] `cryptcache export --private-key=priv.key > .env`
- [ ] `cryptcache --verify`
- [ ] `cryptcache --verify-secret fieldname`
- [ ] Export to ENV file
- [ ] Let users define encryption standard

## Example File

```TOML
title = "Crypt Cache"

profiles = ['Production']

['Production.encryption']
private_key_path = 'pkpath'
public_key_path = 'https://helloworld.com'

[Production.username]
value = "admin"
encrypted = false
hmac  = "9df2a160e6c14a15bcd1a9dc2b18e1cf1f003f8ef14ab4cf2c3ef0f0605d7c6b"

[Production.password]
value = "kJlweiur82ljslakdj=="
encrypted = true
hmac  = "7c3f39e9c5df2e65dff9e8d41e1cf2b3ef57a63a12e01452db6ef7df8f7a09c8"

[Production.validation]
hmac = "2d2e3b4f5a6e7c8d9f0e1d3c4b5a6f7d8c9e0f1a2b3c4d5e6f7e0f1a2d3c4b5a"

```
