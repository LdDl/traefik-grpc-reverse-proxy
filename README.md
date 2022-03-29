## Just an example how to make reverse proxy for gRPC

### Backstory
  Some time ago I had a task to do reverse proxy for multiple gRPC services.

  I've already known how to do common HTTP/HTTPS/WS/WSS stuff with bith [Traefik](https://traefik.io/) and [Ngnix](https://nginx.org), but I had not have familiar with HTTP2 (which is gRPC based on) reverse proxy.
  I've start digging into and now I provide some workaround below.
  

### Requirements:

* Installed Traefik - https://doc.traefik.io/traefik/getting-started/install-traefik/
  
  In my case I've used binary distribution for linux_amd64.tar.gz with [v2.6.2 tag](https://github.com/traefik/traefik/releases/tag/v2.6.2). Your mileage may vary.

* Prepared configuration file.

  I prefer to have *.toml format. I'll describe its contents later.

* Clone this repo and navigate to main directory
  ```shell
  git clone https://github.com/LdDl/traefik-grpc-reverse-proxy.git
  cd traefik-grpc-reverse-proxy
  ```

### Run

* Run first service
   
  ```shell
  go run service_one/cmd/main.go
  ```

  <details>
  <summary>If you want to rebuild *.pb.go from *.proto file(-s) click to expand [Optional]</summary>

  * Be sure that you have protoc with Go compatibility
    
  * Run protoc
    ```shell
    protoc -I service_one/rpc/proto service_one/rpc/proto/*.proto --go_out=plugins=grpc:service_one/rpc/proto/
    ```

  </details>


* Run second service
   
  ```shell
  go run service_one/cmd/main.go
  ```

  <details>
  <summary>If you want to rebuild *.pb.go from *.proto file(-s) click to expand [Optional]</summary>

  * Be sure that you have protoc with Go compatibility
    
  * Run protoc
    ```shell
    protoc -I service_two/rpc/proto service_two/rpc/proto/*.proto --go_out=plugins=grpc:service_two/rpc/proto/
    ```

  </details>

* Start Traefik
  
  ```shell
  ./traefik --configFile=./grpc_proxy_example.toml
  ```

* Start test client
  
  Note: make sure that *.proto files in client folder match same files in services folders (and rebuild it if necessary)
  <details>
  <summary>If you want to rebuild client protobuffs [Optional]</summary>

  ```shell
  protoc -I client/proto_one client/proto_one/*.proto --go_out=plugins=grpc:client/proto_one/
  protoc -I client/proto_two client/proto_two/*.proto --go_out=plugins=grpc:client/proto_two/
  ```

  </details>

  Run:
  ```shell
  go run client/cmd/main.go
  ```
  If everything goes fine then you'll see:
  ```shell
  Response from service #1 is: code:200  text:"Message from service #1"
  Response from service #2 is: code:200  text:"Message from service #2"
  Response from service #1 (reverse proxy) is: code:200  text:"Message from service #1"
  Response from service #2 (reverse proxy) is: code:200  text:"Message from service #2
  ```

* That's all. Now you know how to do gRPC reverse proxy via Traefik