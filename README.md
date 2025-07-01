# Whats Vidder? 
Vidder (as in *VIDeo downloaDER*) is a video downloading CLI, written in Go, that wraps the yt-dlp cli.

# Why Vidder?
1. Most youtube-video downloading apps have some sort of restriction on the freemium version (like only X downloads total, or no playlist downloading) 🤮, as well as chapters or timestamps not downloading correctly. With vidder 💪, you can have it ALL FOR FREE. (though a 🌟 to the project would be great 😉)
2. The yt-dlp cli is way to complex to use in an intuitive manner, whereas the yt-dlp GUI doesn't allow granular control of customization (for example, setting the resolution of the video to 360p.)
3. And why not just use spotify or youtube directly? Sure, if your battery is endless and you have cheap and good internet connection. In my case, with a laptop whose battery runs out faster than Usain Bolt, and an Internet Provider with which I have to save mobile data like pennies, it's just not feasable.
# Requirements:
1. A computer or laptop (Windows, Linux or Mac)
2. Basic terminal knowledge (how to open / close a terminal, how to copy / paste commands)
# Features:
- You can easily download any video or playlist that [yt-dlp](https://github.com/yt-dlp/yt-dlp) can download. (more than 1500 sites)
- You have granular control over how to download the file: from choosing to download it as an audio with maximum possible quality, to downloading as a video with 480p of resolution, including chapters (or timestamps)


# Installation:

## 0. Install yt-dlp:
Yes, real developers start counting from 0. If u need the docs, go to [yt-dlp's github](https://github.com/yt-dlp/yt-dlp). If you're lazy like me 😼, then here's the summary:

### On Windows:
1. open cmd with admin rights and: 
a. Windows 64bit OS (most likely you are using this one): 

`mkdir -p C:\Program Files\vidder & curl -L https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp.exe -o "C:\Program Files\vidder\yt-dlp.exe"`

b. Windows 32bit OS: 

`mkdir -p C:\Program Files\vidder & curl -L https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp_x86.exe -o "C:\Program Files\vidder\yt-dlp.exe"`
  
2. add it to path: 

`setx PATH "C:\Program Files\vidder;%PATH%"` 
3. close terminal, and open either powershell or cmd again, and check if yt-dlp was installed properly: 

`yt-dlp --version`

### On Linux (64bit):
1. Open UNIX terminal (ubuntu, wsl) and download yt-dlp:
`sudo curl -L https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp_linux -o /usr/local/bin/yt-dlp`
2. Make it executable
`sudo chmod +x /usr/local/bin/yt-dlp`
3. Check if it works: 
`yt-dlp --version`
## 1- Install vidder:
### I. Using Pre-Built binaries (recommended approach)
#### On Linux | Mac (bash):

    ```
    #remove previous installation if any
    sudo rm /usr/local/bin/vidder

    #download vidder inside /usr/local/bin/vidder. Make sure to use the correct BINARY_NAME. (vidder-linux64, vidder-linux32, vidder-mac-amd, vidder-mac-arm)
    sudo curl -L https://github.com/E-nkv/vidder/releases/download/1/vidder-linux64 -o /usr/local/bin/vidder
    
    #make it executable
    sudo chmod +x /usr/local/bin/vidder

    #REFRESH THE TERMINAL. (Close and reopen)

    #USE IT!
    vidder --help OR vidder <URL>
    ```
    

Tip: if you get an error like *vidder: command not found*, try the following: 
`export PATH=/usr/local/bin:PATH`. Though most distros have this PATH set by default, so most likely you won't need this.

#### On Windows (cmd only, since powershell is too ugly)
```powershell

#OPEN CMD WITH ADMIN RIGHTS (windows key and search "cmd". right click and hit "open with admin rights")

#remove previous installation if any
del "C:\Program Files\vidder\vidder.exe" 

mkdir "C:\Program Files\vidder" 

#download vidder inside C:\ProgramFiles\vidder. Make sure to use the correct BINARY_NAME. (vidder-win64, vidder-win32)
curl -L "https://github.com/E-nkv/vidder/releases/download/1/vidder-win64.exe" -o "C:\Program Files\vidder\vidder.exe" 

#add it to the PATH to have vidder globally available
setx PATH "C:\Program Files\vidder;%PATH%"
#if the terminal shows some error related to PATH being too large / truncated, you will need to add "C:\Program Files\vidder" from the windows GUI in "advanced system settings --> environment variables"

#CLOSE AND REOPEN THE TERMINAL. it doesnt have to be on with admin rights this time

# USE IT!
vidder --help OR vidder <URL>

```

### II. Manual (requires Go installed)

Open your terminal of choice, and then:

#### 1. go install vidder (recommended). 
```
    go install github.com/E-nkv/vidder
    vidder <URL>
```
#### 2. clone project (if you wanna check the source code and play with it)
```
    #navigate to the folder you want to clone from
    cd FOLDER_TO_CLONE_FROM

    git clone https://github.com/E-nkv/vidder.git .

    #on UNIX-based OS:
    go build -o ./vidder .
    ./vidder <URL>

    #on windows:
    go build -o ./vidder.exe .
    ./vidder.exe <URL>
    
```

# Did you like VIDDER-CLI?
If so, please leave a star 🌟 and share it with your friends (if you have any 🥲☠️). I'm sure they'll love it.
Also, make sure to check [yt-dlp](https://github.com/yt-dlp/yt-dlp), since their magic was what allowed VIDDER-CLI to shine 😃.

# Want to collaborate?
The project works, but there surely are great things to add to it 🤠!
Quick list of some of them:
- Make it available in more languages, both the README instructions and the CLI. as for the CLI itself, we would need to have multiple binaries for each of the languages, which would indeed be an interesting challenge to take.
- Test and modify both /core and /cli to handle downloads in more websites (like IG, Tiktok. note: some of them require API keys or tokens). as of now, only youtube has been tested thoroughly.
- A simple web-based UI (html-css-js bundle) to download and watch videos,that runs locally. This will be quite the challenge, from making the downloader work (since it talks to bash / powershell under the hood. lots of juicy permissions issues to work with at the OS level 😼.) to making the video player work offline, handling different speeds to showing timestamps/chapters .
- Anything else you might have noticed that could be improved

Either way, send a PR and I'll be glad to look into it and make [VIDDER](https://github.com/E-nkv/vidder) much cooler & better 😃.
