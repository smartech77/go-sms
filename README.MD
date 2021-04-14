# Siminn SMS

# API Docs
https://sms.siminn.is/api

# Other Documentation
https://www.siminn.is/adstodarflokkur/sms-magnsendingar

# Example
main_test.go contains an example of how to initialize and use this package. In order to run the test you need to fill in the .env file.

# Notes
According to Siminn employees the proper way to send mass text messages is to send them in a loop. Because of this we will simply implement the send method and a method to verify that the message has been delivered while ignoring the old methods in the documentation.

# TODO
1. add more options for sending text messages
2. Add a method to confirm delivery
3. ? .. more to come