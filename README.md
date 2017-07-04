LiREST: Linux on REST

[![Go Report Card](https://goreportcard.com/badge/github.com/howardplus/lirest)](https://goreportcard.com/report/github.com/howardplus/lirest)

LiREST gives you control of the system remotely by:

* Creates REST API based on description files
* Parse and produce JSON formatted output
* Support system variables or commands

### Quick Start

LiREST runs as a REST server that consumes system settings or commands in the background. Originally designed to read Linux /proc and /sys filesystems, LiREST can be used to expose arbitrary system commands and expose them as REST api based on a set of description files.

    LiREST exposes Linux operating system using REST API

    Usage:
      LiREST [flags]
      LiREST [command]

    Available Commands:
      client      LiRest client commands
      help        Help about any command
      version     Print the version number of LiREST

    Flags:
      -d, --desc-path string   Description file path (default "./descriptions/")
      -u, --desc-url string    Download URL for description files
      -h, --help               help for LiREST
      -i, --ip string          IP address to listen on (default "localhost")
      -s, --no-sysctl          Disable sysctl routes
      -p, --port string        Port to listen on (default "8080")
      -y, --pretty             Pretty-print JSON output
      -q, --quiet              quiet output
      -v, --verbose            verbose output
      -w, --watch              Watch for changes in description files

    Use "LiREST [command] --help" for more information about a command.

By default, LiREST looks for description files in the **./descriptions/** directory. A description file contains **des** extension, and is a JSON file that describes how to retrieve data from the system, and how to format them to produce JSON response. Sample description files can be found under **/descriptions/** directory.

To start, run LiREST server:

    # lirest 

This starts the LiREST server on localhost:8080. Open a browser on the same machine and type in "http://localhost:8080". It returns a **path** response:

    {"subpath":["proc","sys","sysctl","command"]}

In LiREST, each REST API is a **path**, **resource** or **info**. A **path** lists out all the sub-paths available in the current path, as shown in this example. Follow the subpath to reach a specific resource, such as:

    http://localhost:8080/proc/version
    
    {"data":{"full":"Linux version 3.13.0-85-generic (buildd@lgw01-32) (gcc version 4.8.2 (Ubuntu 4.8.2-19ubuntu1) ) #129-Ubuntu SMP Thu Mar 17 20:50:15 UTC 2016","gcc-version":"gcc version 4.8.2 (Ubuntu 4.8.2-19ubuntu1)","kernel-release":"3.13.0-85-generic","kernel-version":"#129-Ubuntu SMP Thu Mar 17 20:50:15 UTC 2016","os-type":"Linux","username":"buildd@lgw01-32"},"name":"version","time":"2017-07-04T23:15:24.791767734Z"}
    
This is an example of a **resource** defined in "/descriptions/proc.des". The data is retrieved from "/proc/version" filesystem and presented as JSON format. To pretty-print the JSON output, use the "-y" option:

    # lirest -y
    
The result will be formatted properly:

    http://localhost:8080/proc/version
    
    {
    "data": {
      "full": "Linux version 3.13.0-85-generic (buildd@lgw01-32) (gcc version 4.8.2 (Ubuntu 4.8.2-19ubuntu1) ) #129-Ubuntu SMP Thu Mar 17 20:50:15 UTC 2016",
      "gcc-version": "gcc version 4.8.2 (Ubuntu 4.8.2-19ubuntu1)",
      "kernel-release": "3.13.0-85-generic",
      "kernel-version": "#129-Ubuntu SMP Thu Mar 17 20:50:15 UTC 2016",
      "os-type": "Linux",
      "username": "buildd@lgw01-32"
    },
    "name": "version",
    "time": "2017-07-04T23:18:06.248795412Z"
    }

Each resource contains an **info**, essentially documents its usage. To access **info** of a particular resource, append it with "/_info":

    http://localhost:8080/proc/version/_info

    {
    "api": [
      {
        "method": "GET",
        "short": "Get kernel version",
        "long": "This file specifies the version of the Linux kernel, the version of gcc used to compile the kernel, and the time of kernel compilation"
      }
    ],
    "path": "/proc/version"
    }
