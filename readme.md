# infinitylink-go

This is just an experiment to write a real world application in Go.
I'll do the same with other languages as well.

## What does infinitylink-go do?
Your links will never break again!

infinitylink saves links to archive.org (aka wabackmachine). That way all
linked websites will be either available at their old location or as archived version.

Just link to infinitylink and don't worry about this any more.

## Usage

Instead of

    <a href="https://blog.com/a-cool-blog-article">Nice link</a>

Simply use

    <a href="https://infinitylink.yourserver.com/https://blog.com/a-cool-blog-article">Nice link</a>

This will
- Save https://blog.com/a-cool-blog-article to archive.org and make it eternal
- Check automatically whether to link to the original article - or - if the link is broken - link to the archived version.

Links will never break again. Couldn't be simpler.

## Deployment
Deployment is done via buildpacks / dokku: https://shellbear.me/blog/go-dokku-deployment

Rough description:

    init git repo
    git remote add dokku dokku@server.com:infinitylink-go
    Make sure to set the proper listening port for the app and we are ready to go.


## Things to explore from here

- Publish the mini-hamcrest library to github
- Experiment with closing the structs down => https://golangbot.com/structs-instead-of-classes/ => unexport New
- Use goroutine to make save call to waybackmachine not block the rest of the request
- Save data into database so that we "know" what has been saved already
- Save data into database so that we "know" articles that are offline (skip long call to the site)
- Make it a full fledged service with frontend login/logout to update the links
- Explore full DI framework

## Observations on Go

### Cool things
- Syntax very easy to understand
- Fast compilation
- Publishing a go lib publicly is so awesome and simple
- Test support is perfect
- Http and client as part of the std lib
- VsCode works perfectly fine as IDE
- Debugging go code works perfectly fine for both test and programs via vscode
- Deploy to live via dokku was a freaking piece of cake

#### Commands are awesome and self-contained
- go build
- go test ./...
- go test -cover ./...

### Not so cool
- Immutability does not exist
- Optional values are not supported by default.
- Default matchers of tests are clunky
- Structs have to be encapsulated to make sure fields are valid. Otherwise hard to trust.

