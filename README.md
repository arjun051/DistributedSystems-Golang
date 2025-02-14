# 🚀 Peer-to-Peer File Storage System

A decentralized peer-to-peer (P2P) file storage system implemented in **Go**. This project enables multiple nodes to connect, store, and retrieve files over a distributed network.

---

## 📌 Features
- ✅ **Custom P2P Library**: Built with a fully custom peer-to-peer networking library.
- 🔗 **Decentralized Storage**: Store and retrieve files across multiple nodes.
- 🔄 **File Replication**: If a node doesn't have a file, it fetches from the network.
- 🔒 **SHA-256 Encryption**: Ensures secure file storage and retrieval.
- ⚡ **Auto-Reconnection**: Nodes attempt to reconnect to peers in case of failure.
- 🔍 **Content-Addressable Storage (CAS)**: Files are stored and retrieved using hash-based addressing.

---

## 🔧 Installation & Setup

### 1️⃣ Clone the Repository
```sh
https://github.com/arjun051/DistributedSystems-Golang.git
```

### 2️⃣ Install Dependencies
```sh
go mod tidy
```

### 3️⃣ Run the File Servers
To start the peer nodes:
```sh
make run
```

### 4️⃣ Run Tests
```sh
make test
```
Or manually run:
```sh
go test ./...
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
- 🧩 **Sharding**: Implementing a sharding mechanism for better file distribution.
- 🌍 **Automatic Peer Discovery**: Nodes will automatically detect and connect to new peers.
- 💻 **Web-Based UI**: A user-friendly interface to manage file uploads and retrievals.
- 📡 **gRPC Support**: Faster and more efficient network communication.
- 🛡 **End-to-End Encryption**: Additional security for file storage.

---

### 🙏 Special Thanks
Huge thanks to **Anthony GG** for his invaluable teachings and guidance in building this project!

---

### ⭐ Like the project? Give it a star on GitHub! ⭐

