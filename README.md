# AWS PROFILE

If you have to many AWS profiles in your machine, and it's hard to remember the name of every profile. That's a tool that can help you.

AWS Control Tower it's one of the guys that you may have too many profiles. You need to change the profile all the time to access certain account or role.



## How it works

* It reads the file `~/.aws/config` and list to you all profiles. 
* After you choose a profile it'll create a new shell session exporting the env variable `AWS_PROFILE` with the name of profile, eg: `AWS_PROFILE=[ Name of Profile ]`
* After exiting the created shell it will bring back the list
  

# Thank you for

https://charm.sh/ and https://github.com/charmbracelet/bubbletea


