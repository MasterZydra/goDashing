[![Go](https://github.com/Hein-Software-Solutions/goDashing/actions/workflows/go.yml/badge.svg)](https://github.com/Hein-Software-Solutions/goDashing/actions/workflows/go.yml)
# GoDashing

**Contentlist**  
- [Common](#common)
	- [Key features](#key-features)
- [Setting up project](#setting-up-project)

# Common
goDashing is [Golang](http://golang.org) based framework that lets you build beautiful dashboards.
This project is a "fork" of the original project [shopify/dashing](http://shopify.github.io/dashing) and [gigablah/dashing-go](https://github.com/gigablah/dashing-go). This dashing project was created at Shopify for displaying custom dashboards on TVs around the office.

![example dashbaord](./docs/screenshot.png)

## Key features
- Works out of the box, no server, runtime or dependency requiered.
- Use premade widgets, or fully create your own with css, html, and js.
- Pull Data from JIRA to your dashboard with a html attribute.
- Schedule and execute any script/binary file to feed data to your dashboard.
- Use the API to push data to your dashboards.


# Setting up project
- Download all dependencies: `go mod vendor`

# Build the project
- `packr build -o ./goDashing ./cmd/godashing/...`

--------------------------------------

# Create a new dashboard
create a name_here.gerb file in the ```dashboards``` folder

* every 20s, goDashing will switch to each dashboard it founds in this folder.
* you can group your dashboard in a folder.
	* example : ```dashboards/subfolder/dashboard1.gerb```  will be available to http://127.0.0.1:8080/subfolder/dashboard1. 
	* doDash will auto switch dashboards it founds in the sub folder.

## Customize layout
* modify ```dashboards/layout.gerb```
	* if you add a layout.gerb in a dashboards/subfolder it will be used by goDashing when displaying a subfolder's dashboard.


# Feed data to your dashboard

## jobs folder usage
When you place a file in ```jobs``` folder Then goDashing will immediatly execute and schedule it according to this convention : ```NUMBEROFSECONDS_WIDGETID.ext```
* the filename has 2 parts :
	* NUMBEROFSECONDS,  interval in seconds for each execution of this file.
	* WIDGETID, the ID of the widget on your dashboard.

The output of the executed file should be a json representing the data to send to your widget, see examples in ```jobs``` folder.

2 cli arguments are provided to each executed file
1. The url of the current running goDashing
2. the token of the current running goDashing API
3. 
You can use this if you want to send data to multiple widgets. (see example)

### PHP files
When you add a php file to the ```job``` folder, goDashing will  assume ```php``` is available on your system, and will run it with it.

## HTTP call usage (dashing API)
```
curl -d '{ "auth_token": "YOUR_AUTH_TOKEN", "text": "Hey, Look what I can do!" } http://127.0.0.1:8080/widgets/YOUR_WIDGET_ID
```


## JIRA Jql and filters
Edit your .gerb dashboard to add jira attributes to your widget :

* ```jira-count-filter='17531'``` - goDashing will search jiras with this filter and feed the widget with issues count.
* ```jira-count-jql='resolution is EMPTY'``` - goDashing will search jiras with this JQL and feed the widget with issues count.
* ```jira-warning-over='10'``` - widget status will pass to warning when there is more dans 10 issues
* ```jira-danger-over='20'``` - widget status will pass to danger when there is more dans 20 issues

You don't need to restart goDashing when editing gerb files to take changes into account.

### jira configuration
create a ```conf/jiraissuecount.ini``` file in goDashing working directory.
* set url, username, password, interval in the file, 

```
url = "https://jira.atlassian.com/"
username = "" #if empty jira will be used anonymously
password =  ""
interval = 30
```


# Use your custom assets, widgets...
* goDashing looks for assets in a ```public``` folder, when it can not found a file in this folder, it will use its embeded one.

## Widgets
To add a custom widget "Test"
* create a ```widgets``` folder in working directory
	* create a ```Test``` folder
		* add the ```Test.js```, ```Test.html```, ```Test.css``` files in it.

goDashing will use them as soon as you set a widget with a ```data-view="Test"```

Be sure to look at the [list of third party widgets][4].


### COFFEESCRIPT ? SCSS ? JS ? CSS ?
* convert coffeescript to js : http://js2.coffee
* convert scss to css : http://www.sassmeister.com


[2]: 
[3]: 
[4]: https://github.com/Shopify/dashing/wiki/Additional-Widgets
