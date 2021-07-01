<!-- PROJECT LOGO -->
<br />
<p align="center">

  <h1 align="center">Stock Alert System</h1>

  <p align="center">
    Set alerts on stocks with Open Source
    <br />
    <!-- <a href="https://github.com/"><strong>Explore the docs »</strong></a> -->
    <br />
    <a href="https://github.com/revulcan/stock-alert-system/issues">Report Bug</a>
    ·
    <a href="https://github.com/revulcan/stock-alert-system/issues">Request Feature</a>
  </p>
</p>

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#install-necessary-programs">Install Necessary Programs</a></li>
        <li><a href="#developer-account">Developer Account</a></li>
        <li><a href="#modify-the-.env.example-file">Modify the .env.example file with the following credentials</a>
        </li>
      </ul>
    </li>
    <li><a href="#communicating-with-the-service">Communicating with the service</a></li>
    <li><a href="#golang-client-example">Golang Client Example</a></li>
    <li><a href="#node-client-example">Node Client Example</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project

It is a microservice to set various types of alerts based on different parameters of a stock such as price to more complicated strategies such as SMA. This helps traders to get automated stock alerts for AlgoTrading.


<br />

[![Contributors][contributors-shield]][contributors-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

### Built With

This project was built using golang along with GRPc to build a distributed system.
* [GRPc](https://github.com/grpc/grpc-go)
* [smartapigo](https://github.com/angelbroking-github/smartapigo)
* [gokiteconnect](https://github.com/zerodha/gokiteconnect)

<br />

<!-- GETTING STARTED -->
## Getting Started

<br />

### Install Necessary Programs
* Golang

<br />

### Developer Account
* [Angel SmartApi Developer account (create one if you don't have)](https://smartapi.angelbroking.com/docs)
* [NOT SUPPORTED YET] KiteConnectAPI Developer account (create one if you don't have)

<br />

### Modify the .env.example file
 
<br />

```environment
# Change endpoint if needed
UPDATES_POST_ENDPOINT=http://localhost:8081/onTrigger 

ANGEL_APIKEY=xjkdvcnd    # your api key
ANGEL_CLIENTID=A12345    # your username
ANGEL_PASSWORD=Password  # your password
```

<br />

## Install the microservice as a linux system service

<br />

```bash
$ chmod +x deploy-service.sh
$ ./deploy-service.sh
password:
```
<br />

## Communicating with the service

<br />

We can communicate with the service easily by compiling `.proto` file to any of the GRPC's abundantly supported list of languages.

I have added client examples under `/grpc/client/` where there are two client:
* Golang Client
* Node Client

You can use the `gen_proto.sh` script in `/grpc/trigger_service/` to recompile the proto and generate the client libraries if any changes are made to the `.proto` file

<br />

## Golang Client Example

<br />

```go
func CreateTrigger(client trigger_service.TriggerServiceClient, req *trigger_service.CreateTriggerReq) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
    // the client here is the GRPC client definitions generated using .proto files
	res, err := client.CreateTrigger(ctx, req)

    // res contains the response of the call
	if err != nil {
		fmt.Println("CreateTrigger encountered error : ", err)
	} else {
		fmt.Println("CreateTrigger Response : ", res)
	}
}
```

<br />

## Node Client Example

<br />

```javascript
createTriggerRes = await client.createTrigger({
        id: '3',
        tAttrib: 'LTP',
        operator: 'GTE',
        tPrice: 302,
        tNearPrice: 292,
        scrip: 'SBIN',
        kiteToken: '',
        exchangeToken: '3045',
        Exchange: 'NSE'
      });
console.log(createTriggerRes)
```
<br />

You can take a look at [GRPC's Supported Languages](https://grpc.io/docs/languages/) to get started with using the trigger service in your application.

<br />


<!-- ROADMAP -->
<!-- ## Roadmap

See the [open issues](https://github.com/revulcan/stock-alert-system/issues) for a list of proposed features (and known issues). -->



<!-- CONTRIBUTING -->
## Contributing

<br />

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<br />

<!-- LICENSE -->
## License

<br />

Distributed under the MIT License. See `LICENSE` for more information.

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/workflow/status/revulcan/stock-alert-system/Test/master
[contributors-url]: https://github.com/revulcan/stock-alert-system/actions/
[stars-shield]: https://img.shields.io/github/stars/revulcan/stock-alert-system?style=social
[stars-url]: https://github.com/revulcan/stock-alert-system/stargazers
[issues-shield]: https://img.shields.io/github/issues/revulcan/stock-alert-system
[issues-url]: https://github.com/revulcan/stock-alert-system/issues
[license-shield]: https://img.shields.io/github/license/revulcan/stock-alert-system
[license-url]: https://github.com/revulcan/stock-alert-system/blob/master/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/nithin-rao/