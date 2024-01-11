
Requirements
	A [Netlify](netlify.com) account (free)
	[Obsidian](https://obsidian.md/) (free)
	[Sophistry](https://github.com/DrVeseli/Sophistry) (free)

The Netlify account is used to host the website, Obsidian is the note taking tool.

In Netlify you can go trough setting up a custom website with custom code, it might be easier to use graphic tools, but all is available trough the CLI you need to install.

First step is to download pnpm:
	on windows, open powershell and paste this in 
	
```
iwr https://get.pnpm.io/install.ps1 -useb | iex
```

	If you are using Linux or MacOS
	
```
curl -fsSL https://get.pnpm.io/install.sh | sh -
```

To install node run:

```
pnpm env use --global lts
```

To install Netlify CLI:

```
pnpm install netlify-cli -g
```

To log in to Netlify from the terminal use:

```
netlify login
```

It should take you to your browser to confirm the login for your Netlify account. You are almost set, Sophistry is a GitHub link, you can download a compiled binary from [here](https://github.com/DrVeseli/sophistry/releases) or compile it yourself, It requires GO. From the source files download the Blog/structure folder, inside you will find a header, footer and some CSS. You are encouraged to change these to fit your vision of the blog, especially the logo and the name in the header.html file.

We are almost done, open Obsidian and create a vault in a folder. Copy the executable for Sophistry and the Blog/structure folders into the Obsidian vault you created.

Start writing, the file structure and the folders you create in your notes will be reflected on your home page, and the induvial notes will become the articles.

Once you are happy with what you written just double click the executable from your file explorer and the Blog folder will be shipped to Netlify for the world to discover. The first time you do this Netlify might ask you to confirm what website you are referring to but the tool is intuitive, just read the prompt and you will have no trouble. After the initial setup you will probably only have to run the executable to publish changes that you made in Obsidian.

disclaimer Netlify CLI is not maintained by me so it is subject to change, It should never break [Sophistry] but it might affect the procedure somewhat!

Happy writing!
