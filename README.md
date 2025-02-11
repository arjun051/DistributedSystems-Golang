# 🚀 Peer-to-Peer File Storage System

A decentralized peer-to-peer (P2P) file storage system implemented in **Go**. This project enables multiple nodes to connect, store, and retrieve files over a distributed network.

---

## 📌 Features
- ✅ **Decentralized Storage**: Store and retrieve files across multiple nodes.
- 🔗 **P2P Communication**: Nodes connect to each other using TCP.
- 🔄 **File Replication**: If a node doesn't have a file, it fetches from the network.
- ⚡ **Auto-Reconnection**: Nodes attempt to reconnect to peers in case of failure.
- 🔒 **Secure Storage**: Files are stored using Content-Addressable Storage (CAS).

---

## 📁 Project Structure

```
📦 project-root
├── 📂 p2p                # Peer-to-peer networking logic
├── 📂 storage            # File storage and retrieval logic
├── 📜 main.go            # Entry point of the program
├── 📜 README.md          # Project documentation
└── 📜 go.mod             # Go module dependencies
```

---

## 🔧 Installation & Setup

### 1️⃣ Clone the Repository
```sh
git clone https://github.com/arjun051/foreverstore.git
cd foreverstore
```

### 2️⃣ Install Dependencies
```sh
go mod tidy
```

### 3️⃣ Run the File Servers
To start the peer nodes:
```sh
go run main.go
```

---

## 🚀 Usage

### Start the P2P Network
In **`main.go`**, nodes are initialized with unique ports:
```go
s1 := makeServer(":3000")
s2 := makeServer(":4000")
s3 := makeServer(":5000", ":3000", ":4000")
```
To run multiple nodes, start them sequentially:
```sh
go run main.go # Starts server on :3000
go run main.go # Starts server on :4000
go run main.go # Starts server on :5000
```

### Store a File
Files are stored in a distributed manner:
```go
data := bytes.NewReader([]byte("my big data file here!"))
err := s3.Store("picture_1.png", data)
```

### Retrieve a File
To fetch a stored file:
```go
r, err := s3.Get("picture_1.png")
b, err := io.ReadAll(r)
fmt.Println("Retrieved file content:", string(b))
```

---

## 🛠 Troubleshooting

### ❌ `Could not connect to :4000 : dial tcp :4000: connect: connection refused`
**Solution:** Ensure all servers are running and increase the startup delay:
```go
time.Sleep(5 * time.Second)  // Give 4000 more time to start!
```

### ❌ `bind: address already in use`
**Solution:** Kill any process using the port:
```sh
lsof -i :4000   # Find process using port
kill -9 <PID>   # Kill process
```

---

## 📜 License
This project is open-source and available under the [MIT License](LICENSE).

---

## 💡 Future Improvements
- 🌍 Implement **DHT (Distributed Hash Table)** for file lookup.
- 📡 Add **gRPC** support for faster communication.
- 🛡 Introduce **encryption** for secure file storage.

---

### ⭐ Like the project? Give it a star on GitHub! ⭐

