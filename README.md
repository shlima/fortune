# Fortune

[![Test](https://github.com/shlima/fortune/actions/workflows/test.yml/badge.svg)](https://github.com/shlima/fortune/actions/workflows/test.yml)

Cracker of bitcoin addresses (p2pkh private keys) by brute forcing 
and brain forcing (includes dataset of **323,156** wallets with non-zero balance).

Notifies you in Telegram about the process and
found (guessed) keys.

![btc cracker telegram screenshot](/docs/screenshot.webp?raw=true)

### Docker
```bash
docker run ghcr.io/shlima/fortune \
    --workers 2 \
    --heartbit-sec 3600 \
    --telegram-ping-sec 21600 \
    --telegram-token "botfather token" \
    --telegram-channel "@channel" \
    bruteforce   
```

## Arguments
| Argument          | Description                                                                               |
|-------------------|-------------------------------------------------------------------------------------------|
| file              | file with a custom dictionary (each public address on a new line) can take many arguments |
| workers           | number of workers for parallel execution                                                  |
| night             | night or silent mode (reduced CPU usage)                                                  |
| heartbit-sec      | print status each N seconds to STDOUT                                                     |
| telegram-ping-sec | send status each N seconds to telegram                                                    |
| telegram-token    | token of the telegram bot                                                                 |
| telegram-channel  | @channel name for the notifications (bot should be added as an administrator)             |

## Command line
```
NAME:
   fortune - bitcoin wallet cracker

USAGE:
   fortune [global options] command [command options] [arguments...]

VERSION:
   0.1

COMMANDS:
   bruteforce  run bruteforce against the dataset of rich addresses
   random      prints random address from the dataset files
   generate    random bitcoin address
   brain       generate brain wallet base on password first argument)
   brainforce  run bruteforce with alphabetical passwords permutations against the dataset of rich addresses
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --file value [ --file value ]  a file with a custom dictionary (default: "addresses/Bitcoin/2023/04/p2pkh_Rich_Max_1.txt", "addresses/Bitcoin/2023/04/p2pkh_Rich_Max_10.txt", "addresses/Bitcoin/2023/04/p2pkh_Rich_Max_100.txt", "addresses/Bitcoin/2023/04/p2pkh_Rich_Max_1000.txt", "addresses/Bitcoin/2023/04/p2pkh_Rich_Max_10000.txt", "addresses/Bitcoin/2023/04/p2pkh_Rich_Max_100000.txt") [$FILE]
   --workers value                number of workers for parallel execution (default: 1) [$WORKERS]
   --night                        night or silent mode (reduced CPU usage) (default: false) [$NIGHT]
   --test-address value           address to test dataset before running brute force (default: "1LQoWist8KkaUXSPKZHNvEyfrEkPHzSsCd") [$TEST_ADDRESS]
   --heartbit-sec value           print status each N seconds to STDOUT (default: 10) [$HEARTBEAT_SEC]
   --telegram-ping-sec value      send status each N seconds to telegram (default: 3600) [$TELEGRAM_PING_SEC]
   --telegram-token value         token of the telegram bot [$TELEGRAM_TOKEN]
   --telegram-channel value       @channel name for the notifications (bot should be added as an administrator) [$TELEGRAM_CHANNEL]
   --help, -h                     show help
   --version, -v                  print the version

COPYRIGHT:
   © github.com/shlima/fortune
```

## Brain wallets

To find a key from a rich bitcoin brain wallet using brute force method, 
you can use the command `brainforce`

```bash
fortune brainforce 
  --pass-length 5  # password length 
  --pass-alphabet english-lower # one of "digits", "symbols", "english-lower", 
                                # or "english-upper" (can take many arguments)
  --pass-alphabet abc # any characters without separator
  --pass-state "" # the end state from the previous run 
                  # (to continue instead of starting all over again)                
```

## Examination

In order to check the correctness of the notifications,
as well as the correctness of the dictionaries with addresses,
run the `bruteforce` command with the number of the wallet
in the dataset as an argument:

```bash
fortune bruteforce 1LQoWist8KkaUXSPKZHNvEyfrEkPHzSsCd
```

To test a notification of a successful found brain wallet:

```bash
# 1) generate a brain bitcoin address with password "foo"
fortune brain foo

# 1) copy one of public address of the brain wallet and 
#    mock the index by adding this address as the first command argument
fortune brainforce --pass-length 3 --pass-alphabet fo 1LEH8BEZgC4onZ4GLm8UpZ3vXGAr6LYKST
```
