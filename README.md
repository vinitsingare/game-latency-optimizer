🚀 Game Latency Optimizer (UDP Overlay Network)

A Go-based networking project that measures, compares, and optimizes network latency using a UDP relay (overlay network) deployed on a VPS.
Inspired by tools like ExitLag, this project focuses on measurement-driven routing, not marketing claims.

📌 Problem Statement

Online games rely on low latency and stable network paths.
However, ISPs may route traffic through suboptimal or congested paths, increasing latency and jitter.

This project explores:

How latency behaves over different routes

Whether routing traffic through an intermediate relay (VPS) can improve performance

How to make routing decisions based on real-time RTT measurements

🧠 Core Idea

We compare two network paths:

Direct Path
Client → ISP → Destination Server

Relay (Overlay) Path
Client → ISP → VPS Relay → Destination Server

By measuring RTT (Round-Trip Time) for both paths, the system can determine which route is better under current network conditions.

🏗️ Architecture Overview
+-------------------+
|   Client (PC)     |
|-------------------|
| - RTT Probing     |
| - Health Checks   |
| - Route Decision  |
+---------+---------+
          |
          | UDP
          v
+-------------------+        (VPS - Singapore)
|   Relay Server    |-------------------------+
|-------------------|                         |
| - UDP Forwarding  |                         |
| - Overlay Routing |                         |
+---------+---------+                         |
          |                                   |
          | UDP                               | UDP
          v                                   v
+-------------------+               +-------------------+
| Destination Server|               | UDP Echo Server   |
| (Game / Simulated)|               | (Port 10000)     |
+-------------------+               +-------------------+

🔧 Components
1️⃣ Client

Runs on local machine

Measures:

Direct RTT

Relay RTT

Performs health checks

Acts as control plane

2️⃣ Relay (VPS)

Runs on Ubuntu VPS (Singapore)

Forwards UDP packets

Acts as overlay network node

3️⃣ UDP Echo Server

Simulates a game server

Used for consistent RTT measurement

Replies immediately to UDP packets

📊 Metrics Collected

Direct RTT – baseline ISP routing latency

Relay RTT – latency via overlay path

Health Status – reachable / unreachable relay

Example output:

Direct RTT: 238ms
Relay RTT: 278ms
Relay HEALTHY

🧪 Key Observations

Relay RTT is often higher than Direct RTT due to extra hop

Overlay routing does not guarantee lower latency

Overlay helps only when ISP routing is suboptimal

Measurement-based decisions are critical

This reflects real-world networking behavior.

🛠️ Tech Stack

Language: Go

Networking: UDP sockets

Concurrency: Goroutines

Infrastructure: Ubuntu VPS (Singapore)

Protocols: UDP (no TCP masking)

OS: Windows (client), Linux (VPS)

🚀 How to Run
(Need To Buy VPS)
On VPS
go run udp_echo_server.go      # Port 10000
go run relay/cmd/main.go       # Port 9999

On Local Machine
go run ./client/cmd/main.go

🎯 Learning Outcomes

Deep understanding of UDP networking

NAT & mobile network behavior

RTT vs one-way latency

Overlay routing limitations

Real infra deployment (VPS)

System design thinking

⚠️ Important Note

This project does not claim to always reduce ping.
Instead, it demonstrates how routing decisions should be data-driven, which is how real networking systems work.

📈 Future Improvements

Auto route selection (Direct vs Relay)

Jitter & packet loss measurement

Multiple relay support (Mumbai, Singapore, etc.)

GUI / dashboard

Systemd-based relay service

Game-specific server testing

🧑‍💻 Author

Vinit Singare
Computer Science & Engineering
