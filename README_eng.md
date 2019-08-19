![brand](./static/images/logo_sm.png)

MM-Wiki is a light software that enables companies for internal knowledge sharing and better collaboration. It serves as a platform for information sharing and wiki building within as well as among teams.

[![stable](https://img.shields.io/badge/stable-stable-green.svg)](https://github.com/phachon/mm-wiki/) 
[![build](https://img.shields.io/shippable/5444c5ecb904a4b21567b0ff.svg)](https://travis-ci.org/phachon/mm-wiki)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/phachon/mm-wiki/master/LICENSE)
[![platforms](https://img.shields.io/badge/platform-All-yellow.svg?style=flat)]()
[![download_count](https://img.shields.io/github/downloads/phachon/mm-wiki/total.svg?style=plastic)](https://github.com/phachon/mm-wiki/releases) 
[![release](https://img.shields.io/github/release/phachon/mm-wiki.svg?style=flat)](https://github.com/phachon/mm-wiki/releases) 

# Features
- Easy deployment. It’s built with [Go](https://golang.org/doc/). You only need to download the package based on your system and execute the binary file.
- Quick installation. It has a clean and concise installing interface that guides you through the process. 
- Private space for every team or department. By setting permissions, other teams/departments can read, edit files.
- Flexible system administration setting. Each user has different roles with various aspects of permissions accordingly.
- The system allows users to log in with certified external system, such as the company’s LDAP log in method.
- Stay synced with your team. You’ll receive email notifications when the file you're following is updated.
- Share and download the file. For now you can only download file as Markdown plain text.

# Installation
## Install by downloading it.
- Linux
  ```
    # Make a directory. 
    $ mkdir mm_wiki
    $ cd mm_wiki
    # Take linux amd64 as an example: download the latest release of the software.
    # Downloading address: https://github.com/phachon/mm-wiki/releases 
    # Unzip the file to the directory you just created.
    $ tar -zxvf mm-wiki-linux-amd64.tar.gz
    # Enter into the installation directory
    $ cd install
    # Execute the file. The default port is 8090. Set another port using: --port=8087
    $ ./install
    # Visit http://ip:8090 in a browser. Now you should see the installation interface. Follow the instruction to finish settings.
    # Ctrl + C to stop installation. Turn on MM-Wiki. 
    $ cd ..
    $ ./mm-wiki --conf conf/mm-wiki.conf
    # Now you can visit the ip address with the port the system is listening.
    # Enjoy using MM-wiki!
    ```
- Windows
  ```
    # Take linux amd64 as an example: download the latest release of the software.
    # Downloading address: https://github.com/phachon/mm-wiki/releases 
    # Unzip the file to a directory that you set before.
    # Enter into install directory.
    # Double click install.exe. 
    # Visit http://ip:8090 in a browser and now you should see the installation interface. Follow the instruction to finish installations.
    # Close the installation window.
    # Use command line（cmd.exe）to enter into the root directory.
    $ execute mm-wiki.exe --conf conf/mm-wiki.conf
    # Now you can visit the ip address with the port the system is listening.
    # Enjoy using MM-wiki!
    ```

## Install with Nginx reverse proxy
```
upstream frontends {
    server 127.0.0.1:8088; # MM-Wiki listening ip:port
}
server {
    listen      80;
    server_name wiki.intra.xxxxx.com www.wiki.intra.xxxxx.com;
    location / {
        proxy_pass_header Server;
        proxy_set_header Host $http_host;
        proxy_redirect off;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Scheme $scheme;
        proxy_pass http://frontends;
    }
    # static resources managed by nginx
    location /static {
        root        /www/mm-wiki; # MM-Wiki 的根目录
        expires     1d;
        add_header  Cache-Control public;
        access_log  off;
    }
}
```
