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
- golang: 1.18~
- react 18

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

- how to write README 
    * https://qiita.com/legitwhiz/items/bb34ef20ba23336e0c87
    * https://qiita.com/Kotabrog/items/fb328b72ac94137897af#requirement

- github shields
    * https://kic-yuuki.hatenablog.com/entry/2019/06/29/173256
    * https://shields.io/

- How to Embed React App into Go Binary
    * https://medium.com/@pavelfokin/how-to-embed-react-app-into-go-binary-12905d5963f0

    * https://www.smartinary.com/blog/how-to-embed-a-react-app-in-a-go-binary/

- sass
    * https://book.scss.jp/code/c2/04.html

- eslint prettier
    * https://zenn.dev/thiragi/scraps/8e988668dbc860

- How to solve "It is not reflected even if you update the value of useState!"
    * https://zenn.dev/syu/articles/3c4aa813b57b8c
    * https://blanktar.jp/blog/2020/06/react-why-state-not-updated
    * https://blanktar.jp/blog/2020/06/react-why-state-not-updated

- React AbortController 
    * https://lightbulbcat.hatenablog.com/entry/2019/08/06/040939

- golang error handling
    * https://zenn.dev/nobonobo/articles/a7f41596220a1b

- Change build path in create-react-app
    * https://qiita.com/yakimeron/items/7a4f8d9e70a4a2b1b96b
    * https://github.com/facebook/create-react-app/issues/1354#issuecomment-299956790

- react markdown
    * https://zenn.dev/rinka/articles/b260e200cb5258
    * https://harkerhack.com/react-markdown-editor/

## Author

[twitter@geniusmaaakun](https://twitter.com/geniusmaaakun)

## Licence

[MIT](https://opensource.org/license/mit/)
