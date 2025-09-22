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