# Whats Vidder? 
Vidder (as in *VIDeo downloaDER*) is a video downloading CLI, written in Go, that wraps the yt-dlp cli.

# Why Vidder?
0. Its just cooler to do stuff with terminals and CLIs rather than GUIs. It will make you feel like a hacker 😈. (and yes, real developers start counting from 0.)
1. Most youtube-video downloading apps have some sort of restriction on the freemium version (like only X downloads total, or no playlist downloading) 🤮, as well as chapters or timestamps not downloading correctly. With vidder 💪, you can have it ALL FOR FREE. (though a 🌟 to the project would be great 😉)
2. The yt-dlp cli is way to complex to use in an intuitive manner, whereas the yt-dlp GUI doesn't allow granular control of customization (for example, setting the resolution of the video to 360p.)
3. And why not just use spotify or youtube directly? Sure, if your battery is endless and you have cheap and good internet connection. In my case, with a laptop whose battery runs out faster than Usain Bolt, and an Internet Provider with which I have to save mobile data like pennies, it's just not feasable.
# Requirements:
1. A computer or laptop (Windows, Linux or Mac)
2. Basic terminal knowledge (how to open / close a terminal, how to copy / paste commands)
# Features:
- You can easily download any video or playlist that [yt-dlp](https://github.com/yt-dlp/yt-dlp) can download. (more than 1500 sites)
- You have granular control over how to download the file: from choosing to download it as an audio with maximum possible quality, to downloading as a video with 480p of resolution, including chapters (or timestamps)
- You can download a batch of videos given a txt file with the video URLS, separated by \n (newlines or ENTER). This happens concurrently, spawning up to 5 * *Number_Of_CPU_Cores* goroutines to make downloads faster.

# Installation:

## 0. Install yt-dlp:
If u need the docs, go to [yt-dlp's github](https://github.com/yt-dlp/yt-dlp). If you're lazy like me 😼, then here's the summary:

### On Windows:
Open *cmd* with admin rights and RUN ONE BY ONE THE FOLLOWING COMMANDS: 
 
```
mkdir -p "C:\Program Files\vidder" & curl -L https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp.exe -o "C:\Program Files\vidder\yt-dlp.exe"
setx PATH "C:\Program Files\vidder;%PATH%"

#close and reopen the terminal
yt-dlp --version #to check if it works
```
### On Linux:
```
sudo curl -L https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp_linux -o /usr/local/bin/yt-dlp
sudo chmod +x /usr/local/bin/yt-dlp
yt-dlp --version #check if it works
```
## 1- Install vidder:
### I. Using Pre-Built binaries (recommended)
#### On Bash terminal:

    ```
    #COPY, PASTE, AND RUN THESE COMMANDS ONE BY ONE IN THE TERMINAL. they will remove previous installation if any, then download vidder inside /usr/local/bin/vidder. then make it executable. 
    #Make sure to use the correct BINARY_NAME to download. We will use linux64 in the example below. available options are: (vidder-linux64, vidder-linux32, vidder-mac-amd, vidder-mac-arm), 

    sudo rm /usr/local/bin/vidder 
    sudo curl -L https://github.com/E-nkv/vidder/releases/download/1/vidder-linux64 -o /usr/local/bin/vidder
    sudo chmod +x /usr/local/bin/vidder 

    #Use it! 
    vidder videoURL 
    vidder -f absolute-path-to-txt.txt
    vidder --help 
    vidder -h
    ```
    
    

Tip: if you get an error like *vidder: command not found*, try the following: 
`export PATH=/usr/local/bin:PATH`. Though most distros have this PATH set by default, so most likely you won't need this.

#### On Windows (cmd only, since powershell is too ugly)
```powershell
#COPY, PASTE AND RUN ONE BY ONE THE FOLLOWING COMMANDS
#Make sure to use the correct BINARY_NAME. (vidder-win64 [default], vidder-win32)

del "C:\Program Files\vidder\vidder.exe" 
mkdir "C:\Program Files\vidder" 
curl -L "https://github.com/E-nkv/vidder/releases/download/1/vidder-win64.exe" -o "C:\Program Files\vidder\vidder.exe" 
setx PATH "C:\Program Files\vidder;%PATH%"

#Close and reopen the terminal. either powershell or cmd, without admin rights if you want.

# USE IT!
vidder videoURL 
vidder -f absolute-path-to-txt.txt
vidder --help 
vidder -h

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
