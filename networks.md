* `host <ip>` get hostname by IP
* `traceroute linkmeup.ru` 
* `tcpdump <ip>`
* use wireshark
![Screenshot 2025-09-17 at 13.40.49.png](Screenshot%202025-09-17%20at%2013.40.49.png)

### DZ
*  `host  bigbarrel.co.nz`
```
bigbarrel.co.nz has address 8.47.69.0
bigbarrel.co.nz has address 8.6.112.0
```

* `ping 8.47.69.0`
* `sudo mtr  bigbarrel.co.nz`
* `sudo tcpdump -i any -v host bigbarrel.co.nz`

### IP
* plugin with diagrams for wireshark
`https://www.cellstream.com/download/a-better-default-profile-with-packet-diagram-for-v3-4-0-and-later/`
* netowrk masks (<ip>/<mask>) like 132.11.0.14/26 and their calculations
![Screenshot 2025-09-22 at 15.05.47.png](Screenshot%202025-09-22%20at%2015.05.47.png)
* subnetworks
![Screenshot 2025-09-22 at 15.05.47.png](Screenshot%202025-09-22%20at%2015.05.47.png)

### Tools
* `ip -4 address show`
* `ip -6 route`

### UDP vs TCP
* TCP has a handshake, UDP does not
* TCP is reliable and slow, UDP is fast and unreliable (in terms of packages delivery)

### Well-know ports
![Screenshot 2025-09-26 at 13.57.40.png](Screenshot%202025-09-26%20at%2013.57.40.png)

### Static routing
![Screenshot 2025-09-29 at 21.52.12.png](Screenshot%202025-09-29%20at%2021.52.12.png)
![Screenshot 2025-09-29 at 22.03.59.png](Screenshot%202025-09-29%20at%2022.03.59.png)