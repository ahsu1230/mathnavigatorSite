# Directions

# VPC + RDS

- Create VPC
- Create RDS (MySQL)
- Create SecurityGroups

# Managing EC2 Instances

- Select Instance Type
- What is an AMI?
- Determine if we want to use a launch template?
- Determine if we want to use spot requests?


## Creating Auto-scaling Group

- First decide if you want to use a **launch template** or **launch configuration**.
A launch template is for
A launch configuration allows you to specify what `user data` or commands to run on instance startup.

- Then create an Auto-Scaling Group. 
  - Choose if you want to use a launch template or launch configuration.
  - Choose if you would like to use a load balancer. A load balancer can be created by a Target Group OR by selecting an AWS-created Load Balancer.
  - Select the desired capacity of your scaling group (min, max, desired).
  - Select health check endpoint (try not to use :80/index.html)
  - When instance is healthy, select listeners. Listeners are used to map incoming network (for load balancer) and delegate to a port to an instance. Be careful of whether you're using HTTP / HTTPS / TCP
  - Make sure to review your Security Group to ensure load balancer is allowed to port network over to the EC2 instance. Are they in the same VPC? Is TCP on the listener ports allowed?

## Registering Domain and using Route53

- Go to Route53 and register a domain
- Once domain is registered or fully transferred from previous domain service, double check the domain has a single `NS` Record with 4 namespace servers defined.
- You can then define an additional Record `A` or `AAAA` (Simple Route) to route to the load balanced you created in the previous step.
- Double check your security group here! Your load balancer will need to accept HTTP / HTTPS / certain TCP requests from anywhere!

## Using CodeDeploy (unused)

- Create IAM role to allow service to interact with your EC2 instance (in this case, CodeDeploy)