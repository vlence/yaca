# yaca
Yet another chat app.

Why let Slack or another app manage your chats and then have all your embarrasing texts exposed via a data leak?
Do it yourself, on your own server!

## Technical Design

A chat is a series of messages over time.
```
Msg1->Msg2->Msg3->...
```

Each chat is just a sqlite database with all the chat messages.
```
UserA-UserB.chats = Msg1->Msg2
```
