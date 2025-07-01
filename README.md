# Whats Vidder? 
Vidder (as in *VIDeo downloaDER*) is a video downloading CLI, written in Go, that wraps the yt-dlp cli.

# Why Vidder?
1. Most youtube-video downloading apps have some sort of restriction on the freemium version (like only X downloads total, or no playlist downloading) 🤮, as well as chapters or timestamps not downloading correctly. With vidder 💪, you can have it ALL FOR FREE. (though a ⭐ to the project would be great 😉)
2. The yt-dlp cli is way to complex to use in an intuitive manner, whereas the yt-dlp GUI doesn't allow granular control of customization (for example, setting the resolution of the video to 360p.)
3. And why not just use spotify or youtube directly? Sure, if your battery is endless and you have cheap and good internet connection. In my case, with a laptop whose battery runs out faster than Usain Bolt, and an Internet Provider with which I have to save mobile data like pennies, it's just not feasable.
# Requirements:
1. A computer or laptop (Windows, Linux or Mac)
2. Basic terminal knowledge (how to open / close a terminal, how to copy / paste commands)
# Features:
- You can easily download any video or playlist that [yt-dlp](https://github.com/yt-dlp/yt-dlp) can download. (more than 1500 sites)
- You have granular control over how to download the file: from choosing to download it as an audio with maximum possible quality, to downloading as a video with 480p of resolution, including chapters (or timestamps)


# Installation:

## 1. Using Pre-Built binaries (recommended approach):
### On Linux | Mac (bash):

    ```
    #remove previous installation if any
    sudo rm /usr/local/bin/vidder

    #download vidder inside /usr/local/bin/vidder. Make sure to use the correct BINARY_NAME. (vidder-linux64, vidder-linux32, vidder-mac-amd, vidder-mac-arm)
    sudo curl -L https://github.com/E-nkv/vidder/releases/download/1/vidder-linux64 -o /usr/local/bin/vidder
    
    #make it executable
    sudo chmod +x /usr/local/bin/vidder

    #USE IT!
    vidder --help OR vidder <URL>
    ```
    

Tip: if you get an error like *vidder: command not found*, try the following: 
`export PATH=/usr/local/bin:PATH`. Though most distros have this PATH set by default, so most likely you won't need this.

### On Windows (cmd only, since powershell is too ugly)
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

## 2. Manual (requires Go installed)

Open your terminal of choice, navigate to the folder you want to clone from, and then:

a (preferred). 
```
    go install github.com/E-nkv/vidder
    vidder <URL>
```
b.  (if you wanna check the source code and play with it)
```
    git clone https://github.com/E-nkv/vidder.git .

    #on UNIX-based OS:
    go build -o ./vidder .
    ./vidder <URL>

    #on windows:
    go build -o ./vidder.exe .
    ./vidder.exe <URL>
    
```