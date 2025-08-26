# Gedis

A lightweight **Redis client CLI** written in Go.  
Supports both **RESP2** and **RESP3** protocols, including parsing of nested arrays, maps, doubles, booleans, and bulk strings.  
fully-custom implementation of `redis-cli`.

---

## Features

- **RESP2 Support**
  - Simple strings (`+`)
  - Errors (`-`)
  - Integers (`:`)
  - Bulk strings (`$`)
  - Arrays (`*`)

- **RESP3 Support**
  - Maps (`%`)
  - Sets (`~`)
  - Doubles (`,`)
  - Booleans (`#`)
  - Nulls (`_`)
  - Verbatim strings (`=`)
  - Nested structures (arrays inside maps, maps inside arrays, etc.)
  - Attribute responses (`|`)

- **Command-Line Client**
  - Interactive mode (`./gedis -h 127.0.0.1 -p 6379`)
  - Send any Redis command (`GET`, `SET`, `HELLO`, etc.)
  - Pretty-prints nested RESP2/RESP3 replies

---

## Installation

Clone the repository:

```bash
git clone https://github.com/0xEbrahim/Gedis.git
```
Enter the directory
```bash
cd Gedis
```
Build
```bash
go build -o gedis ./
```
Run
```bash
./gedis -h <host> -p <port>
```
 - note: -h & -p are optional, if not provided the application will run on the localhost

---
## Example
```
127.0.0.1:6379> SET name "mahmoud"
OK
127.0.0.1:6379> GET name
"mahmoud"
127.0.0.1:6379> EXISTS name
1
```

# Feature Support Comparison: gedis vs official redis-cli

| Feature / Behavior                     | Gedis | Official redis-cli |
|---------------------------------------|---------------------------|------------------|
| commands (SET, GET, DEL, EXISTS, ...etc)| ✅                        | ✅               |
| RESP2 protocol support                 | ✅                        | ✅               |
| RESP3 protocol support                 | ✅ (partial / nested parsing) | ✅ (full)        |
| Double values (`,`)                    | ✅                        | ✅               |
| Boolean values (`#`)                   | ✅                        | ✅               |
| Maps (`%`)                             | ✅                        | ✅               |
| Sets (`~`)                             | ❌                        | ✅               |
| Null (`_`)                             | ✅                        | ✅               |
| Verbatim strings (`=`)                 | ✅                        | ✅               |
| Attributes (`\|`)                        | ❌                        | ✅               |
| Nested arrays / maps                    | ✅                        | ✅               |
| Interactive prompt                      | ✅                        | ✅               |
| Command history / autocomplete          | ❌                        | ✅               |
| Authentication (ACL / password)        | ❌                        | ✅               |
| TLS / SSL support                        | ❌                        | ✅               |
| Cluster / Sentinel support              | ❌                        | ✅               |
| Pipelining / scripting support          | ❌                        | ✅               |
| Pretty RESP3 map output                 | ✅ (raw formatting)       | ✅ (numbered & readable) |

