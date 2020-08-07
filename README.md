# web-ipt-d
 
web-ipt-d is a very simple web interface to iptables. It runs under systemd control, uses socket activation and reloads the config on SIGHUP.
 
 
 
## Scenario
 
It was written as a proof of concept for a scenario, where one wants to block an ip address on a remote machine.

For example, if one has a couple of containers running the same service, even on multiple worker hosts, one container can detect and block a source ip based on local log file analysis.
Local means either the container or the host running that container.

But:
* How about the other containers running the same service? Synchronize between them?
* What if a container is restarted and the local rules get lost?
 
Another scenario would be the protection of a network based on log file analysis somewhere within the network. The entry point of that network could block ip addresses that were found to send malicious traffic to some service within that network.
 
Both scenarios could profit from blocking unwanted traffic as early as possible.
 
This web service exposes iptables so that unwanted traffic can be blocked from remote.
 
So if for instance [fail2ban](https://www.fail2ban.org) can make use of this service by issuing a block action like this:

`curl --cacert public.pem -i -u user:pass https://server:8820/chain/myChain/ipv4/1.2.3.4 -X PUT`
 
### Flaws:
 
* Unblock action: Each ip is only added once. If multiple containers want to block the same ip, one would need to handle the unblock action/timing in web-ipt-d. As the information on the container could be lost. This is not implemented.
* Currently, the ip is blocked - no port information can be added
* When reloading the configuration, the reading of config parameters during write is not blocked. This data race condition is not solved.
 
## Feature Description
 
The command line parameters are displyed with the -help flag as follows:
`./web-ipt-d -help`

This will print something like this to the console:
`Usage of ./web-ipt-d:
  -config string
    	(absolute) path and name of the configuration file
  -version
    	version of the running code
  -write_config
    	write sample config.json file to local directory`
 
The configuration can be edited and loaded into a running daemon by sending a hangup or interrupt (CTRL-C) to the service.
 
Example: `sudo kill -SIGHUP xxxx` or short `sudo kill -HUP xxxx` with xxxx the process id.
 
## Installation
 
This is also a reminder to me on how I did it.
 
1. Name the files the same so one does not need to refer to each other in the unit files:
    * executable    web-ipt-d
    * service file  web-ipt-d.service
    * socket file   web-ipt-d.socket
 
2. Copy executable according to the .service file (ExecStart). E.g. /usr/sbin/web-ipt-d and chmod it to 0700 for user root
 
3. Copy the configuration file to -config=/etc/web-ipt-d/config.json and chmod it to 0600 for user root
 
4. Find directory for systemd unit files with `pkg-config systemd --variable=systemdsystemunitdir`

    Edit the .service file - see e.g. https://wiki.debian.org/ServiceSandboxing

    Copy the .service file to the systemdsystemunitdir found above
 
5. Edit and copy the .socket file to the same directory

    Note: Here is the server ip and port configured. Perhaps it is a good idea to bind the service to only the internal network.
 
6. Dance the systemctl:
    * `systemctl daemon-reload` and `systemctl start web-ipt-d.socket`

    Here it is the .socket that is started. (Quick check the executable with `systemd-socket-activate -l 8820 ./web-ipt-d`)
 
    * Stop the service with `systemctl stop web-ipt-d.service`

    Here it is the .service that is stopped.
             
              * Usually, both socket and service are started / stopped like here:
              `systemctl start web-ipt-d.socket
              systemctl start web-ipt-d.service`
 
Also note `sudo systemctl status web-ipt-d.socket` or `sudo journalctl -u web-ipt-d.socket` to see the listening socket.

The log of the service (errors only): `sudo journalctl -u web-ipt-d.service`

The log of the daemon (all log levels) is at `/var/log/web-ipt-d.log` A sample lograte file is included.
 