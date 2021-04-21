# GO-Port-Listener
Port listener in Go. Allows listening on multiple TCP ports simultaneously.

Typical use: segmentation testing, etc.

Use: 

#> go run listen.go

Responds with HTTP 200, IP addresses, ports, and timestamp:

"2021-04-21 11:27:26 - server: 192.168.0.156:2999 - client: 192.168.0.156:34452"


![image](https://user-images.githubusercontent.com/40667621/115566940-b8866700-a288-11eb-85db-fd1e0c8df664.png)

![image](https://user-images.githubusercontent.com/40667621/115580533-2afd4400-a295-11eb-9ec9-df52d674cde3.png)

![image](https://user-images.githubusercontent.com/40667621/115580460-16b94700-a295-11eb-99ec-857350e769f4.png)



