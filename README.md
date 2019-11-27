# DealzNotifier

Context: 
I'm a regular user of applications / websites such as Dealabs which is a website used to regroup best deals.
You have the native function on the official dealabs app to create alert, so you can just done it like that.

This project was initially made just to learn Golang and to see whats going on with the language.
Doing my researchs, I found that Dealabs is the french website but there are many more (Chollometro, Hot UK Dealz, Pepper.nl...)
which are all made by the Pepper company. So, each structures of the websites are nearly the same.
The differences are that, each deals are posted by users, so from a website to another, you can find differents deals.
 
My idea was to develop a scrapper and to use it to combine alerts for all thoses websites.
This alert will be sent to me phone. At first, wanted to use Whatsapp to send my alert, but as I haven't lots of time to develop the project,
I use telegram, with a library which was easier to implement.

The code is probably not the best, but as I do what I want.. whatever.

The program is permanently running on a Raspberry, so that the handler with Telegram API is fully operationnal everytime.

Execution:
Change the token to fit with yours in `main.go` :
```
	Token:  "672292993:AAEY5S2ETZcc1_tCUMEPqE4GnshRwtLp3PM"
```

Warning : Don't try to use this one, this is an obsolete one...

- go build src/*
- ./main
