# Security

This document details some best practices to ensure the security of the host running the `circuit-breaker` service.

# Hosting Provider

* Restrict inbound port access for SSH to your public IP address
  * This part isnt always necessary but usually I like to disable inbound ssh access always, and only enable it when I need to access the host
  * The downside to this is that in the event of emergency maintenance, you will first need to enable inbound ssh access before accessing the host
* 2FA on your hosting provider account

# Operating System

* Ideally use Ubuntu 20.04 LTS, otherwise try to use the latest available LTS version of any debian based linux distribution. 
  * If you cant do that, LTS releases of a redhat operating system work as well.
* Secure SSH by disabling root login, enabling public key authentication, and disabling password authentication
  * Make sure you copy your ssh public key to the server before doing this otherwise you can lock yourself out
* Use a dedicated user account for managing the docker service


## Securing SSH

The config file is usually stored at `/etc/ssh/sshd_config`. Make sure you override any of the previous values for the given settings to the specified values below:

* `PermitRootLogin no`
* `PubkeyAuthentication yes`
* `PasswordAuthentication no`

After updating these values restart the ssh daemon with `sudo systemctl restart sshd`

## Dedicated User Account

* Create an account like so: `sudo useradd -s /bin/bash -m circuit-breaker-svc`
* Change password for the account: `sudo passwd circuit-breaker-svc`
* Grant the user account permissions to manage docker: `sudo usermod -aG docker circuit-breaker-svc`

Now anytime you need to manage the docker services you can follow a pattern like so:

```shell
$> ssh user@host # ssh to the target host
$> su - circuit-breaker-svc # become the circuit-breaker-svc user
$> docker-compose up -d # star the docker services!
```