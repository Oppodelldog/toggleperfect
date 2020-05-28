# toggle perfect - dev UI

This supports local development of the device software by emulating
the UI of the device.

for local development turn off the original device UI and enable  
the remote UI by settings two env variables:
```.env
TP_ENABLE_DEVICE_UI=false
TP_ENABLE_REMOTE_UI=true
```

Then open the index.html file in your Browser.
