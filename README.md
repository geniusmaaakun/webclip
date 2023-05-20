<div style="text-align: center;">

# WebClip

![ソースコード言語](https://img.shields.io/github/languages/top/geniusmaaakun/webclip)

![ソースコード数](https://img.shields.io/github/languages/count/geniusmaaakun/webclip)

![ソースコードサイズ](https://img.shields.io/github/languages/code-size/geniusmaaakun/webclip)


![logo](./frontend/public/favicon.ico)

</div>



## Overview
Get HTML files from a website on the command line, convert them to markdown format, and save them in a specified folder.
By specifying the option, image files are also downloaded.
It is used when you want to save the material of a specific URL locally.


## Requirement
- MacOS, Windows, Linux
- golang: 1.19
- react 17.0.2

## install
```
make install
```

## Usage
### List of command options
* -u: specify the URL
* -o: specify the save directory
* -save: save to database
* -d: also save the image file of the target page

```
// Converts the HTML of the target URL to markdown and saves it locally in the specified directory
webclip -u "{URL}" -o "{OUTPUT_DIR}"

// By adding --save, you can save it to the database and list it with the server command described later.
webclip -u "{URL}" -o "{OUTPUT_DIR}" -save

// It also downloads images in HTML
webclip -u "{URL}" -o "{OUTPUT_DIR}" -save -d
```

###  List of subcommands

* search: Search the database for files that match the conditions from the saved files. Only if --save is specified
```
webclip search -b "{body}"
webclip search -t "{title}"
```

* server: Display the list page of saved files. Simple as it is a command line tool. http://localhost:8080
```
webclip server
```

* clean: Delete data from DB if file path does not exist. You can save DB space.
```
webclip clean
```

* zip: zip the file
```
webclip zip
webclip zip -t "{title}"
webclip zip -b "{body}"
webclip zip -t "{title}" -b "{body}"
```



## Features and Description

- Created to automate manual copying and pasting of materials.
By executing it on the command line, the HTML file of the specified URL is acquired, converted to markdown format, and saved in the specified folder.

- clean architecture \
I made it for clean architecture practice.

- react + golang \
I wanted to try using react and golang, so I adopted react for the front end.


## Reference

- README書き方
* https://qiita.com/legitwhiz/items/bb34ef20ba23336e0c87
* https://qiita.com/Kotabrog/items/fb328b72ac94137897af#requirement

- バッジ
* https://kic-yuuki.hatenablog.com/entry/2019/06/29/173256
* https://shields.io/


## Author

[twitter@geniusmaaakun](https://twitter.com/geniusmaaakun)

## Licence

[MIT](https://opensource.org/license/mit/)
