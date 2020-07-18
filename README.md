# visitor-count

A simple micro-application to generate a "Visitor Count" SVG, which can be
embeded to GitHub Profile Readme's, for example.

Example:  
![](https://ghvisitorcount.zekro.de/repo-demo.svg)

## Configuration

The application server is looking for a `config.yaml` or `config.json` file either at the root directory of the application or in a sub directory named `config/`.

Take a look at the [**example config**](config/config.yml) to see how to configure the application.

You can also specify environment variables for configuration. The variable keys need to have the prefix `VC_`, are all uppercase and groups are seperated with periods (`.`). Example:
```
VC_WS.ADDR="0.0.0.0:8080"
VC_WS.IPWHITELIST="140.82.115.* 172.217.23.174 62.171.147.124"
```

---

© 2020 Ringo Hoffmann (zekro Development)  
Covered by the MIT Licence.