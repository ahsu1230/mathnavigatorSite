# Potential Problems

## AWS Setup

- Currently just have one EC2 instance running. We need a auto-scaling thing to always have one EC2 available.
- Need a load balancer to link instead of individual EC2 instances.
- Need to move domain name

## Orion

- I'm using admin user to connect Orion to RDS... That's kind of scary. Need new user groups but can't grant them privileges for some reason.

## Gemini-Admin

- ~~Gemini Admin is currently on `/var/www/html` folder on EC2 instance.~~ Will not be published anymore.
- How do we create a subdomain? Do we need to wait until we move over domain?
- How do we ONLY enable HTTPS and not HTTP?

- Handling images has a bug. For some reason, the images keep mapping to `/dist` and not `/dist/`. Have to manually fix it with a bash command.
- ~~Admin site needs a Login screen!~~ Will only be local run from home network. Will NOT publish to AWS.
- Currently manually setting ORION_HOST_PROD in webpack. Use a real environment variable to hide.

## Gemini-User

- ~~Will place Gemini User on EC2 instance at `/var/www/html`... OR should we place it on Bluehost???~~ Created a docker instance express server to handle this and always serve files. How to make HTTPS?
- How do we ONLY enable HTTPS and not HTTP?
- Currently manually setting ORION_HOST_PROD in webpack. Use a real environment variable to hide.
- Will need a "Maintenance page..."