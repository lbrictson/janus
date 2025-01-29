# Frequently Asked Questions

## Does Janus support a high availability configuration?

Not currently, if you attempt to run more than one Janus server scheduled jobs will run on all servers. Tailing logs of 
running jobs will have unexpected results as well.

## What databases does Janus support?

Janus currently supports SQLite and PostgreSQL.  SQLite is the default and is recommended for small installations.  PostgreSQL
is recommended for larger installations or where protecting a flat file database is not trivial.

## Can I run Janus in a container?

Yes, Janus can be run in a container.  The container is available on Docker Hub as `lbrictson/janus`.  The container is
configured to use SQLite by default, but can be configured to use PostgreSQL as well.

The Janus container comes in both a regular version and a `slim` version.  The slim version only contains bash and the Janus program itself.
The regular version contains many helper tools that admins might find useful.

If you want to build your own Janus container that is customized with your own tools you should start from the `slim` version.

## What tools come preconfigured in the docker container?

The regular version of the Janus container comes with the following tools preinstalled:
- Bash
- Curl
- Wget
- Git
- Ansible
- Python3
- AWSCLI
- pip
- pipx
- ssh
- sshpass
- pipenv

The slim version only contains bash.

## Can I run Janus on Windows or OSX?

You can compile Janus to run on those OS's, however it is only tested on Linux.

## What license is Janus released under?

Janus is released under the MIT license.  You can find the full text of the license in the repository.

## Do you have a helm chart for Janus?

Not yet, soon though.