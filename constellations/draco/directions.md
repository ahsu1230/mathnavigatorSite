# Directions

# VPC + RDS

- Create VPC
- Create RDS (MySQL)
- Create SecurityGroups

# Managing EC2 Instances

## Creating Auto-scaling Group

- First decide if you want to use a **launch template** or **launch configuration**.
A launch template is for
A launch configuration allows you to specify what `user data` or commands to run on instance startup.

- Create a Target Group.
A target group specifies a "group" of instances / lambdas / etc. to auto-scale and check for health. Given a host & port, it will occasionally query that endpoint to check health of the instance.

- Then create an Auto-Scaling Group. 
  - Choose if you want to use a launch template or launch configuration.
  - Choose if you would like to use a load balancer. A load balancer can be created by a Target Group OR by selecting an AWS-created Load Balancer. If you want to do it the latter way, will need to setup Certificate Manager to create Load Balancer.
  - Select the desired capacity of your scaling group (min, max, desired).

## Using CodeDeploy

- Create IAM role to allow service to interact with your EC2 instance (in this case, CodeDeploy)