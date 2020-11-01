# Deploying your app to AWS

## Create VPC + RDS

- Create VPC
  - A **Virtual Private Cloud** is your isolated subsection of the AWS cloud. Inside your VPC, entities can discover each other but are kept isolated from other entities of other clouds.
- Create RDS (MySQL)
- Create Security Groups

Security Groups are virtual firewalls that dictate traffic between your instances. Entities of your VPC can select multiple security groups so it's often nice to have security groups de-compositioned to handle a small set of permissions. For instance, I can have one security group for my HOME access and another default security group for only the VPC. That way any entity in my security group can be accessed by other entities of the VPC and can be SSH'd into by me as long as I'm at home.

# Managing EC2 Instances

- What is an AMI?
  - **Amazon Machine Image** is an archetype / blueprint of what the EC2 instance should be (think Linux, Windows, # SSD, #GB, etc.)
- Select Instance Type
  - The tiers of how large an Amazon machine can be (t1, t2, micro, large, etc.)
- Determine if we want to use a launch template?
  - A launch template is often used for auto-scaling groups. Use a launch template to configure how to bring up additional instances when scaling. *NOTE* launch templates only determine the AMI. **Launch configurations** have the additional bonus of allowing user data scripts to be run on initialization.

## Creating Auto-scaling Group

- First decide if you want to use a **launch template** or **launch configuration**.
  - A launch template is for determining the AMI of an instance. A launch configuration allows you to also specify what `user data` or commands to run on instance startup.
- Then create an **Auto-Scaling Group**. 
  - Choose if you want to use a launch template or launch configuration.
- You have the option to use a load balancer, which is explained below. The load balancer will be associated with your auto-scaling group and can be used to determine if groups / instances are healthy. If they are not healthy, we may need to scale up on instances to create healthy instances.

## Creating Load Balancers

- Recommended to use the Application Load Balancer (HTTP / HTTPS).
- Create *Target Groups* which are entities that route requests from your load balancer to your instances. 
  - Determine how to route (HTTP / HTTPS) requests from one port to another. 
  - Configure health check settings for your target groups.
- Select the desired capacity of your scaling group (min, max, desired)
- Make sure to review your Security Group to ensure load balancer is allowed to port network over to the EC2 instance. Are they in the same VPC? Are we safely allowing HTTP / HTTPS requests from the public internet?

## Registering Domain and using Route53

- Go to Route53 and register a domain
- Once domain is registered or fully transferred from previous domain service, double check the domain has a single `NS` Record with 4 namespace servers defined.
- You can then define an additional Record `A` (Simple Route) to route to the load balancer you created in the previous step. I had 3 different `A` Records (Example: `andymathnavigator.com`, `www.andymathnavigator.com` and `*.andymathnavigator.com`).
  - Double check your security group here! Your load balancer will need to accept HTTP / HTTPS / certain TCP requests from anywhere!

*Difference between A vs. AAAA vs. CNAME*

- A records are the most basic form of a DNS record. The A record points to a specific IPv4 address.
- AAAA records point to a specific IPv6 address (similar to an A record).
- CNAME records point a name to another name (NOT an IP address, unlike an A record). So you can think of it as a synonymous "alias" that points a name to another name.

## Amazon Certificate Manager

  - If you configured your domain correctly, you can create a certificate for your domain to allow HTTPS transactions.

## Using CodeDeploy (unused)

- Create IAM role to allow service to interact with your EC2 instance (in this case, CodeDeploy)

## Creating S3 buckets and using Cloudfront
<https://aws.amazon.com/premiumsupport/knowledge-center/cloudfront-access-to-amazon-s3/>
Allow Cloudfront to access images / blobs / objects from S3. Ideally, you want to not allow public access directly to S3 but allow cloudfront to publicly distribute those images.