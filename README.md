[![Go](https://github.com/Hein-Software-Solutions/goDashing/actions/workflows/go.yml/badge.svg)](https://github.com/Hein-Software-Solutions/goDashing/actions/workflows/go.yml)
# GoDashing

**Contentlist**  
- [Key features](#key-features)
- [Dependencies](#dependencies)
- [Getting started](#getting-started)
- [Developer informations](#developer-informations)
	- [Setting up the project](#setting-up-the-project)
	- [Build the project](#build-the-project)
	- [Add your own widgets](#add-your-own-widgets)
	- [Helpful converters](#helpful-converters)

**Links**  
- [Source code annotations](./docs/code/README.md)

GoDashing is a [Golang](http://golang.org) based port of the original project [shopify/dashing](http://shopify.github.io/dashing) and [gigablah/dashing-go](https://github.com/gigablah/dashing-go) that lets you build beautiful dashboards. This dashing project was created at Shopify for displaying custom dashboards on TVs around the office.

```
> ./goDashing_darwin_amd64

 ..|'''.|          '||''|.                  '||       ||
.|'     '    ...    ||   ||   ....    ....   || ..   ...  .. ...     ... .
||    .... .|  '|.  ||    || '' .||  ||. '   ||' ||   ||   ||  ||   || ||
'|.    ||  ||   ||  ||    || .|' ||  . '|..  ||  ||   ||   ||  ||    |''
 ''|...'|   '|..|' .||...|'  '|..'|' |'..|' .||. ||. .||. .||. ||.  '||||.
                                                                   .|....'

------------------------------------------------------------
                       Startup
------------------------------------------------------------
2021/07/28 00:03:59 Check for asset folder 'dashboards'
2021/07/28 00:03:59 Extract asset folder 'dashboards'
2021/07/28 00:03:59 Check for asset folder 'jobs'
2021/07/28 00:03:59 Extract asset folder 'jobs'
2021/07/28 00:03:59 Check for asset folder 'public'
2021/07/28 00:03:59 Extract asset folder 'public'
2021/07/28 00:03:59 Check for asset folder 'widgets'
2021/07/28 00:03:59 Extract asset folder 'widgets'

Listen on http://localhost:8080

------------------------------------------------------------
                      Running
------------------------------------------------------------
2021/07/28 00:03:59 JiraJob : can not read config file conf/jiraissuecount.ini
2021/07/28 00:03:59 ExecJob - 1_doughnutchart.php - scheduled every 1s
2021/07/28 00:03:59 ExecJob - 1_linechart.php - scheduled every 1s
2021/07/28 00:03:59 ExecJob - 1_water_main_city.php - scheduled every 1s
2021/07/28 00:03:59 ExecJob - 2_valuation.php - scheduled every 2s
2021/07/28 00:03:59 ExecJob - 3_buzzwords.php - scheduled every 3s
2021/07/28 00:03:59 ExecJob - 4_barchart.php - scheduled every 4s
2021/07/28 00:03:59 ExecJob - 5_convergence.php - scheduled every 5s
2021/07/28 00:03:59 ExecJob - 5_piechart.php - scheduled every 5s
2021/07/28 00:03:59 ExecJob - 5_polarchart.php - scheduled every 5s
2021/07/28 00:03:59 ExecJob - 5_radarchart.php - scheduled every 5s
2021/07/28 00:03:59 ExecJob - 5_sampleTXT.php - scheduled every 5s
2021/07/28 00:03:59 ExecJob - 5_web-update.php - scheduled every 5s
```

![example dashbaord](./docs/screenshot.png)

## Key features
- **Easy to setup**: It works without setting up a webserver. The server is contained in GoDashing itself. The necessary files get extracted on the first start of the program.
- **Premade widgets**: The program already contains a small list of widgets (BarChart, graph, image, meter, PieChart, Sparkline, TwelveHourClock, comments, html, LineChart, myclock, PolarChart, switcher, DoughnutChart, iframe, list, number, RadarChart, text).  
The list can be extended with you own creations using CSS, HTML and JS.

## Dependencies
For running the jobs successfully your system must have `PHP` installed.

## Getting started
1. Get the app here https://github.com/Hein-Software-Solutions/goDashing/releases
2. Start goDashing `$ ./goDashing`
3. Go to http://127.0.0.1:8080

**Note on macOS**  
macOS requires to add the application to the Gatekeeper Approval. This can be done with the terminal:  
`spctl --add /Path/To/Application.app`  
For more Details please visit [OSXDaily.com](https://osxdaily.com/2015/07/15/add-remove-gatekeeper-app-command-line-mac-os-x)

# Developer informations
## Setting up the project
1. Download the source code
2. Download all dependencies:  
`go mod vendor`
3. Install packr for building the project:  
`go get github.com/gobuffalo/packr/packr`

## Build the project
To build the project in the terminal run the command  
`> packr build -o ./goDashing ./cmd/godashing/...`.  
Packr is a package used for including the necessary files into the binary itself.

To build a version for every operating system the script *release* can be executed. The binaries will be saved in the folder *release*.  
`> ./release.sh`

## Add your own widgets
To add a custom widget add the folder with the widget name in the `widgets` folder.
The following file extensions will be delivered by goDashing: `.css`, `.html`, `.js`
goDashing will use them as soon as you set a widget with a ```data-view="<YourWidgetName>"```

**Important hint:** Widget names are *case sensitive*!

A list of third party widgets for the original Shopify dashing is linked [here](https://github.com/Shopify/dashing/wiki/Additional-Widgets).  
Some porting is required. Helpful converters are listed in the section [Helpful converters](#helpful-converters).

If you have successfully ported a widget or created a new widget, please consider to create a pull request to add it to the default widgets.

## Helpful converters
- CoffeeScript to JS: http://js2.coffee
- SCSS to CSS: http://www.sassmeister.com

## Used dependencies
**JavaScript**  
https://github.com/es-shims  
https://jquery.com/download/  
https://github.com/aterrien/jQuery-Knob  
https://github.com/shutterstock/rickshaw

-------------------------------
# TODO
- Pull Data from JIRA to your dashboard with a html attribute.
- Schedule and execute any script/binary file to feed data to your dashboard.
- Use the API to push data to your dashboards.

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

