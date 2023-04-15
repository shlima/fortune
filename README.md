# Fortune

Cracker of bitcoin addresses (private keys) by brute forcing 
(includes dataset of **323,156** wallets with non-zero balance).

![btc cracker telegram screenshot](/docs/telegram.png?raw=true)

Notifies you in Telegram about the generation process and 
found (guessed) keys.

### Docker
```bash
docker run ghcr.io/shlima/fortune \
--workers 2
--heartbit-sec 3600 
--telegram-ping-sec 21600
--telegram-token ""
--telegram-channel "" 
bruteforce   
```

## Arguments
| Argument    | Description                                                                   |
|-------------|-------------------------------------------------------------------------------|
| file | file with a custom dictionary (each address on a new line) can take many arguments |
| heartbit-sec | print status each N seconds to STDOUT                                         |
| telegram-ping-sec | send status each N seconds to telegram                                        |
| telegram-token | token of the telegram bot                                                     |
| telegram-channel | @channel name for the notifications (bot should be added as an administrator) |

## Examination

In order to check the correctness of the notifications to Telegram, 
as well as the correctness of the dictionaries with addresses, 
run the `bruteforce test` command with the number of the wallet 
in the dataset as an argument:

```bash
fortune bruteforce test 1LQoWist8KkaUXSPKZHNvEyfrEkPHzSsCd
```
