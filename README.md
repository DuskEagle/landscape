This is my experiment on what an infrastructure as code tool could look
like if it was available as a library. 

Existing tools like Terraform, Chef, and Pulumi run in separate processes.
This adds friction to managing infrastructure within a codebase.
They also operate in a model where you will want to update all of the
resources within a project every time you perform an update. Aside from
Pulumi, they also have their own configuration languages. 

Landscape aims to give you the control to manage your cloud resources
alongside other code in your programming language. Unlike other tools, it
gives you control to update just the cloud resources you wish to update,
when you want to update them..

Examples are available in the `examples/` directory. As I work out how I
would want this tool to work, I will add more examples, and write code
to turn those examples into something that works.

This is more of an exploration project than anything else. It's unlikely
to get far enough to actually use.
