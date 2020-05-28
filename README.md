# toggle perfect

Hobby project to control a 3d printed device.
Find the model on [www.thingiverse.com](https://www.thingiverse.com/thing:4295022)

## Platform
* Raspberry Pi Zero (linxy arm)

## Controlled devices
* waveshare ink (2.7")
* 4 display input keys
* 4 LEDs

![Gopher starring at the case](https://raw.githubusercontent.com/Oppodelldog/toggleperfect/master/gopher_staring_at_the_case.jpg)

### Plugin Applications
Even if the project is an experiment, and the whole design is all the time in progress I introduced  
some nice concepts that allows to adopt this project very easily for your custom use case.

You can either hang in your app like you would expect it in go by implementing some interface [App Interface](internal/apps/apps.go) 

Also it is possible to compile your app(s) separately and just plug them in by providing an ```*.so``` file. 

### Toggle Perfect
I put some samples into the ```internal/apps``` folders, but "timetoggle" aka "toggle perfect" shows  
a working application.  

During this COVID-19 lock-down / home-office period I created this tool to capture my project times.   
I am one of those who like to have a clean working environment, especially on the computer screen.  
So I try to reduce the applications or browser tabs that are open simultaneously.  
This little tool helps me with that.

When I switch my focus to another project I click those buttons until I see the propert ticket ID on the screen.  
The app from then captures the time I am working on that project.
At the end of day I just switch back to main menu and see a table of tickets I worked on and the appropriate times. 
This makes it quite easy to put the times one time a day into the ticket system.    

I really like this app and won't miss it anymore, maybe because I made it, maybe because it really is useful?!  

### Dev UI
![Dev UI Screenshot](https://raw.githubusercontent.com/Oppodelldog/toggleperfect/master/devui.png)

Dev UI enables developing the software without having to upload and build it on the device.  
For more information visit:   
[internal/remote/web/README.md](internal/remote/web/README.md)
