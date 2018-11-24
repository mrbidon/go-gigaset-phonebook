## Phonebook webservice for Gigaset

Aim of this app is to serve a phonebook webservice on a raspberrypi (or any linux device)
for Gigaset C530-IP in order to display incoming phone number thought a shared phonebook.

Database is stored in a csv file.

Ansible script are provided to setup a service.

Gigaset API Spec can be found here : https://teamwork.gigaset.com/gigawiki/display/GPPPO/Online+directory

## How it work

All code is in the `phonebook.go` file.

You can build and run it localy thought the following command line
```
go build
./phonebook numbers.csv 1234
```

The service can be tested thought this url (the only query my Gigagset id doing) :
```
curl "http://localhost:1234/?command=get%5flist&type=pb&fn=%2a&ln=%2a&ct=%2a&st=%2a&hm=<PHONENUMBER>&nr=%2a&mb=%2a&fx=%2a&sip=%2a&zc=%2a&em=%2a&in=%2a&bp=%2a&lang=9&first=1&count=1&mac=7C2F80D0C183&reqsrc=auto&handsetid=23546&limit=2048"
```


## Disclaimer

I'm not still using this device, but I'm sharing this code. I think it should be
usefull to anybody that need a personnal phonebook for its . I can only answer
questions thought issue.  
